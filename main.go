package main

import (
	"fmt"
	"goprojet/dictionnaire/dictionary"
	"sync"
)

func main() {
	dict := dictionary.NewDictionary("dictionary.txt")
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		dict.Add("estiam", "école.")
		dict.Add("chat", "An animal.")
		dict.Add("Air france", "Entreprise.")
	}()

	go func() {
		defer wg.Done()
		dict.Remove("chaimae")
		dict.Remove("go")
	}()

	wg.Wait()

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
