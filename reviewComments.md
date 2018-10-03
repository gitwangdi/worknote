# gofmt
vscode does it when we use ctrl+s.

# Comment Sentences 
comment should be a full sentence, begin with the name of the thing being described. 

# context ??
Ctx should be the first parameter.  

# copy 
take care the name

# delcare empty slices 
var t []string  
declare nil string slice  
t := []string{}  
declare non-nil but zero-lenth string slice  

# crypto rand 
use:  
crypto/rand  
instead of:  
math/rand  
  
(use hex.EncodeToString(buf) or base64.StdEncoding.EncodeToString(buf) to get text)  

# doc comment
Export function or variables need comment   
Package need /* */  

# panic
尽量不使用panic  
使用error返回值  

# error strings
错误信息中的字符串不要大写开头，因为在后面可能的输出过程（log）中要在前面加入其他字符  

# example 
一个新包，要带有对应的测试 

# go routine  
保证包退出的时机  
慎重使用channel的阻塞功能  

# handle error
不要掠过错误  

# imports 
Avoid renaming imports.  
Devide imports in groups, and standard library packages in the first group.  
(goimports)

# import dot 
只有当构成了循环依赖的情况下才使用import . ，否则造成代码读困难  

# in-band error 
go 可以使用多返回值，不必再遇到错误的时候返回空或者0.  

# indent error flow
if处理error信息的时候，normal code尽量不要放到else里面

# initalisms
特定缩写 比如URL instead of Url  

# interface 
interface 放在使用interface的包里面（以interface为参数），而不是具体实现interface的包里。

# line length
别太长，根据语义去控制换行，不要单纯因为换行而换行。

# named result parameters
返回值命名，一般是不必要的，除非：有两个以上相同类型的返回值，可以用变量名区分， 或者其他必要原因。  

# package name
最好是有实际意义的名字，用来：  
package chubby中有一个变量为FILE,在别的包里面使用则名为chubby.FILE。尽量不要使用util,common,misc,api一类的包名 

# pass values
注意指针参数的使用

# receiver names 
函数接受的参数尽量命名简短：client->cl/c  
根据功能命名，而不是面向对象  

# receiver type 
慎重使用指针

# test failures
打印错误的输入 输出 和 正确期望结果

# value names
名字的适用范围越远，名字的长度才越长，如果是在不大函数中参数的名字可以使一到两个字母，循环中只需要i或r，当变量更加特殊的时候待会需要更具有描述性的名字