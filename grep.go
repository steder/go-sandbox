/*
 golang version of grep

 Usage:

 For now we just support the following usage:

   grep <pattern> <path>

 */

package main

import "bufio"
import "fmt"
import "io"
import "os"
import "regexp"
import "strings"

func main() {
	usage := "Usage:" + strings.Join([]string{"\n",
		"grep <pattern> <file>",
		"\n"}, "\n\t")
	if len(os.Args) >= 3 {
		pattern := os.Args[1]
		filename := os.Args[2]
		re := regexp.MustCompile(`(?i)`+pattern)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Unable to open file: " + filename)
			return
		}
		defer file.Close()
		r := bufio.NewReader(file)
		for {
			line, err := r.ReadString('\n')
			if err != io.EOF {
				if re.FindString(line) != "" {
					fmt.Print(filename + ": " + line)
				}
			} else {
				break
			}
		}
	} else {
		fmt.Print(usage)
	}
}