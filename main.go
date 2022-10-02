package main

type person struct {
	first string
}

type mongo map[int]person
type postg map[int]person

func (m mongo) save(n int, p person) {
	m[n] = p
}

func (m mongo) retrieve(n int) person {
	return m[n]
}
func (pg postg) save(n int, p person) {
	pg[n] = p
}

func (pg postg) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(a int) person
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}
func get(a accessor, n int) person {
	return a.retrieve(n)
}

func main() {
	dbm := mongo{}
	dbp := postg{}

	p := person{
		first: "Ralph",
	}

	put(dbm, 1, p)

	get(dbp, 2)

}
