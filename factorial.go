package main

import "fmt"

func factorial(n int) int  {
  if n == 0 {
    return 1 //cumpliendo la ley de que el factorial de 0 por definicion es 1
  }
  return n * factorial(n - 1) //si es distinto de 0 ejecuta la funcion de factorial
}

func main()  {
  fmt.Println(factorial(5)) //imprime el factorial de 5
}
