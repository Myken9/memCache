package inmemory

import (
	"sync"
	"testing"
)

func TestStorage_Set(t *testing.T) {
	st := Storage{}

	t.Run("Currency set", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				st.Set("555", "444")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				st.Set("111", "333")
			}
		}()
		wg.Wait()
	})
}

func TestStorage_Delete(t *testing.T) {
	st := Storage{}

	t.Run("Currency delete", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				st.Delete("555")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				st.Delete("555")
			}
		}()
		wg.Wait()
	})
}

func TestStorage_Get(t *testing.T) {
	st := Storage{}

	t.Run("Currency get", func(t *testing.T) {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				st.Set("555", "333")
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				st.Get("555")
			}
		}()
		wg.Wait()
	})
}
