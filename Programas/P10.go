package main

//ESTRUCTURAS Y LISTAS
import (
	"fmt"
)


type Persona struct {
	id int
	nombre string
	edad int
}
var(
	L_Persona [] Persona
)
func Menu() int {
	var x int
	fmt.Println("\nQue Desea Hacer?")
	fmt.Println("--------------------------------------------")
	fmt.Print("1->Agregar Persona\n2->Ver Todos los Datos \n3->Eliminar Persona\n4->Limpiar Todos los datos\n5->Salir\n")
	fmt.Scan(&x)
	return x
}
func main() {
	res := 0
	res= Menu()
	for ; res!=5; {
		switch res {
			case 1:		
				//AGREGAR PERSONA
				AgregaPersona()
				break
			case 2:
				//VER TODAS LAS PERSONAS
				ImprimePersonas()
				break
			case 3:
				//ELIMINAR PERSONA
				EliminarPersona()
				break
			case 4:
				L_Persona= nil
				fmt.Printf("\n**DATOS BORRADOS**\n")
				//LIMPIAR TODO
				break
			}
			res= Menu()
		}
}

func AgregaPersona() {
	var id int
	var nombre string
	var edad int
	fmt.Println("\nIngrese ID:")
	fmt.Println("-----------------------")
	fmt.Print("-> ")
	fmt.Scan(&id)
	fmt.Println("\nIngrese Nombre:")
	fmt.Println("-----------------------")
	fmt.Print("-> ")
	fmt.Scan(&nombre)
	fmt.Println("\nIngrese Edad:")
	fmt.Println("-----------------------")
	fmt.Print("-> ")
	fmt.Scan(&edad)
	ii_persona:=Persona{id,nombre,edad}
	L_Persona = append(L_Persona,ii_persona)
	fmt.Printf("\n**AGREGADA**\n")
	
}

func ImprimePersonas() {
	fmt.Printf("\n*******************************************\n")
	fmt.Printf("\nID\tNombre\t\tEdad\t\n")
	for _, ii_persona := range L_Persona {
		fmt.Printf("\n%-3d\t%-10s\t%-2d\n",ii_persona.id,ii_persona.nombre,ii_persona.edad)
	}
	fmt.Printf("\n*******************************************\n")
	return 
}
func EliminarPersona() {
	ImprimePersonas()
	var id int
	fmt.Printf("\nIngrese ID A Borrar:")
	fmt.Println("-----------------------")
	fmt.Print("-> ")
	fmt.Scan(&id)
	for i, ii_persona := range L_Persona {
		if ii_persona.id==id{
			L_Persona = append(L_Persona[:i], L_Persona[i+1:]...)
		}
	}
	fmt.Printf("\n**Borrado**\n")
	ImprimePersonas()
	return 
}