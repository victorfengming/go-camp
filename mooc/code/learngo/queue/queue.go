package queue

// An FIFO queue.
type Queue []int

// Pushes the element into the queue.
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// callback the wheather or not the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

/**

PS E:\Projects\GolandProjects\go-camp\mooc\code\learngo\queue> go doc queue
package queue // import "learngo/queue"

type Queue []int
PS E:\Projects\GolandProjects\go-camp\mooc\code\learngo\queue>


PS E:\Projects\GolandProjects\go-camp\mooc\code\learngo\queue> go doc queue Pop
package queue // import "learngo/queue"

func (q *Queue) Pop() int
    Pops element from head.

PS E:\Projects\GolandProjects\go-camp\mooc\code\learngo\queue>

*/
