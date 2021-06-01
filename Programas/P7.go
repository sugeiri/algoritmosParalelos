package main

//PALABRAS PALINDROMAS
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
		if ValidaPalindroma(palabra) {
			fmt.Printf("**la Palabra %s es Palindroma**\n",palabra)
		}else{
			fmt.Printf("**la Palabra %s No es Palindroma**\n",palabra)
		}
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

func ValidaPalindroma(n string ) bool {
	var s string
	for a := len(n)-1; a>=0; a-- {
		s+=string(n[a])
	}
	if n == s{
		return true
	}
	return  false
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
