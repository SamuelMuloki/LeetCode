package methods

type Stack interface {
	IsEmpty() bool
	Push(athlete Person) int
	Pop() Person
	Length() int
}

type Person struct {
	Name string
	Age  int
}

type People []Person

func New() Stack {
	return &People{}
}

func (p *People) IsEmpty() bool {
	return len(*p) == 0
}

func (p *People) Push(athlete Person) int {
	*p = append(*p, athlete)

	return len(*p)
}

func (p *People) Pop() Person {
	if p.IsEmpty() {
		return Person{}
	}

	lastIndex := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]

	return lastIndex
}

func (p *People) Length() int {
	return len(*p)
}
