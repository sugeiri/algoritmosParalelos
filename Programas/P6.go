package main

//NUMEROS PRIMOS
import (
	"fmt"
)
var (
	i int
)
func main()	{
	res := true
	for ; res; {
		lectura()
		ValidaPrimo(i)
		res = finalizar()
	}
}
func lectura() int {
	fmt.Println("\nIngrese un Numero")
	fmt.Println("---------------------")
    fmt.Print("-> ")
	fmt.Scan(&i)
	return i
}

func ValidaPrimo(n int ) {
	sum := 0
	for a := 1; a <= n; a++ {
		if n%a==0 {
			sum++
		}
	}
	if sum>2{
		fmt.Printf("**El numero %d No Es primo**\n",n)
		return
	}
	fmt.Printf("**El numero %d Es primo**\n",n)
	return 
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
