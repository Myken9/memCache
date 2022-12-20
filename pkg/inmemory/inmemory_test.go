package inmemory

import (
	"memcach/pkg/memcache"
	"sync"
	"testing"
)

func TestStorage_Set(t *testing.T) {
	mc, _ := memcache.New("localhost:11211")
	defer mc.Close()
	ns := NewStorage(mc)

	t.Run("Currency set", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				err := ns.Set("555", "444")
				if err != nil {
					panic(err)
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				err := ns.Set("111", "333")
				if err != nil {
					panic(err)
				}
			}
		}()
		wg.Wait()
	})
}

func TestStorage_Delete(t *testing.T) {
	mc, _ := memcache.New("localhost:11211")
	defer mc.Close()
	ns := NewStorage(mc)

	t.Run("Currency delete", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				err := ns.Delete("555")
				if err != nil {
					panic(err)
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				err := ns.Delete("555")
				if err != nil {
					panic(err)
				}
			}
		}()
		wg.Wait()
	})
}

func TestStorage_Get(t *testing.T) {
	mc, _ := memcache.New("localhost:11211")
	defer mc.Close()
	ns := NewStorage(mc)

	err := ns.Set("asd", "asd")
	if err != nil {
		panic(err)
	}

	t.Run("Currency get", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				err := ns.Set("555", "333")
				if err != nil {
					panic(err)
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				_, err := ns.Get("asd")
				if err != nil {
					panic(err)
				}
			}
		}()
		wg.Wait()
	})
}
