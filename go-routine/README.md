

# 保护共享资源的方法

1、sync.Mutex互斥锁

2、sync.atomic原子写入

3、使用channel进行同步

PS：一定要注意第3种方法，在我们的代码中，当使用unbuffered channel时，会出现deadlock的情况，原因如下：

1. 主routine在等待wg.Wait()执行完毕，才能执行 fmt.Println(<-ch2)
2. 另外两个routine中较慢的一个叫做慢routine。慢routine在最后一次循环执行 ch2 <- count时，必须得有人接收才行，目前唯一有接收能力的语句就是主routine中的 fmt.Println(<-ch)，但是fmt.Println(<-ch)必须在慢routine执行完毕后才能执行，因此deadlock了。