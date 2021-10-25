





# 目录

•Error Type

•Handling Error

•Go 1.13 errors

•Go 2 Error Inspection

•References



---

![1635135863822](README/1635135863822.png)





Struct Type Error这个是不会打印的

因为在栈上new的两个对象,他们是不同的(地址)

从我们的设计上,来避免一些情况,比如error的string内容相等的情况





![1635136052139](README/1635136052139.png)

 这个时候就会相等了

因为重写了这个NewError

这样还是会导致意外的 判定相等

这样也是不行的

---

各个语言的演进历史：

•C

  *单返回值，一般通过传递指针作为入参，返回值为* *int* *表示成功还是失败。*

•C++

  *引入了 exception，但是无法知道被调用方会抛出什么异常。*

•Java

  *引入了* *checked exception，**方法的所有者必须申明，调用者必须处理。在启动时抛出大量的异常是司空见惯的事情，并在它们的调用堆栈中尽职地记录下来。Java 异常不再是异常，而是变得司空见惯了。它们从良性到灾难性都有使用，异常的严重性由函数的调用者来区分。*



---

#      **Error vs Exception**  

Go 的处理异常逻辑是不引入 exception，支持多参数返回，所以你很容易的在函数签名中带上实现了 error interface 的对象，交由调用者来判定。

*如果一个函数返回了* *(value, error)**，你不能对这个* *value* *做任何假设，必须先判定* *error**。唯一可以忽略* *error* *的是，如果你连* *value* *也不关心。*

Go 中有 panic 的机制，如果你认为和其他语言的 exception 一样，那你就错了。当我们抛出异常的时候，相当于你把 exception 扔给了调用者来处理。

*比如，你在* *C++* *中，把* *string* *转为* *int**，如果转换失败，会抛出异常。或者在* *Java* *中转换* *String* *为* *Date* *失败时，会抛出异常。*

*Go panic* *意味着* *fatal error**（就是挂了）。不能假设调用者来解决* *panic**，意味着代码不能继续运行。*

使用多个返回值和一个简单的约定，Go 解决了让程序员知道什么时候出了问题，并为真正的异常情况保留了 panic。



## 演示

![1635136654124](README/1635136654124.png)

![1635136686875](README/1635136686875.png)

![1635136712479](README/1635136712479.png)

![1635136749354](README/1635136749354.png)

![1635136789970](README/1635136789970.png)





---





异常和错误是两种情况



错误是可以处理的

异常是不可恢复的一种情况

---



对于真正意外的情况，那些表示不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出，我们才使用 panic。对于其他的错误情况，我们应该是期望使用 error 来进行判定。

You only need to check the error value if you care about the result. -- Dave

This [blog post ](https://devblogs.microsoft.com/oldnewthing/?p=36693)from Microsoft’s engineering blog in 2005 still holds true today, namely:

My point isn’t that exceptions are bad. My point is that exceptions are too hard and I’m not smart enough to handle them.

•简单

•考虑失败，而不是成功（plan for failure, not success）

•没有隐藏的控制流

•完全交给你来控制 error

•Error are values



![1635137165097](README/1635137165097.png)

> 作者自谦: exception 这种情况比较难,以至于我不会更聪明的处理它,所以我们go里面用 panic和error



比如上面的代码,在AddToLeague(guy)这行如果报错, Exception 会出现

就会有这种场景: 一个guy他加入了这个联盟,但是你在 查看Team的时候蹦掉了

所以在用error这种方式,能尽可能的避免这种情况,因为goer 鼓励你做 及时的处理

