package demolib

type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}

func NewPerson(personName string) Person {
	return Person{
		name: personName,
	}
}
