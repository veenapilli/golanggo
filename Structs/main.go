package main

import "fmt"

/**
Passed as a value (Value type)
Known at compile time
Can be formed with multiple type
Fields cant be iterated upon like a map
Represents a single object with multiple properties
*/
type person struct {
	firstName  string
	middleName string
	lastName   string
	contactInfo
}

type contactInfo struct {
	email    string
	postcode int
}

func (p person) describe() {
	fmt.Printf("\n%+v\n", p)
}

/**
without pointers, update works as call by value not reference

(p *person) -> is the type description. pointer to an object
(*p) -> operator. Manipulate the referenced value

&p -> converts value to address
*p -> converts address to value
*/
func (p *person) updateName(s string) {
	(*p).firstName = s
}

func main() {
	// populated person object
	p := person{
		firstName: "a",
		lastName:  "b",
		contactInfo: contactInfo{
			email:    "a@b.com", // need the comma at the end [if struct is defined on multiple lines]
			postcode: 1234,
		},
	}
	fmt.Println(p)
	p.updateName("c")
	p.describe()
	// empty person object
	var q person
	q.describe()

	fmt.Println("\n------- slices -------")
	/**
	Working with slices
	We dont need pointers here !
	We pass them as variables and not as receivers
	*/
	sI := []int{1, 2, 3}
	sS := []string{"One", "Two", "Three"}

	update(sI)
	updateStr(sS)

	fmt.Println(sI)
	fmt.Println(sS)

}

func update(i []int) {
	i[0] = 4
}

func updateStr(s []string) {
	s[0] = "Four"
}
