package main;

import "fmt"

func greeting(name string) {
  fmt.Printf("Hello, %s!", name) 
}

func main() {
  // explicit
  var name string = "John"

  // type inferred
  var surname = "Johnson"

  // shorthand
  human := true

  const birthDate = 1991
}
