// Copyright 2016 Chao Wang <hit9@icloud.com>

package ketama

import (
	"math/rand"
	"runtime"
	"testing"
)

// Must asserts the given value is True for testing.
func Must(t *testing.T, v bool) {
	if !v {
		_, fileName, line, _ := runtime.Caller(1)
		t.Errorf("\n unexcepted: %s:%d", fileName, line)
	}
}

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandString returns a random string with the fixed length.
func RandString(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		j := rand.Intn(len(letters))
		b[i] = letters[j]
	}
	return string(b)
}

func TestBalance(t *testing.T) {
	nodes := []*Node{
		NewNode("127.0.0.1:8000", nil, 1),
		NewNode("127.0.0.1:8001", nil, 1),
		NewNode("127.0.0.1:8002", nil, 1),
		NewNode("127.0.0.1:8003", nil, 1),
		NewNode("127.0.0.1:8004", nil, 1),
		NewNode("127.0.0.1:8005", nil, 1),
		NewNode("127.0.0.1:8006", nil, 1),
		NewNode("127.0.0.1:8007", nil, 1),
		NewNode("127.0.0.1:8008", nil, 1),
		NewNode("127.0.0.1:8009", nil, 1),
		NewNode("127.0.0.1:8010", nil, 1),
		NewNode("127.0.0.1:8011", nil, 1),
		NewNode("127.0.0.1:8012", nil, 1),
	}
	ring := NewRing(nodes)
	Must(t, len(ring.nodes) == len(nodes)*160)
	N := 4096 * len(nodes)
	m := make(map[string]int, 0)
	for i := 0; i < N; i++ {
		key := RandString(128)
		n := ring.Get(key)
		m[n.Key()]++
	}
	for _, v := range m {
		// rate 0.8 ~ 1.2
		Must(t, float64(v) > float64(N/len(nodes))*0.8)
		Must(t, float64(v) < float64(N/len(nodes))*1.2)
	}
}

func TestConsistence(t *testing.T) {
	nodes := []*Node{
		NewNode("192.168.0.1:9527", nil, 1),
		NewNode("192.168.0.2:9527", nil, 1),
		NewNode("192.168.0.3:9527", nil, 2),
		NewNode("192.168.0.4:9527", nil, 2),
		NewNode("192.168.0.5:9527", nil, 4),
	}
	ring := NewRing(nodes)
	Must(t, len(ring.nodes) == (1+1+2+2+4)*160)
	for i := 0; i < 1024; i++ {
		key := RandString(128)
		n1 := ring.Get(key)
		n2 := ring.Get(key)
		Must(t, n1.Key() == n2.Key())
	}
}
