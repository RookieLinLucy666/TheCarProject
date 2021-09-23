package main

type Item interface {}

// 队列结构体
type Queue struct {
	Items []Item
}

type IQueue interface {
	New() Queue
	Enqueue(t Item)
	Dequeue(t Item)
	IsEmpty() bool
	Size() int
}

// 新建
func (q *Queue)New() *Queue  {
	q.Items = []Item{}
	return q
}

// 入队
func (q *Queue) Enqueue(data Item)  {
	q.Items = append(q.Items, data)
}

// 出队
func (q *Queue) Dequeue() *Item {
	// 由于是先进先出，0为队首
	item := q.Items[0]
	q.Items = q.Items[1: len(q.Items)]
	return &item
}

// 队列是否为空
func (q *Queue) IsEmpty() bool  {
	return len(q.Items) == 0
}

// 队列长度
func (q *Queue) Size() int  {
	return len(q.Items)
}

var q Queue

func initQueue() *Queue  {
	if q.Items == nil{
		q = Queue{}
		q.New()
	}
	return &q
}

