package main

import (
	"fmt"
	"time"
)

func main() {
	amarillo := time.Tick(50 * time.Millisecond)
	rojo := time.After(100 * time.Millisecond)
	for {
		select {
		case <-amarillo:
			fmt.Println("Semaforo en Amarillo")
		case <-rojo:
			fmt.Println("Semaforo en Rojo")
			return
		default:
			fmt.Println("Semaforo en Verde")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
