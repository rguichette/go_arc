package mongo

import "github.com/rguichette/go_arc"

type Mongo map[int]go_arc.Person

func (m Mongo) Save(n int, p go_arc.Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) go_arc.Person {
	return m[n]
}
