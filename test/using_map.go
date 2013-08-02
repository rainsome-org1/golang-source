package main
import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo
	personDB = make(map[string] PersonInfo)

	personDB["1"] = PersonInfo{"12345", "Smith", "Xi'an"}
	personDB["2"] = PersonInfo{"12346", "Bob", "Beijing"}

	person, ok := personDB["2"]

	if ok {
		fmt.Println("Find the person:", person.Name, "lives in", person.Address)
	} else {
		fmt.Println("Not find the perspn 2.")
	}

	// delete the id "2"
	delete(personDB, "2")
	person1, ok1 := personDB["2"]

	if ok1 {
		fmt.Println("Find the person:", person1.Name, "lives in", person1.Address)
	} else {
		fmt.Println("Not find the perspn 2.")
	}


}
