# pool

## install

`go get github.com/snowmerak/object-pool`

## syncpool

`SyncPool` is a pool wrapper for `sync.Pool`. It is safe for concurrent use.

```go
package main

import (
    "fmt"
    pool "github.com/snowmerak/object-pool"
)

func main() {
    p := pool.NewSync[int]()

    intPtr := p.Get()
    *intPtr = 1

    p.Put(intPtr)
}
```

## sizedpool

`SizedPool` is a pool with a fixed size. 

If `block` is true, `Get` will return nil if no object is available.

```go
package main

import (
    "fmt"
    pool "github.com/snowmerak/object-pool"
)

func main() {
    p := pool.NewSized[int](10, true)

    intPtr := p.Get()
    *intPtr = 1

    p.Put(intPtr)
}
```

## softpool

`SoftPool` is a pool with a fixed size and `sync.Pool`.

```go
package main

import (
    "fmt"
    pool "github.com/snowmerak/object-pool"
)

func main() {
    p := pool.NewSoft[int](10)

    intPtr := p.Get()
    *intPtr = 1

    p.Put(intPtr)
}
```
