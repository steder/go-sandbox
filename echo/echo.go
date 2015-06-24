/*
 Implementing echo command in go
*/

package main

import "fmt"
import "os"
import "strings"

func main() {
	message := strings.Join(os.Args[1:], " ")
	fmt.Println(message)
}
