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

	tick := tickets.Ticket{}
	tick.RegisterAll("tickets.csv")

	destino := "Brazil"
	total := tickets.GetTotalTickets(destino)
	fmt.Println("Requerimiento 1")
	fmt.Printf("%d personas viajaron a %s\n", total, destino)
	fmt.Println("-------------")

	intervaloHorario := tickets.IntervaloHorario{}
	ih, err := intervaloHorario.GetCountByPeriod()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Requerimiento 2")
	fmt.Println("Madrugada:", ih.Madrugada)
	fmt.Println("Mañana:", ih.Maniana)
	fmt.Println("Tarde:", ih.Tarde)
	fmt.Println("Noche:", ih.Noche)
	fmt.Println("-------------")
	total2 := tickets.AverageDestination(destino)
	fmt.Println("Requerimiento 3")
	fmt.Println("Promedio:", total2, "%")
	fmt.Println("-------------")
}
