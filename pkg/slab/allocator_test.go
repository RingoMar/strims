package slab

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/MemeLabs/go-ppspp/pkg/mathutil"
	"github.com/MemeLabs/go-ppspp/pkg/sortutil"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

type Meme struct {
	foo int64
	bar [16]byte
}

func TestAlloc(t *testing.T) {
	s := NewWithSize[Meme](16)
	var ms []*Meme
	for i := 0; i < 65000; i++ {
		ms = append(ms, s.Alloc())
	}
	for _, m := range ms {
		s.Free(m)
	}
}

func BenchmarkAlloc(b *testing.B) {
	ms := make([]*Meme, 0, b.N)

	b.ResetTimer()
	s := NewWithSize[Meme](1024)

	for i := 0; i < b.N; i += 65000 {
		n := mathutil.Min(b.N-i, 65000)

		for j := 0; j < n; j++ {
			ms = append(ms, s.Alloc())
		}

		for j := 0; j < n; j++ {
			s.Free(ms[j])
			ms[j] = nil
		}
		ms = ms[:0]
	}
}

func BenchmarkNativeAlloc(b *testing.B) {
	ms := make([]*Meme, 0, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i += 65000 {
		n := mathutil.Min(b.N-i, 65000)

		for j := 0; j < n; j++ {
			ms = append(ms, &Meme{})
		}

		for j := 0; j < n; j++ {
			ms[j] = nil
		}
		ms = ms[:0]
	}
}

func BenchmarkSyncPool(b *testing.B) {
	ms := make([]*Meme, 0, b.N)

	b.ResetTimer()
	p := sync.Pool{
		New: func() interface{} {
			return &Meme{}
		},
	}

	for i := 0; i < b.N; i += 65000 {
		n := mathutil.Min(b.N-i, 65000)

		for j := 0; j < n; j++ {
			ms = append(ms, p.Get().(*Meme))
		}

		for j := 0; j < n; j++ {
			p.Put(ms[j])
			ms[j] = nil
		}
		ms = ms[:0]
	}
}

func TestListGrow(t *testing.T) {
	cases := []struct {
		count int
	}{
		{100},
		{10000},
		{1000000},
	}
	for _, c := range cases {
		ms := make([]ref, 0, 4096)
		s := newFreeList(4096)

		for i := 0; i < c.count; i++ {
			if len(ms) > 0 && rand.Float64() < 0.5 {
				n := rand.Intn(len(ms))
				s.Free(ms[n])
				ms[n] = ms[len(ms)-1]
				ms = ms[:len(ms)-1]
			} else {
				n := s.Alloc()
				if n != nilRef {
					ms = append(ms, n)
				}
			}
		}

		sortutil.Ordered(ms)
		assert.Equal(t, len(slices.Compact(ms)), len(ms))
	}
}

func BenchmarkListSerial(b *testing.B) {
	ms := make([]ref, 0, b.N)
	s := newFreeList(65000)

	b.ResetTimer()

	for i := 0; i < b.N; i += 65000 {
		n := mathutil.Min(b.N-i, 65000)

		for j := 0; j < n; j++ {
			ms = append(ms, s.Alloc())
		}

		for j := 0; j < n; j++ {
			s.Free(ms[j])
		}
		ms = ms[:0]
	}
}

func BenchmarkListRandom(b *testing.B) {
	ms := make([]ref, 0, b.N)
	s := newFreeList(65000)

	n := mathutil.Min(b.N, 65000)
	is := make([]int, n)
	for i := range is {
		is[i] = i
	}
	rng := rand.New(rand.NewSource(1234))
	d := b.N % n
	dis := is[d:]
	rng.Shuffle(n-d, func(i, j int) { dis[i], dis[j] = dis[j], dis[i] })
	rng.Shuffle(d, func(i, j int) { is[i], is[j] = is[j], is[i] })

	b.ResetTimer()

	for i := 0; i < b.N; i += 65000 {
		n := mathutil.Min(b.N-i, 65000)

		for j := 0; j < n; j++ {
			ms = append(ms, s.Alloc())
		}

		for j := 0; j < n; j++ {
			s.Free(ms[is[j]])
		}
		ms = ms[:0]
	}
}
