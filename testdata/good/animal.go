package testenums

type Animal int

// There are no other animals.
const (
	Dog Animal = iota + 1
	Cat
	Bird
	Giraffe
)
