package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(interface{}) *ListItem
	PushBack(interface{}) *ListItem
	Remove(el *ListItem)
	MoveToFront(el *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(i interface{}) *ListItem {
	newFront := ListItem{Value: i, Next: l.Front()}

	if l.Front() != nil {
		l.Front().Prev = &newFront
	}

	if l.len == 0 {
		l.back = &newFront
	}

	l.front = &newFront
	l.len++

	return l.front
}

func (l *list) MoveToFront(el *ListItem) {
	if l.front == el {
		return
	}

	l.resetLinks(el)

	l.front.Prev = el
	el.Next = l.front
	l.front = el
}

func (l *list) PushBack(i interface{}) *ListItem {
	newBack := ListItem{Prev: l.Back(), Value: i}

	if l.Back() != nil {
		l.Back().Next = &newBack
	}

	if l.len == 0 {
		l.front = &newBack //TODO: cover test
	}

	l.back = &newBack
	l.len++

	return l.back
}

func (l *list) Remove(delEl *ListItem) {
	l.resetLinks(delEl)

	l.len--
}

func (l *list) resetLinks(el *ListItem) {
	if l.back == el {
		l.back = el.Prev
	}

	if el.Prev != nil {
		el.Prev.Next = el.Next
	}

	if el.Next != nil {
		el.Next.Prev = el.Prev
	}

	el.Prev = nil
	el.Next = nil
}
