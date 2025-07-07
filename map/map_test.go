package mymap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	m := New(50)

	fmt.Println(m.Get(5))
	fmt.Println(m.Put(5, "hello"))
	fmt.Println(m.Get(5))
	fmt.Println(m.Put(5, "world"))
	fmt.Println(m.Get(5))
	fmt.Println(m.Del(5))
	fmt.Println(m.Get(5))

	fmt.Println(m.Put(6, "sad"))
	fmt.Println(m.Put(7, "xxxj"))
	fmt.Println(m.Put(3, "aasss"))

	fmt.Println(m)
}
