package ExercismGo

import "fmt"

// Shape struct is the equivalent of a Class from OOP
// Once declared, it can then be instantiated and used
// by creating and instantiating a variable with it.
type Shape struct {
	name string
	size int
}

func main() {
	s := Shape{
		name: "Square",
		size: 12,
	}
	s.name = "Circle"
	s.size = 34

	fmt.Println(s)

	NewShape("Triangle", 10)

	// It is recommended to use this only as a last resort.
	// Try using struct literals or New functions instead.
	quickInitialisedShape := new(Shape)
	fmt.Println("name: %s, size: %d", quickInitialisedShape.name, quickInitialisedShape.size)
}

// NewShape In Go, we create constructors by creating a function with the same name as the Struct
// but starting with "New" in the name, ex: "NewBook", "NewStruct2"
func NewShape(name string, size int) Shape {
	return Shape{
		name: name,
		size: size,
	}
}
