package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

// Key in this struct allows us to delete values from map with
// lower alghoritm complexity - O(1).
// Without this it is necessary to make traversing through
// the map every time when we want to delete a key.
type valueToMapLink struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	element, existsFlag := l.items[key]
	if existsFlag {
		l.queue.Remove(element)
	} else {
		if l.queue.Len() == l.capacity {
			delete(l.items, l.queue.Back().Value.(valueToMapLink).key)
			l.queue.Remove(l.queue.Back())
		}
	}
	l.items[key] = l.queue.PushFront(valueToMapLink{key, value})
	return existsFlag
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	element, existsFlag := l.items[key]
	if existsFlag {
		l.queue.MoveToFront(element)
		return l.items[key].Value.(valueToMapLink).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
