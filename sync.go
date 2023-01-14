package pool

import "sync"

type SyncPool[T any] struct {
	pool sync.Pool
}

func NewSync[T any]() *SyncPool[T] {
	return &SyncPool[T]{
		pool: sync.Pool{
			New: func() interface{} {
				return new(T)
			},
		},
	}
}

func (p *SyncPool[T]) Get() *T {
	return p.pool.Get().(*T)
}

func (p *SyncPool[T]) Put(x *T) {
	p.pool.Put(x)
}
