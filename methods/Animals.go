package methods

// Animal interface is satisfied if a type defines its methods
// There is no implements keyword in go whether or not a type satisfies an interface is determined automatically
type Animal interface {
	Speak() string
}

// Dog defines actions for dogs
type Dog struct {
}

// Speak method is from interface type animal and is implemented by type Dog
func (d Dog) Speak() string {
	return "Woof"
}

// Cat defines actions for cats
type Cat struct {
}

// Speak method is from interface type animal and is implemented by type Cat
func (c Cat) Speak() string {
	return "Meow"
}

// Llama defines actions for dogs
type Llama struct {
}

// Speak method is from interface type animal and is implemented by type Llama
func (l Llama) Speak() string {
	return "?????"
}

// JavaProgrammer defines actions for JavaProgrammer
type JavaProgrammer struct {
}

// Speak method is from interface type animal and is implemented by type JavaProgrammer
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
