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
	listTickets, err := tick.RegisterAll("tickets.csv")
	if err != nil {
		panic(err.Error())
	}

	destino := "Brazil"
	total := tickets.GetTotalTickets(destino, listTickets)
	fmt.Println("Requerimiento 1")
	fmt.Printf("%d personas viajaron a %s\n", total, destino)
	fmt.Println("-------------")

	intervaloHorario := tickets.IntervaloHorario{}
	ih, err := intervaloHorario.GetCountByPeriod(listTickets)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Requerimiento 2")
	fmt.Println("Madrugada:", ih.Madrugada)
	fmt.Println("Mañana:", ih.Maniana)
	fmt.Println("Tarde:", ih.Tarde)
	fmt.Println("Noche:", ih.Noche)
	fmt.Println("-------------")
	total2 := tickets.AverageDestination(destino, listTickets)
	fmt.Println("Requerimiento 3")
	fmt.Println("Promedio:", total2, "%")
	fmt.Println("-------------")
}
