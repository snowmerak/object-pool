package pool

import "sync"

type Pool[T any] struct {
	pool sync.Pool
}

func New[T any]() *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() interface{} {
				return new(T)
			},
		},
	}
}

func (p *Pool[T]) Get() *T {
	return p.pool.Get().(*T)
}

func (p *Pool[T]) Put(x *T) {
	p.pool.Put(x)
}
