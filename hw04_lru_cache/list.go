package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) PushFront(v interface{}) *ListItem {
	newValue := ListItem{v, l.front, nil}
	if l.len > 0 {
		l.front.Prev = &newValue
	} else {
		l.back = &newValue
	}
	l.front = &newValue
	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newValue := ListItem{v, nil, l.back}
	if l.len > 0 {
		l.back.Next = &newValue
	} else {
		l.front = &newValue
	}
	l.back = &newValue
	l.len++
	return l.back
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if l.len == 0 {
		panic("Attempt to remove element from empty list")
	}
	if i.Next != nil && i.Prev != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	} else if l.front == i && l.len > 1 {
		l.front.Next.Prev = nil
		l.front = l.front.Next
	}
	if l.back == i && l.len > 1 {
		l.back.Prev.Next = nil
		l.back = l.back.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Len() int {
	return l.len
}

func NewList() List {
	return new(list)
}
