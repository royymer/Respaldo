package main

import "fmt"
import "math"

//asi se declara una interface
type calculosGeometricos interface{
  area() float64
  perimetro() float64
}

//a continuacion estructura del cuadro
type cuadrado struct{
  ancho, alto float64
}

//a continuacion estructura para cirulos
 type circulo struct{
   radio float64
 }

 //funcion para calcular el area del cuadrado
 //la funcion se le pasa un parametro "s" de tipo cuadro
 //y regresara el area de la interface que es de tipo float 
 func (s cuadrado) area() float64  {
   return s.ancho * s.alto
 }

 //funcion para calcular el perimetro del cuadrado

 func (s cuadrado) perimetro() float64  {
      return 4 * (s.ancho + s.alto)
 }

 //funcion para calcular el area de un circulo

 func (c circulo) area() float64  {
   return math.Pi * c.radio * c.radio
 }

 //funcion para calcular el perimetro delcirculo

 func (c circulo) perimetro() float64  {
   return 2 * math.Pi * c.radio
 }

 //utilizo la interface geometria para los calculosGeometricos
 func medida(g calculosGeometricos)  {
   fmt.Println(g)
   fmt.Println(g.area())
   fmt.Println(g.perimetro())
 }

 //funcion principal

 func main()  {
   s := cuadrado{ancho: 3 , alto: 4}
   c := circulo{radio : 5}

   medida(s)
   medida(c)
 }
