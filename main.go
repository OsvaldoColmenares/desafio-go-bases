package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Se recuperó")
		}
		fmt.Println("ejecución finalizada")
	}()

	destino := "Brazil"
	total, err := tickets.GetTotalTickets(destino)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Requerimiento 1")
	fmt.Printf("%d personas viajaron a %s\n", total, destino)
	fmt.Println("-------------")
	madrugada, maniana, tarde, noche, err := tickets.GetCountByPeriod()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Requerimiento 2")
	fmt.Println("Madrugada:", madrugada)
	fmt.Println("Mañana:", maniana)
	fmt.Println("Tarde:", tarde)
	fmt.Println("Noche:", noche)
	fmt.Println("-------------")
	total2, err := tickets.AverageDestination(destino)
	fmt.Println("Requerimiento 3")
	fmt.Println("Promedio:", total2, "%")
	fmt.Println("-------------")
}
