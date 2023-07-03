package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple cache", func(t *testing.T) {
		c := NewCache(10)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)
	})

	t.Run("clear cache", func(t *testing.T) {
		c := NewCache(10)

		c.Set("aaa", 100)
		c.Set("bbb", 100)
		_, existsFlag := c.Get("bbb")
		require.True(t, existsFlag)
		c.Clear()
		_, existsFlag = c.Get("bbb")
		require.False(t, existsFlag)
		_, existsFlag = c.Get("aaa")
		require.False(t, existsFlag)
	})

	t.Run("remove oldest item", func(t *testing.T) {
		c := NewCache(5)
		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)
		c.Set("d", 400)
		c.Set("e", 500)
		c.Set("f", 600)
		_, existsFlag := c.Get("a")
		require.False(t, existsFlag)
		_, existsFlag = c.Get("f")
		require.True(t, existsFlag)
	})

	t.Run("remove oldest item with get", func(t *testing.T) {
		c := NewCache(5)
		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)
		c.Set("d", 400)
		c.Set("e", 500)
		_, existsFlag := c.Get("a") // make "b" oldest value
		require.True(t, existsFlag)
		c.Set("f", 600)
		_, existsFlag = c.Get("f")
		require.True(t, existsFlag)
		_, existsFlag = c.Get("a")
		require.True(t, existsFlag)
		_, existsFlag = c.Get("b")
		require.False(t, existsFlag)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
