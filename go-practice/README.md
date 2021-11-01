
## Golang Context

### Select 

select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。


### ctx.Done()
ctx.Done() means that we receive the order to finish our work.  In this case, we will interrupt the loop, log a message and return.

详细请参考[这一章](https://www.practical-go-lessons.com/chap-37-context) 

> 在上面的程序中
we will use a channel to communicate with the caller. We create a for loop, and inside that loop, we will put a select statement. In this select statement, we have two cases :

The channel returned by ctx.Done() has been closed. It means that we receive the order to finish our work

In this case, we will interrupt the loop, log a message and return.
The default case (executed if any previous case are not executed)

In this default case, we will increment the sum.

If the variable sum is becoming greater strictly than 1.000, we will send the result on the result channel