# Connection

## Dial()  
Connect to the server, try erver 10 seconds, deadline is 30 seconds.  
**DialConfig()**  Dial with config for connection setup  
**DialTLS()** Safe Dial(). witl tls.Config  
  
## Open()
Open an established connection, or use other **custem transport**.  

## Connection.Channel()
Open a chennal to pass the message. Any error will **close this channel**, like exchangeDeclare error.

## Connection.Close()
Close the connection and channel and so on.

## Connection.ConnectionState()
Return TLS detail

## Connection.LocalAddr()
Return net.Addr.

## Connection.NotifyBlocked()
Block the server when TCP server is busy.  [Example](https://godoc.org/github.com/streadway/amqp#Connection.NotifyBlocked)

## Connection.NotifyClose()
Listener for close event?

# Type 

## Delivery
The field comes from Channel.Consume() or Channel.Get()  
**Ack()** Acknowledgement.Multiple, the prior messages??  
**Nack()** Negatively acknowledgement. If requeue is true, the message will be send to other consumers. Multiple ???
**Reject()** Negatively acknowledgement for a single message.

## Type PlainAuth
Authentication including username and password.  

## Type publishing
Include the message to be sent to the server.  
Header, property and body.

# Channel
Used as a context for message exchange. Any error will close it.  
**Ack()** is same to Delivery.Ack()
**Nack()** the same

## Channel.Cancel()
Stop Channel.Consume, before stopping receive all the message in the server. Set noWait to stop immediately and the message may be dropped.  

## Channel.Close()
Close the Channel.

## Channel.Confirm()
Used to get the reply from publishing

## Channel.Consume()
Must range over the chan to ensure all deliveries are received, or blocking.??  
autoAck: Acknowledge automatic.  
exclusive: The only consumer for the queue.  
noLocal: Not supported by RabbitMQ.
noWait: Do not wait for server to confirm?

## Channel.Exchange

### ExchangeDeclare()
Durable: Survive server restarts.  
autoDelete: Delete when no binding.  
internal: 内部交易  
noWait: 
### ExchangeDeclarePassive()
### ExchangeDelete()
ifUnused: Only delete unbinded exchange  
noWait: 
### ExchangeBind()
### ExchangeUnbind()

## Channel.Flow()
False to pauses the delivery. True to react.  
Need NotifyFlow() ??  

## Channel.Get()

## Channel.Publish()
Send message.

## Channel.Qos()
?? limit the number of the messages in the channel  

## Channel.Queue
### QueueDeclare()
### QueueDeclarePassive()
### QueueDelete()
### QueueBind()
### QueueUnbind()
### QueueInspect()
Return queue with the given name. Can be used to get queue informations.
### QueuePurge()
Remove the messages in the queue, but the messages that have been delivered will not be removed.

## Channel.Recover()
Redelivers the unack deliveries on the channel.  
If requeue is false, delivered to the original consumer.
If requeue is true, delivered to any available consumer.  
Cannot be used on rabbitMQ.

## Channel.Tx()
Atomically(transaction mode)  
It is said to be a mode that we write series of codes and implement together.
### Channel.TxCommit()
### Channel.TxRoollback()

## Channel.Notify
Used to get reply, and do further operation.