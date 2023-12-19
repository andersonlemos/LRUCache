package LRUCache

type Item struct {
	key   string
	value any
	next  *Item
	prev  *Item
}

type Cache struct {
	dict     map[string]*Item
	items    *Item
	head     *Item
	capacity int
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		dict:     make(map[string]*Item),
	}
}

func (c *Cache) Get(key string) any {
	item, ok := c.dict[key]

	if !ok {
		return -1
	}

	if item == c.head {
		return item.value
	}

	if item.prev != nil {
		item.prev.next = item.next
	}

	if item.next != nil {
		item.next.prev = item.prev
	}

	if item == c.items {
		c.items = item.next
	}

	c.head.next = item
	item.prev = c.head
	c.head = item
	c.head.next = nil

	return item.value
}

func (c *Cache) Set(key string, value any) {
	if item, ok := c.dict[key]; ok {
		// remove old item with the same key
		if item.prev != nil {
			item.prev.next = item.next
		}

		if item.next != nil {
			item.next.prev = item.prev
		}

		if item == c.items {
			c.items = c.items.next
		}

		if item == c.head {
			c.head = c.head.prev
		}
		delete(c.dict, key)
	}

	if len(c.dict) == c.capacity {
		leastItem := c.items

		delete(c.dict, leastItem.key)
		c.items.key = key
		c.items.value = value
		c.dict[key] = c.items
		return
	}

	newItem := &Item{key: key, value: value}
	newItem.next = c.items
	if c.items != nil {
		c.items.prev = newItem
	}
	c.items = newItem
	c.dict[key] = newItem

	if c.items.next == nil {
		c.head = c.items
	}

}
