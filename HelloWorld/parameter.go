
package main
import (
    "fmt"
    "os"
)
func main () {
	command := os.Args
	args := os.Args[1:]

	fmt.Println(command)
	fmt.Println(args)
}