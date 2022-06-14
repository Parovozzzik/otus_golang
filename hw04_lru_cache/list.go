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
	item ListItem
	size int
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.item.Prev
}

func (l list) Back() *ListItem {
	return l.item.Next
}

func (l *list) PushFront(v interface{}) *ListItem {
	var newItem ListItem
	if l.size == 0 {
		newItem.Value = v
		newItem.Prev = nil
		newItem.Next = &l.item
		l.item.Next = &newItem
	} else {
		firstItem := l.Front()
		newItem.Value = v
		newItem.Prev = nil
		newItem.Next = firstItem
		firstItem.Prev = &newItem
		l.Back().Next = nil
	}
	l.item.Prev = &newItem
	l.size++
	return &newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	var newItem ListItem
	newItem.Value = v
	newItem.Prev = l.Back()
	newItem.Next = nil
	l.Back().Next = &newItem
	l.item.Next = &newItem
	l.size++
	return &newItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if l.Back() == i {
		l.item.Next = i.Prev
	}
	if l.Front() == i {
		l.item.Prev = i.Next
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.item.Next = i.Prev
	}
	i.Prev = nil
	i.Next = l.Front()
	i.Next.Prev = i
	l.item.Prev = i
}
