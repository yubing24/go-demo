package main

import "fmt"

type Contact struct {
	Name string
	Age  int
}

// SelfIntroduction prints a self-intro message
func (contact Contact) SelfIntroduction() {
	fmt.Printf("Hi, I am %s and I am %d years old.\n", contact.Name, contact.Age)
}

func (contact *Contact) Aging() {
	contact.Age += 1
}

func main() {
	// Instantiate the Contact
	mycontact := Contact{
		Name: "Bob Smith",
		Age:  22,
	}
	mycontact.SelfIntroduction()
	mycontact.Aging() // this will increment Age by 1
	mycontact.SelfIntroduction()
}
