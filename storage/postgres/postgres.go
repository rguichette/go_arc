package postgres

import "github.com/rguichette/go_arc"

type Postg map[int]go_arc.Person

func (m Postg) Save(n int, p go_arc.Person) {
	m[n] = p
}

func (m Postg) Retrieve(n int) go_arc.Person {
	return m[n]
}
