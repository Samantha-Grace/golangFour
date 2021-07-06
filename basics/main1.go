// Go Bootcamp: Master Golang with 1000+ Exercises and Projects

//type conversion
package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	//speed = speed * intforce
	//speed = speed * int(force)

	speed = int(float64(speed) * force)

	fmt.Println(speed)
}
