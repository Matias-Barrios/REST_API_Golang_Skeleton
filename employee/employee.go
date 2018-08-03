package employee

type Employee struct {
	Id       int
	Name     string
	Lastname string
	Age      int
	Phones   []string
	Email    string
}

func (t Employee) getName() string {
	return t.Name
}
