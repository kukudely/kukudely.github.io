package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"reflect"
	"testing"
)

func equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	equal(t, c.Get(1), 1)
	c.Put(3, 3)
	equal(t, c.Get(2), -1)
	c.Put(4, 4)
	equal(t, c.Get(1), -1)
	equal(t, c.Get(3), 3)
	equal(t, c.Get(4), 4)
	rng := rand.New(rand.NewSource(146))
	for capacity := 0; capacity <= 8; capacity++ {
		cache := Constructor(capacity)
		entries := map[int]int{}
		last := map[int]int{}
		clock := 0
		for operation := 0; operation < 2000; operation++ {
			key := rng.Intn(15)
			clock++
			if rng.Intn(2) == 0 {
				want := -1
				if value, ok := entries[key]; ok {
					want = value
					last[key] = clock
				}
				equal(t, cache.Get(key), want)
			} else {
				value := rng.Intn(100)
				cache.Put(key, value)
				if capacity == 0 {
					continue
				}
				entries[key] = value
				last[key] = clock
				if len(entries) > capacity {
					victim, minTime := 0, clock+1
					for k, when := range last {
						if when < minTime {
							victim, minTime = k, when
						}
					}
					delete(entries, victim)
					delete(last, victim)
				}
			}
		}
	}
}
