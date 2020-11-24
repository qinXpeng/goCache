package lru

type KLfunc func(key string, value ValueInterface)

type LRUCache struct {
	maxBytes  int64
	nBytes    int64
	list      *Link
	cache     map[string]*ListNode
	onEvicted KLfunc
}

type ValueInterface interface {
	Len() int
}

type keyNode struct {
	key   string
	value ValueInterface
}

func NewCache(maxBytes int64, onEvicted KLfunc) *LRUCache {
	return &LRUCache{
		maxBytes:  maxBytes,
		list:      NewList(),
		cache:     make(map[string]*ListNode),
		onEvicted: onEvicted,
	}
}

func (c *LRUCache) Insert(key string, value ValueInterface) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		kv := ele.Val.(*keyNode)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		c.list.Push_front(&keyNode{key, value})
		c.cache[key] = c.list.Front_pointer()
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.Erase()
	}
}
func (c *LRUCache) Erase() {
	ele := c.list.Back_pointer()
	if ele != nil {
		c.list.Pop_back()
		kv := ele.Val.(*keyNode)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}
func (c *LRUCache) Get(key string) (ValueInterface, bool) {
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		kv := ele.Val.(*keyNode)	
		val := kv.value
		return val, ok
	}
	return nil, false
}
