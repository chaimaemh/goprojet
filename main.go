package main

import (
	"fmt"
	"goprojet/dictionnaire/dictionary"
)

func main() {
	dict := dictionary.NewDictionary("dictionary.txt")

	err := dict.Add("chaimae", "Prénom")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	err = dict.Add("chat", "animal")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	err = dict.Add("estiam", "école")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
	}

	
	wordToGet := "estiam"
	definition, err := dict.Get(wordToGet)
	if err != nil {
		fmt.Println("Erreur lors de la récupération :", err)
	} else {
		fmt.Printf("Definition of %s: %s\n", wordToGet, definition)
	}

	
	wordToRemove := "chaimae" 
	err = dict.Remove(wordToRemove)
	if err != nil {
		fmt.Println("Erreur lors de la suppression :", err)
	} else {
		fmt.Printf("Mot %s supprimé du dictionnaire.\n", wordToRemove)
	}

	fmt.Println("\nDictionary entries:")
	entries, err := dict.List()
	if err != nil {
		fmt.Println("Erreur lors de la récupération de la liste :", err)
	} else {
		for _, entry := range entries {
			fmt.Println(entry)
		}
	}
}
