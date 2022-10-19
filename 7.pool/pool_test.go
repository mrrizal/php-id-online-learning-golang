// source: https://github.com/slaise/GoJsonBenchmark/blob/main/benchmark_syncpool_test.go
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

// var buf, _ = json.Marshal(User{Name: "Test", ID: 1})

var userPool = sync.Pool{
	New: func() interface{} {
		return &User{Name: "Test", ID: 1}
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		user := &User{}
		data := fmt.Sprintf("{'name': 'Test', 'id': '%d'}", n)
		json.Unmarshal([]byte(data), user)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		user := userPool.Get().(*User)
		data := fmt.Sprintf("{'name': 'Test', 'id': '%d'}", n)
		json.Unmarshal([]byte(data), user)
		userPool.Put(user)
	}
}
