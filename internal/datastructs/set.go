package datastructs

import "fmt"

type Set struct {
	fruitsSet fruitsSet
}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Run() {
	fmt.Println("Filling set WAY 1")

	// TWO WAYS
	// WAY 1

	// set of strings
	fruits := map[string]struct{}{}

	// We can add members to the set
	// by setting the value of each key to an
	// empty struct
	fruits["banana"] = struct{}{}
	fruits["apple"] = struct{}{}
	fruits["grape"] = struct{}{}

	fmt.Println(fruits)

	// Adding a new member to the set
	fruits["orange"] = struct{}{}

	fmt.Println(fruits)

	// Adding an existing to the set
	fruits["banana"] = struct{}{}

	fmt.Println(fruits)

	// Removing a member from the set
	delete(fruits, "apple")
	fmt.Println(fruits)

	for fruit := range fruits {
		fmt.Println(fruit)
	}

	fmt.Println("Filling set WAY 2")
	// WAY 2
	// We can add members to the set
	// by setting the value of each key to an
	// empty struct
	s.fruitsSet = make(map[string]struct{}, 0)
	s.fruitsSet.add("banana")
	s.fruitsSet.add("apple")
	s.fruitsSet.add("grape")
	fmt.Println(s.fruitsSet)

	// Adding a new member to the set
	s.fruitsSet.add("orange")
	fmt.Println(s.fruitsSet)

	// Adding an existing to the set
	s.fruitsSet.add("banana")
	fmt.Println(s.fruitsSet)

	// Removing a member from the set
	if s.fruitsSet.has("apple") {
		s.fruitsSet.remove("apple")
	}
	fmt.Println(s.fruitsSet)
	if s.fruitsSet.has("coconut") {
		s.fruitsSet.remove("coconut")
	}
	fmt.Println(s.fruitsSet)
}

type fruitsSet map[string]struct{}

// Adds an animal to the set
func (s fruitsSet) add(fruit string) {
	s[fruit] = struct{}{}
}

// Removes an animal from the set
func (s fruitsSet) remove(fruit string) {
	delete(s, fruit)
}

// Returns a boolean value describing if the animal exists in the set
func (s fruitsSet) has(fruit string) bool {
	_, ok := s[fruit]
	return ok
}
