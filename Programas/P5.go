package main

//ARREGLO ENTERO ORDENADO
import (
	"fmt"
	"sort"
)
var (
	arreglo []int
)
func main()	{
	lectura()
	Imprimir()
	  
}
func lectura() {
	var x int
	x=1
	fmt.Println("\nIngrese los Valores (Ingrese 0 para terminar)")
	fmt.Println("-------------------------------------------")
	for ;x!=0;
	{
		fmt.Print("-> ")
		fmt.Scan(&x)
		if x!=0{
			arreglo=append(arreglo, x)
		}
	}
}
func Imprimir() {
	fmt.Printf("\n ARREGLO")
	fmt.Println(arreglo)
	fmt.Printf("\n ORDENADO")
	sort.Ints(arreglo)
	fmt.Println(arreglo)
	
}


