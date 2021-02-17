package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, val interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, val interface{}) bool {
	item := cacheItem{key: key, value: val}

	if _, ok := c.items[key]; ok {
		c.items[key].Value = item
		c.queue.MoveToFront(c.items[key])

		return true
	}

	if c.capacity == c.queue.Len() {
		back := c.queue.Back()

		c.queue.Remove(back)
		delete(c.items, back.Value.(cacheItem).key)
	}

	c.items[key] = c.queue.PushFront(item)

	return false
}

func (c lruCache) Get(key Key) (interface{}, bool) {
	if el, ok := c.items[key]; ok {
		c.queue.MoveToFront(c.items[key])

		return el.Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
