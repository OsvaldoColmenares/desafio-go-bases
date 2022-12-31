package tickets

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID      int
	Nombre  string
	Email   string
	Destino string
	Horario string
	precio  float64
}

type IntervaloHorario struct {
	Madrugada int
	Maniana   int
	Tarde     int
	Noche     int
}

// type ITicket interface {
// 	GetTotalTickets()
// }

func (t Ticket) RegisterAll(rutaFile string) ([]Ticket, error) {
	file, err := os.Open(rutaFile)
	if err != nil {
		return nil, errors.New("No se pudo abrir el archivo csv")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("No se leer el archivo csv")
	}
	var Tickets = make([]Ticket, 0, 1000)
	for _, row := range records {
		if row[0] != "" && row[1] != "" && row[2] != "" && row[3] != "" && row[4] != "" {
			t.ID, err = strconv.Atoi(row[0])
			t.Nombre = row[1]
			t.Email = row[2]
			t.Destino = row[3]
			t.Horario = row[4]
			t.precio, err = strconv.ParseFloat(row[5], 64)
			Tickets = append(Tickets, t)
		}
	}
	return Tickets, nil
}

func GetTotalTickets(destination string, tickets []Ticket) int {
	cantidadViajes := 0
	for _, row := range tickets {
		if row.Destino == destination {
			cantidadViajes++
		}
	}
	return cantidadViajes
}

func (ih IntervaloHorario) GetCountByPeriod(tickets []Ticket) (*IntervaloHorario, error) {
	for _, row := range tickets {
		before, _, _ := strings.Cut(row.Horario, ":")
		var horario, _ = strconv.Atoi(before)
		switch {
		case horario >= 0 && horario <= 6:
			ih.Madrugada++
		case horario >= 7 && horario <= 12:
			ih.Maniana++
		case horario >= 13 && horario <= 19:
			ih.Tarde++
		case horario >= 20 && horario <= 23:
			ih.Noche++
		default:
			return &IntervaloHorario{}, errors.New("Horario invalido.")
		}
	}
	return &ih, nil
}

func AverageDestination(destination string, tickets []Ticket) float64 {
	var cantidad float64 = float64(len(tickets))
	var promedio float64 = (float64(GetTotalTickets(destination, tickets)) / cantidad) * 100
	return promedio
}
