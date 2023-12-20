package main

import (
	"fmt"
	"example/dictionnaire/dictionary" 
)

func main() {
	dictionary := make(dictionary.Dictionary)

	dictionary.Add("estiam", "école.")
	dictionary.Add("Chaimae", "Prénom.")
	dictionary.Add("shaddow", "chat.")

	wordToGet := "go"
	fmt.Printf("Definition of %s: %s\n", wordToGet, dictionary.Get(wordToGet))

	wordToRemove := "Estiam" 
	fmt.Printf("Removing %s from the dictionary.\n", wordToRemove)
	dictionary.Remove(wordToRemove)

	fmt.Println("\nDictionary entries:")
	entries := dictionary.List()
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
