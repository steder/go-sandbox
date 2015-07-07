package main

/* This is a test! */

import "fmt"
import "io/ioutil"
import "os"
import "sort"
import "strings"

func crappy_character_sort(s string) string {
	letters := strings.Split(strings.ToLower(s), "")
	sort.Strings(letters)
	sorted := strings.Join(letters, "")
	return sorted
}

func load_dictionary() map[string][]string {
	content, err := ioutil.ReadFile("/usr/share/dict/words")
	var lines []string
	dictionary := map[string][]string{}

	if err == nil {
		lines = strings.Split(string(content), "\n")
	} else {
		fmt.Print("Error reading words file!")
	}

	fmt.Println("Loading up words list...")
	for _, line := range lines {
		sorted := crappy_character_sort(line)
		// fmt.Println(line + " " + sorted)
		dictionary[sorted] = append(dictionary[sorted], line)
	}
	return dictionary
}

func main() {
	if len(os.Args) > 1 {
		sorted := crappy_character_sort(os.Args[1])
		dictionary := load_dictionary()
		fmt.Println("Anagrams for " + os.Args[1])
		//fmt.Println(dictionary[sorted])
		for _, word := range dictionary[sorted] {
			fmt.Println("\t" + word)
		}
	} else {
		fmt.Println(`Usage:

 $ ./scrabble <word>
 $ ./scrabble <letters>

`)
	}
}
