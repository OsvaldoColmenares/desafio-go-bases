package tickets

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct { //TODO: Usar esta struct para mapear los datos
	// ID      int
	// Nombre  string
	// Email   string
	// Destino string
	// Horario string
}

//var Tickets = make([]Ticket, 1000) //TODO: Definir un slice de tipo Ticket

func GetTotalTickets(destination string) (int, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0, errors.New("No se pudo abrir el archivo csv")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	records, _ := reader.ReadAll()

	cantidadViajes := 0
	//TODO: Mapear aqui los datos hacia Ticket
	for _, row := range records {
		if row[3] == destination {
			cantidadViajes++
		}
	}

	return cantidadViajes, nil
}

func GetCountByPeriod() (int, int, int, int, error) {
	// TODO: Utilizar el slice Ticket
	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0, 0, 0, 0, errors.New("No se pudo abrir el archivo csv")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	records, _ := reader.ReadAll()

	madrugada := 0
	maniana := 0
	tarde := 0
	noche := 0

	for _, row := range records {
		var hora = strings.Split(row[4], ":")
		var horario, err = strconv.Atoi(hora[0])
		if err != nil {
			return 0, 0, 0, 0, errors.New("No se pudo convertir la hora a int")
		}
		switch {
		case horario >= 0 && horario <= 6:
			madrugada++
		case horario >= 7 && horario <= 12:
			maniana++
		case horario >= 13 && horario <= 19:
			tarde++
		case horario >= 20 && horario <= 23:
			noche++
		default:
			return 0, 0, 0, 0, errors.New("Ese horario no pertenece a ningún turno.")
		}
	}

	return madrugada, maniana, tarde, noche, nil
}

func AverageDestination(destination string) (float64, error) {
	// TODO: Utilizar el slice Ticket
	file, err := os.Open("tickets.csv")
	if err != nil {
		return 0, errors.New("No se pudo abrir el archivo csv")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	records, _ := reader.ReadAll()

	c := 0
	for _, row := range records {
		if row[3] == destination {
			c++
		}
	}

	var cantidad float64 = float64(len(records))
	var promedio float64 = (float64(c) / cantidad) * 100

	return promedio, nil
}
