

# Goroutine





将异步函数的决定权交给该函数的调用放通常更容易的多

也不是更容易吧,更合理一点儿

# Memory model





一个指针的赋值一原子的



go map 在runtime 里面是一个 HMap 的结构体



满足原子,但是不满足可见性



一个goroutine改掉的值,在另一个gouroutine里面看不到

---

sync 纪要控制好并发的执行,又要控制好goroutine 的生命周期



# Package Sync



go memory model 

 

互斥锁



