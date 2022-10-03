package go_arc

//Person is how the arc package stores a Person
type Person struct {
	First string
}

//Accessor is how to store or retrieve a person
//When retriving a person , if they do not exist, return the zero value
type Accessor interface {
	Save(n int, p Person)
	Retrieve(a int) Person
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}
func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}

// func main() {
// 	// dbm := Mongo{}
// 	// dbp := Postg{}

// 	// p := Person{
// 	// 	first: "Ralph",
// 	// }

// 	// put(dbm, 1, p)

// 	// get(dbp, 2)

// }
