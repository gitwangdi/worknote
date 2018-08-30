## HelloWorld

![示意图](helloworld.png)

最简单的rabbitMQ应用，不使用exchange，直接使用queue进行producer和consummer的对接。  
发送信息的流程是：建立连接->开启Channel->建立Queue->Publish消息  
接收消息的流程是：建立连接->开启Channel->建立Queue->Consume监听信息  

```go
// 建立连接：
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

// 开启Channel：
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

// 建立Queue：
q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable  true：当服务器重启，queue仍然存在
  false,   // delete when unused
  false,   // exclusive  true：当connecting断开，则queue删除
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")

// Publish信息：
body := "Hello World!"
err = ch.Publish(
  "",     // exchange  当使用Exchange的时候用来做标识
  q.Name, // routing key  用来作为sender和receiver的联络暗号，exchange只发送对应暗号的消息
  false,  // mandatory
  false,  // immediate
  amqp.Publishing {
    ContentType: "text/plain",
    Body:        []byte(body),
  })
failOnError(err, "Failed to publish a message")

// 监听信息：
msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  true,   // auto-ack  自动回应时间
  false,  // exclusive  
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")
```

处理信息的时候可以使用一个chan来保持server的运行，使用go func()处理信息。

```go
forever := make(chan bool)
go func() {
  for d := range msgs {
    log.Printf("Received a message: %s", d.Body)
  }
}()
log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever
```

## Work Queue

![示意图](workqueue.png)

当有两个以上的接收者的情况。  
queue会进行平均分配，即1、3、5分给receiver1，2、4分给receiver2。  
这一分配过程是在两个接收者执行任务之前的，即使两个接收者没有完成哪怕一个任务，仍然会把所有该给他的消息给他。这一情况下，当一个接收者dies，则他的任务就会丢失。  
这一问题的解决方法是：将consume过程中的auto-ack变量设置为false，然后在执行过程的最后执行ack

```go
// 设置auto-ack为false
msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  false,  // auto-ack
  ...
)

// 处理过程
go func() {
  for d := range msgs {
    // 具体工作
    d.Ack(false)
  }
}()
```

之前的例子中queue是依赖于rabbit-server而存在的，所以当rabbit-server重启之后，其上的queue是会丢失的。想要使queue不消失，需要对queue进行设置，修改其中的durable变量为true。  
但是这只能保证在下一次rabbit-server重启后，queue仍然存在，但不能保证其上的信息不丢失，想要保证信息不丢失还需要在发送信息的时候，在Publish中添加persisten的属性。  

```go
err = ch.Publish(
  "",           // exchange
  q.Name,       // routing key
  false,        // mandatory
  false,
  amqp.Publishing {
    DeliveryMode: amqp.Persistent,
    ContentType:  "text/plain",
    Body:         []byte(body),
})
```

## Publish/Subscribe

![publish](publish.png)

在前面的例子中，两个接收者使用的是同一个queue，这一情况狂下再分配任务的时候就会出现问题。  
这一节针对的就是两个接收者分开使用queue。  
流程更改如下：  
发送方：建立连接->开启Channel->构建Exchange->Publish信息  
接收方：建立连接->开启Channel->构建Exchange->构建Queue->构建QueueBind->Consume接收消息  
事实上，在前面的应用中也使用到了Exchange，但是由于没有分开queue的需要，我们使用的是default exchange，不需要我们去定义。

```go
// 发送方
err = ch.ExchangeDeclare(
  "logs",   // name
  "fanout", // type  fanout设置为广播，即所有消息播报给每一个queue， 自动忽略routine key
  true,     // durable
  false,    // auto-deleted
  false,    // internal
  false,    // no-wait
  nil,      // arguments
)
err = ch.Publish(
  "logs", // exchange
  "",     // routing key
  ...
)
```

```go
// 接收方
// queue可以使用任意名字，可以设为“”设定随机名
// QueueBind
err = ch.QueueBind(
  q.Name, // queue name
  "",     // routing key
  "logs", // exchange
  false,
  nil
)
```

## Direct exchange

![direct](direct.png)

在Exchange中设定type为direct后，在分配信息的时候会进行筛选，凡是符合routine key的信息才会被发送到对应的queue中。  
具体到代码，sender方在Publish中设置第二个参数routine-key，接收方在Bind中设置此参数。  

## Topic exchange

![topic](topic.png)

在Exchange中设定type为topic后，同样和direct一样对信息进行筛选，但是使用的过滤方法是c1.c2.c3的形式。  
*.*.orange,可以接收到所有标记中第三个属性为orange的信息,但是一定要有前两个属性，同时不能有第四个属性。lazy.#，则可以收到所有第一个属性为lazy的信息，不限后面还有多少个属性。  
