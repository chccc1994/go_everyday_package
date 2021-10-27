## Sync
[Sync](https://studygolang.com/pkgdoc)
>sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。本包的类型的值不应被拷贝。

golang 并不推荐这个包中的大多数并发控制方法，但还是提供了相关方法，主要原因是golang中`提倡以共享内存的方式来通信`:不要以共享内存的方式来通信，作为替代，我们应该`以通信的手段来共享内存`。共享内存的方式使得多线程中的通信变得简单，但是在并发的安全性控制上将变得异常繁琐。

- type  
>包中主要有: 
    Mutex：互斥锁
    RWMutex：读写锁
    WaitGroup：等待组
    Once：单次执行
    Cond：信号量
    Pool：临时对象池
    Map：自带锁的map


[Golang并发的次优选择：sync包](https://lessisbetter.site/2019/01/04/golang-pkg-sync/)