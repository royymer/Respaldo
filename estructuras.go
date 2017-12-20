package main

import "fmt"

type persona struct{
  nombre string
  cedula string
  direccion string
}

func main()  {
    persona1 := persona{"Roymer Centeno", "V-20073158", "Valencia Edo Carabobo"}
    fmt.Println(persona1)
}
