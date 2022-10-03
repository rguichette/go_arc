package go_arc

import (
	"testing"
)

type Db map[int]Person

func (m Db) Save(n int, p Person) {
	m[n] = p
}
func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	mdb := Db{}
	p := Person{First: "John"}
	Put(mdb, 1, p)

	got := mdb.Retrieve(1)
	if mdb.Retrieve(1) != p {
		t.Fatalf("want %v, got %v ", p, got)
	}
}
