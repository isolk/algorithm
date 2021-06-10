package queue

type Queue struct {
	data     []interface{}
	tail     int
	head     int
	size     int
	capasity int
}

func NewQueue(caposity int) *Queue {
	return &Queue{data: make([]interface{}, caposity), capasity: caposity}
}

func (q *Queue) Enqueue(val interface{}) {
	// 已满
	if q.Full() {
		panic("full")
	}

	// 达到尾部,跳到头
	if q.tail == q.capasity {
		q.tail = 0
	}

	q.data[q.tail] = val
	q.tail++
	q.size++
}

func (q *Queue) Dequeue() interface{} {
	if q.Empty() {
		panic("empty")
	}

	if q.head == q.capasity {
		q.head = 0
	}

	val := q.data[q.head]
	q.head++
	q.size--
	return val
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Full() bool {
	return q.size == q.capasity
}
