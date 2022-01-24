package main

import "fmt"

func main() {

	staircase(6)
}


func staircase(jml int) {

	totalSpace:=jml-1
	for i:=0; i<=jml; i++ {
		for j:=0; j<=jml; j++ {
			if j> totalSpace {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
		totalSpace--
	}

}