package queue

//Использование строго типизированной очереди обусловлено отсутствием
//необходимости наличия нескольких типов в очереди.

type item[T comparable] struct {
	next *item[T]
	data T
}

type Queuer[T comparable] interface {
	Push(data T)
	Pop() T
	Peek() T
	IsEmpty() bool
	Len() int
}

type Queue[T comparable] struct {
	head *item[T]
	tail *item[T]
	len  int
}

func (q *Queue[T]) Push(data T) {
	newItem := &item[T]{data: data}
	switch q.head {
	case nil:
		q.head = newItem
		q.tail = newItem
	default:
		q.tail.next = newItem
		q.tail = newItem
	}
	q.len++
}

func (q *Queue[T]) Pop() T {
	var dv T
	if q.head == nil {
		return dv
	}
	v := q.head.data
	q.head = q.head.next
	q.len--
	return v
}

func (q *Queue[T]) Peek() T {
	if q.head == nil {
		var dv T
		return dv
	}
	return q.head.data
}

func (q *Queue[T]) Get(index int) T {
	if index < 0 || index >= q.len {
		var dv T
		return dv
	}

	current := q.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current.data
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Len() int {
	return q.len
}

func (q *Queue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.len = 0
}

func New[T comparable]() Queue[T] {
	return Queue[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}
