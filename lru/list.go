package lru

type ListNode struct {
	prev *ListNode
	next *ListNode
	Val  interface{}
}

type Link struct {
	root *ListNode
	n    int
}

func (list *Link) Len() int {
	return list.n
}
func NewList() *Link {
	var rt = NewNode()
	rt.next = rt
	rt.prev = rt
	return &Link{
		root: rt,
		n:    0,
	}
}
func NewNode() *ListNode {
	return &ListNode{}
}
func (list *Link) addSize(x int) {
	list.n += x
}
func (list *Link) Push_back(val interface{}) {
	rt := NewNode()
	rt.Val = val
	list.push_back(rt)
}
func (list *Link) Push_front(val interface{}) {
	rt := NewNode()
	rt.Val = val
	list.push_front(rt)
}

// insert rhs, lhs->rhs
func (list *Link) linkNode(lhs, rhs *ListNode) {
	lhs.next = rhs
	rhs.prev = lhs
}
func (list *Link) push_back(rt *ListNode) {
	list.linkNode(list.Back_pointer(), rt)
	list.linkNode(rt, list.root)
	list.addSize(1)
}
func (list *Link) push_front(rt *ListNode) {
	list.linkNode(rt, list.Front_pointer())
	list.linkNode(list.root, rt)
	list.addSize(1)
}
func (list *Link) Front_value() interface{} {
	return list.root.next.Val
}
func (list *Link) Back_value() interface{} {
	return list.root.prev.Val
}
func (list *Link) Front_pointer() *ListNode {
	return list.root.next
}

func (list *Link) Back_pointer() *ListNode {
	return list.root.prev
}
func (list *Link) remove(rt *ListNode) {
	list.linkNode(rt.prev, rt.next)
	list.addSize(-1)
	rt.next = nil
	rt.prev = nil
}
func (list *Link) Erase(rt *ListNode) {
	list.remove(rt)
}
func (list *Link) Pop_back() {
	list.remove(list.root.prev)
}
func (list *Link) Pop_front() {
	list.remove(list.root.next)
}
func (list *Link) MoveToFront(rt *ListNode) {
	list.remove(rt)
	list.push_front(rt)
}
func (list *Link) MoveToBack(rt *ListNode) {
	list.remove(rt)
	list.push_back(rt)
}
