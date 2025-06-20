package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, implementer Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", implementer.LanguageName(), implementer.Greet(name))
}
