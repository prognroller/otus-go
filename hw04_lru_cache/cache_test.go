package hw04_lru_cache //nolint:golint,stylecheck

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

	t.Run("purge set logic", func(t *testing.T) {
		tests := []struct {
			key Key
			val interface{}
		}{
			{key: "a", val: 1},
			{key: "b", val: 2},
			{key: "c", val: 3},
			{key: "d", val: 4},
		}

		cache := NewCache(3)

		for _, pair := range tests {
			cache.Set(pair.key, pair.val)
		}

		for i := 3; i > 0; i-- {
			_, ok := cache.Get(tests[i].key)
			require.True(t, ok)
		}

		_, ok := cache.Get("a")
		require.False(t, ok)
	})

	t.Run("purge set/get logic", func(t *testing.T) {
		tests := []struct {
			key Key
			val interface{}
		}{
			{key: "a", val: 1},
			{key: "b", val: 2},
			{key: "c", val: 3},
		}
		expectedKeys := []Key{"a", "c", "d"}

		cache := NewCache(3)

		for i := 0; i < len(tests); i++ {
			cache.Set(tests[i].key, tests[i].val)
		}

		cache.Get("a")
		cache.Get("b")
		cache.Set("c", 33)
		cache.Get("a")

		cache.Set("d", 4)

		for _, k := range expectedKeys {
			_, ok := cache.Get(k)
			require.True(t, ok)
		}

		_, ok := cache.Get("b")
		require.False(t, ok)
	})

	t.Run("Clear", func(t *testing.T) {
		tests := []struct {
			key Key
			val interface{}
		}{
			{key: "a", val: 1},
			{key: "b", val: 2},
		}
		c := NewCache(2)

		for _, k := range tests {
			c.Set(k.key, k.val)
		}

		c.Clear()

		for _, k := range tests {
			_, ok := c.Get(k.key)
			require.False(t, ok)
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // NeedRemove if task with asterisk completed

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
