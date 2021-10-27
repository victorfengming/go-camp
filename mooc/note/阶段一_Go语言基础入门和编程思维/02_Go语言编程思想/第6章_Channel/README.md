

# 6-1 channel



## code1

```go
package main

import "fmt"

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)
	c <- 1
	c <- 2
	c <- 3
	n:= <-c
	fmt.Println(n)

}

func main() {
	chanDemo()
}

/**
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.chanDemo()
	E:/Projects/GolandProjects/go-camp/mooc/code/learngo/channel/channel.go:8 +0x37
main.main()
 */
```



> 死锁了

发的数据没人收是会deadlock的

![1635301183810](README/1635301183810.png)

## code2

```go
package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)

	go func() {
		for  {
			n:= <- c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2
	//c <- 3
	//n := <-c
	//fmt.Println(n)
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**
1
2

Process finished with the exit code 0
*/

```





> go语言函数是一等公民

> go语言中的channel 也是一等公民

## code3

```go
package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)

	go worker(c)

	c <- 1
	c <- 2
	//c <- 3
	//n := <-c
	//fmt.Println(n)
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**
1
2

Process finished with the exit code 0
*/

```





---



## code4

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int) {
	for {
		fmt.Println(id,<-c)
	}
}

func chanDemo() {

	var cahnneles [10]chan int
	for i := 0; i < 10; i++ {
		//var c chan int // c == nil
		cahnneles[i] = make(chan int)
		go worker(i,cahnneles[i])
	}

	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**
5 102
1 98
2 99
3 100
4 101
0 97
6 103
7 104
9 106
8 105

Process finished with the exit code 0

*/

```





在打印

## code5

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("%d,%c\n", id, <-c)
	}
}

func chanDemo() {

	var cahnneles [10]chan int
	for i := 0; i < 10; i++ {
		//var c chan int // c == nil
		cahnneles[i] = make(chan int)
		go worker(i, cahnneles[i])
	}

	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**
5,f
2,c
0,a
1,b
3,d
0,A
1,B
9,j
8,i
4,e
4,E
3,D
7,h
2,C
6,g
6,G
7,H
9,J
5,F
8,I

Process finished with the exit code 0


*/

```

> goroutine 调度之后,先发的不一定会先收到

## code6

```go
package main

import (
	"fmt"
	"time"
)

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("@%d---%c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {

	var cahnneles [10]chan int
	for i := 0; i < 10; i++ {
		//var c chan int // c == nil
		cahnneles[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**

@0---a
@8---i
@1---b
@1---B
@3---d
@4---e
@5---f
@6---g
@7---h
@0---A
@9---j
@2---c
@2---C
@5---F
@6---G
@7---H
@3---D
@4---E
@9---J
@8---I

Process finished with the exit code 0

*/

```

> // // 告诉外面用的人 , 我这个channel怎么用

![1635302653609](README/1635302653609.png)



这样就不能收数据了,只能发 属性 (send-only type)

```go
//n:= <-cahnneles[i] 
// ↑ Invalid operation: <-cahnneles[i] (receive from the send-only type chan<- int)
```



---



## code7

```go
package main

import (
	"fmt"
	"time"
)

// // 告诉外面用的人 , 我这个channel怎么用
func createWorker(id int) chan<- int { // 告诉外面用的人 , 我这个channel怎么用
	//
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("@%d---%c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {

	var cahnneles [10]chan<- int
	for i := 0; i < 10; i++ {
		//var c chan int // c == nil
		cahnneles[i] = createWorker(i)
		//n:= <-cahnneles[i]
		// ↑ Invalid operation: <-cahnneles[i] (receive from the send-only type chan<- int)
	}

	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)

}

func bufferedChannel() {
	c := make(chan int)
	c <- 1
}

func main() {
	//chanDemo()
	bufferedChannel()
}

/**
goroutine 1 [chan send]:
main.bufferedChannel(...)
	E:/Projects/GolandProjects/go-camp/mooc/code/learngo/channel/channel.go:42
main.main()
	E:/Projects/GolandProjects/go-camp/mooc/code/learngo/channel/channel.go:47 +0x31

Process finished with the exit code 2

 */

```



> 这样就是到4个 追加的时候才报错



## code8

```go
func bufferedChannel() {
	// 加上缓冲区,大小为3
	c := make(chan int,3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

}
```





## code9

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int) {
	for {
		fmt.Printf("@%d---%c\n", id, <-c)
	}
}


func bufferedChannel() {
	// 加上缓冲区,大小为3
	c := make(chan int,3)
	go worker(0,c)
	c <- '1'
	c <- '2'
	c <- '3'
	//c <- 4
	time.Sleep(time.Millisecond)



}

func main() {
	//chanDemo()
	bufferedChannel()
}

/**
@0---1
@0---2
@0---3

Process finished with the exit code 0

 */

```

> 告诉接收方,何时发完了



## code10

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		n,ok := <-c
		if ok{
			fmt.Printf("@%d---%d\n", id,n)
		}else{
			break
		}
	}
}


func bufferedChannel() {
	// 加上缓冲区,大小为3
	c := make(chan int)
	go worker(0, c)
	c <- '1'
	c <- '2'
	c <- '3'
	c <- 'd'
	//c <- 4
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {
	//chanDemo()
	bufferedChannel()
}

/**
@0---49
@0---50
@0---51
@0---100

Process finished with the exit code 0

*/

```







# 6-2 使用Channel等待任务结束







# 6-3 使用Channel进行树的遍历



# 6-4 Select



# 6-5 传统同步机制



# 6-6 并发模式（上）



# 6-7 并发模式（下）



# 6-8 并发任务的控制