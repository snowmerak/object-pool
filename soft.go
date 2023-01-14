package pool

import "sync"

type SoftPool[T any] struct {
	sizedPool chan *T
	syncPool  sync.Pool
}

func NewSoft[T any](size int) *SoftPool[T] {
	return &SoftPool[T]{
		sizedPool: make(chan *T, size),
		syncPool: sync.Pool{
			New: func() interface{} {
				return new(T)
			},
		},
	}
}

func (p *SoftPool[T]) Get() *T {
	select {
	case x := <-p.sizedPool:
		return x
	default:
		return p.syncPool.Get().(*T)
	}
}

func (p *SoftPool[T]) Put(x *T) {
	select {
	case p.sizedPool <- x:
	default:
		p.syncPool.Put(x)
	}
}
