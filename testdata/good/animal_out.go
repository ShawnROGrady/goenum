// Code generated by goenum. DO NOT EDIT.

package testenums

import "fmt"

var animalName = map[Animal]string{
	Dog:     "Dog",
	Cat:     "Cat",
	Bird:    "Bird",
	Giraffe: "Giraffe",
}

var animalValue = map[string]Animal{
	"Dog":     Dog,
	"Cat":     Cat,
	"Bird":    Bird,
	"Giraffe": Giraffe,
}

func (e Animal) String() string {
	return animalName[e]
}

type InvalidAnimalNameError struct {
	Name string
}

func (e *InvalidAnimalNameError) Error() string {
	return fmt.Sprintf("invalid Animal %q", e.Name)
}

func FromString(name string) (Animal, error) {
	val, ok := animalValue[name]
	if !ok {
		return Animal(0), &InvalidAnimalNameError{
			Name: name,
		}
	}

	return val, nil
}

func (e Animal) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

func (e *Animal) UnmarshalText(text []byte) error {
	name := string(text)
	val, ok := animalValue[name]
	if !ok {
		return &InvalidAnimalNameError{
			Name: name,
		}
	}

	*e = val
	return nil
}
