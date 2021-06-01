package main

//IMPRIMIR STRING AL DERECHO Y AL REVES
import (
	"fmt"
)
var (
	palabra string
)
func main()	{
	res := true
	for ; res; {
		lectura()
		Imprimir(palabra)
		res = finalizar()
	}
}
func lectura() string {
	fmt.Println("\nIngrese Palabra")
	fmt.Println("---------------------")
    fmt.Print("-> ")
	fmt.Scan(&palabra)
	return palabra
}

func Imprimir(n string ) {
	var s string
	fmt.Printf("\nPalabra: %s ",n)
	for a := len(n)-1; a>=0; a-- {
		s+=string(n[a])
	}
	fmt.Printf("\nAl reves: %s",s)
	
}
func finalizar() bool {
	var x int
	fmt.Println("\nPulse 0 Para seguir y 1 para finalizar")
	fmt.Println("--------------------------------------------")
    fmt.Print("-> ")
    fmt.Scan(&x)
	if x==0{
		return true
	}
	return false
	
}
