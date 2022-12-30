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

var Tickets = make([]Ticket, 0, 1000)

func (t Ticket) RegisterAll(rutaFile string) error {
	file, err := os.Open(rutaFile)
	if err != nil {
		return errors.New("No se pudo abrir el archivo csv")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ','
	records, _ := reader.ReadAll()

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
	return nil
}

func GetTotalTickets(destination string) int {
	cantidadViajes := 0
	for _, row := range Tickets {
		if row.Destino == destination {
			cantidadViajes++
		}
	}
	return cantidadViajes
}

func (ih IntervaloHorario) GetCountByPeriod() (*IntervaloHorario, error) {
	for _, row := range Tickets {
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

func AverageDestination(destination string) float64 {
	var cantidad float64 = float64(len(Tickets))
	var promedio float64 = (float64(GetTotalTickets(destination)) / cantidad) * 100
	return promedio
}
