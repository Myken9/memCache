package inmemory

import (
	"github.com/bradfitz/gomemcache/memcache"
	"sync"
	"testing"
)

func TestStorage_Set(t *testing.T) {
	mc := memcache.New("localhost:11211")
	ns := NewStorage(mc)

	t.Run("Currency set", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ns.Set("555", "444")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ns.Set("111", "333")
			}
		}()
		wg.Wait()
	})
}

func TestStorage_Delete(t *testing.T) {
	mc := memcache.New("localhost:11211")
	ns := NewStorage(mc)

	t.Run("Currency delete", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ns.Delete("555")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ns.Delete("555")
			}
		}()
		wg.Wait()
	})
}

func TestStorage_Get(t *testing.T) {
	mc := memcache.New("localhost:11211")
	ns := NewStorage(mc)

	t.Run("Currency get", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ns.Set("555", "333")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				ns.Get("555")
			}
		}()
		wg.Wait()
	})
}
