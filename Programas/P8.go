package main

//SIGNO ZODIACAL
import (
	"fmt"
)
var (
	palabra string
)
func main()	{
	res := true
	for ; res; {
		Busca_Signo(lectura())
		res = finalizar()
	}
}
func lectura() [2]int {
	var nac [2]int
	fmt.Println("\nIngrese Dia")
	fmt.Println("--------------------------")
    fmt.Print("-> ")
	fmt.Scan(&nac[0])
	fmt.Println("\nIngrese Mes")
	fmt.Println("--------------------------")
    fmt.Print("-> ")
	fmt.Scan(&nac[1])
	return nac
}

func Busca_Signo(n [2] int ) {
	switch n[1]{
	case 1:
		if n[0]<=20{
			fmt.Printf("**Su Signo Zodiacal es Capricornio**\n")
			return 
		} 
		if n[0]>=21{
			fmt.Printf("**Su Signo Zodiacal es Acuario**\n")
			return
		} 
		
		break
	case 2:
		if n[0]<=19{
			fmt.Printf("**Su Signo Zodiacal es Acuario**\n")
			return 
		} 
		if n[0]>=20{
			fmt.Printf("**Su Signo Zodiacal es Piscis**\n")
			return
		} 
		break
	case 3:
		if n[0]<=20{
			fmt.Printf("**Su Signo Zodiacal es Piscis**\n")
			return 
		} 
		if n[0]>=21{
			fmt.Printf("**Su Signo Zodiacal es Aries**\n")
			return
		} 
		break
	case 4:
		if n[0]<=20{
			fmt.Printf("**Su Signo Zodiacal es Aries**\n")
			return 
		} 
		if n[0]>=21{
			fmt.Printf("**Su Signo Zodiacal es Tauro**\n")
			return
		} 
		break
	case 5:
		if n[0]<=21{
			fmt.Printf("**Su Signo Zodiacal es Tauro**\n")
			return 
		} 
		if n[0]>=22{
			fmt.Printf("**Su Signo Zodiacal es Geminis**\n")
			return
		} 
		break
	case 6:
		if n[0]<=21{
			fmt.Printf("**Su Signo Zodiacal es Geminis**\n")
			return 
		} 
		if n[0]>=22{
			fmt.Printf("**Su Signo Zodiacal es Cancer**\n")
			return
		} 
		break
	case 7:
		if n[0]<=22{
			fmt.Printf("**Su Signo Zodiacal es Cancer**\n")
			return 
		} 
		if n[0]>=23{
			fmt.Printf("**Su Signo Zodiacal es Leo**\n")
			return
		} 
		break
	case 8:
		if n[0]<=23{
			fmt.Printf("**Su Signo Zodiacal es Leo**\n")
			return 
		} 
		if n[0]>=24{
			fmt.Printf("**Su Signo Zodiacal es Virgo**\n")
			return
		} 
		break
	case 9:
		if n[0]<=23{
			fmt.Printf("**Su Signo Zodiacal es Virgo**\n")
			return 
		} 
		if n[0]>=24{
			fmt.Printf("**Su Signo Zodiacal es Libra**\n")
			return
		} 
		break
	case 10:
		if n[0]<=23{
			fmt.Printf("**Su Signo Zodiacal es Libra**\n")
			return 
		} 
		if n[0]>=24{
			fmt.Printf("**Su Signo Zodiacal es Escorpio**\n")
			return
		} 
		break
	case 11:
		if n[0]<=22{
			fmt.Printf("**Su Signo Zodiacal es Escorpio**\n")
			return 
		} 
		if n[0]>=23{
			fmt.Printf("**Su Signo Zodiacal es Sagitario**\n")
			return
		} 
		break
	case 12:
		if n[0]<=21{
			fmt.Printf("**Su Signo Zodiacal es Sagitario**\n")
			return 
		} 
		if n[0]>=22{
			fmt.Printf("**Su Signo Zodiacal es Capricornio**\n")
			return
		} 
		break	
	}
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
