package pool

type SizedPool[T any] struct {
	pool  chan *T
	block bool
}

func NewSized[T any](size int, block bool) *SizedPool[T] {
	return &SizedPool[T]{
		pool:  make(chan *T, size),
		block: block,
	}
}

func (p *SizedPool[T]) Get() *T {
	select {
	case x := <-p.pool:
		return x
	default:
		if p.block {
			return nil
		}
		return new(T)
	}
}

func (p *SizedPool[T]) Put(x *T) {
	select {
	case p.pool <- x:
	default:
		x = nil
	}
}
