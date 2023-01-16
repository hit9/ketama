Ketama
======

Package ketama implements a consistent hashing ring (on md5).

https://pkg.go.dev/github.com/hit9/ketama

Example
-------

```go
package main

import (
	"fmt"
	"github.com/hit9/ketama"
)

func main() {
	ring := ketama.NewRing([]*ketama.Node{
		ketama.NewNode("127.0.0.1:8000", "binding data0", 1),
		ketama.NewNode("127.0.0.1:8001", "binding data1", 1),
		ketama.NewNode("127.0.0.1:8002", "binding data2", 1),
		ketama.NewNode("127.0.0.1:8003", "binding data3", 1),
		ketama.NewNode("127.0.0.1:8004", "binding data3", 1),
	})
	fmt.Printf("%v\n", ring.Get("key1"))
}
```

Please checkout [ketama_example.go](ketama_example.go) for more .

Links
-----

- [一致性哈希算法 - 哈希环法](https://writings.sh/post/consistent-hashing-algorithms-part-2-consistent-hash-ring)

License
-------

BSD.
