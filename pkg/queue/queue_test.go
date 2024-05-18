package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("TestNewQueue", TestNewQueue)
	t.Run("TestQueue_PushAndPop", TestPushAndPop)
	t.Run("TestQueue_Peek", TestPeek)
	t.Run("TestQueue_IsEmpty", TestIsEmpty)
	t.Run("TestQueue_Len", TestLen)
	t.Run("TestQueue_Clear", TestClear)
}

func TestNewQueue(t *testing.T) {
	q := New[int]()
	if q.Len() != 0 {
		t.Errorf("Expected length of queue to be 0, got %d", q.Len())
	}
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
	if q.Peek() != 0 {
		t.Errorf("Expected Peek() to return 0, got %d", q.Peek())
	}
}

func TestPushAndPop(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Len() != 3 {
		t.Errorf("Expected length of queue to be 3, got %d", q.Len())
	}

	if q.Peek() != 1 {
		t.Errorf("Expected Peek() to return 1, got %d", q.Peek())
	}

	if popped := q.Pop(); popped != 1 {
		t.Errorf("Expected Pop() to return 1, got %d", popped)
	}

	if q.Len() != 2 {
		t.Errorf("Expected length of queue to be 2 after Pop(), got %d", q.Len())
	}

	if q.Peek() != 2 {
		t.Errorf("Expected Peek() to return 2 after Pop(), got %d", q.Peek())
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)

	if peeked := q.Peek(); peeked != 1 {
		t.Errorf("Expected Peek() to return 1, got %d", peeked)
	}

	q.Pop()

	if peeked := q.Peek(); peeked != 2 {
		t.Errorf("Expected Peek() to return 2 after Pop(), got %d", peeked)
	}
}

func TestIsEmpty(t *testing.T) {
	q := New[int]()
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty")
	}

	q.Push(1)
	if q.IsEmpty() {
		t.Error("Expected queue not to be empty after Push()")
	}

	q.Pop()
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty after Pop()")
	}
}

func TestLen(t *testing.T) {
	q := New[int]()
	if q.Len() != 0 {
		t.Errorf("Expected length of queue to be 0, got %d", q.Len())
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Len() != 3 {
		t.Errorf("Expected length of queue to be 3, got %d", q.Len())
	}

	q.Pop()

	if q.Len() != 2 {
		t.Errorf("Expected length of queue to be 2 after Pop(), got %d", q.Len())
	}
}

func TestClear(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Clear()

	if q.Len() != 0 {
		t.Errorf("Expected length of queue to be 0 after Clear(), got %d", q.Len())
	}

	if !q.IsEmpty() {
		t.Error("Expected queue to be empty after Clear()")
	}
}
