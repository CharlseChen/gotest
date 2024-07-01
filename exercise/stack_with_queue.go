package exercise

type myStack struct {
	queue []int
}

var MyStack = &myStack{
	queue: make([]int, 0),
}

func (q *myStack) Push(x int) {
	q.queue = append(q.queue, x)
}

func (q *myStack) Pop() int {
	n := len(q.queue) - 1
	for n != 0 {
		val := q.queue[0]
		q.queue = q.queue[1:]
		q.queue = append(q.queue, val)
		n--
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

func (q *myStack) Peek() int {
	val := q.Pop()
	q.queue = append(q.queue, val)
	return val
}

func (q *myStack) IsEmpty() bool {
	return len(q.queue) == 0
}
