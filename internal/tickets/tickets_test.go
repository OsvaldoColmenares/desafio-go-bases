// Crear test unitarios para cada uno de los requerimientos anteriores, mínimo 2 casos por requerimiento.

package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTotalTickets_CountryDestination_ReturnTotalTicket(t *testing.T) {
	//Arrange
	const resultadoEsperado int = 2
	const destination string = "China"
	TicketsTests := []Ticket{
		{1, "Tait Mc Caughan", "tmc0@scribd.com", "Brazil", "17:11", 785},
		{2, "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537},
		{3, "Yalonda Jermyn", "yjermyn2@omniture.com", "Argentina", "18:11", 579},
		{4, "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
		{5, "Saree Nobes", "snobes4@google.com.au", "China", "0:31", 1398},
	}

	//Act
	result := GetTotalTickets(destination, TicketsTests)

	//Assert
	assert.Equal(t, resultadoEsperado, result)
}

func Test_GetTotalTickets_CountryDestination_ReturnZeroTickets(t *testing.T) {
	//Arrange
	const resultadoEsperado int = 0
	const destination string = "Francia"
	TicketsTests := []Ticket{
		{1, "Tait Mc Caughan", "tmc0@scribd.com", "Brazil", "17:11", 785},
		{2, "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537},
		{3, "Yalonda Jermyn", "yjermyn2@omniture.com", "Argentina", "18:11", 579},
		{4, "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
		{5, "Saree Nobes", "snobes4@google.com.au", "China", "0:31", 1398},
	}

	//Act
	result := GetTotalTickets(destination, TicketsTests)

	//Assert
	assert.Equal(t, resultadoEsperado, result)
}

func Test_GetCountByPeriod_AllPeriod_ReturnTotalCount(t *testing.T) {
	//Arrange
	const resultadoMadrugada int = 1
	const resultadoManiana int = 0
	const resultadoTarde int = 2
	const resultadoNoche int = 2

	TicketsTests := []Ticket{
		{1, "Tait Mc Caughan", "tmc0@scribd.com", "Brazil", "17:11", 785},
		{2, "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537},
		{3, "Yalonda Jermyn", "yjermyn2@omniture.com", "Argentina", "18:11", 579},
		{4, "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
		{5, "Saree Nobes", "snobes4@google.com.au", "China", "0:31", 1398},
	}
	tp := IntervaloHorario{}

	// Act
	result, _ := tp.GetCountByPeriod(TicketsTests)

	//Arrange
	assert.Equal(t, resultadoMadrugada, result.Madrugada)
	assert.Equal(t, resultadoManiana, result.Maniana)
	assert.Equal(t, resultadoTarde, result.Tarde)
	assert.Equal(t, resultadoNoche, result.Noche)
}

func Test_GetCountByPeriod_AllPeriod_ReturnError(t *testing.T) {
	//Arrange
	const errorEsperado string = "Horario invalido."

	TicketsTests := []Ticket{
		{1, "Tait Mc Caughan", "tmc0@scribd.com", "Brazil", "17:11", 785},
		{2, "Padget McKee", "pmckee1@hexun.com", "China", "34:19", 537},
		{3, "Yalonda Jermyn", "yjermyn2@omniture.com", "Argentina", "18:11", 579},
		{4, "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
		{5, "Saree Nobes", "snobes4@google.com.au", "China", "0:31", 1398},
	}
	tp := IntervaloHorario{}

	// Act
	_, err := tp.GetCountByPeriod(TicketsTests)

	//Arrange
	assert.Equal(t, errorEsperado, err.Error())
}

func Test_AverageDestination_CountryDestinationPercentage_ReturntCountryPercentage(t *testing.T) {
	//Arrange
	const resultadoEsperado float64 = 40
	const destination string = "China"
	TicketsTests := []Ticket{
		{1, "Tait Mc Caughan", "tmc0@scribd.com", "Brazil", "17:11", 785},
		{2, "Padget McKee", "pmckee1@hexun.com", "China", "20:19", 537},
		{3, "Yalonda Jermyn", "yjermyn2@omniture.com", "Argentina", "18:11", 579},
		{4, "Diannne Pharrow", "dpharrow3@icio.us", "Mongolia", "23:16", 1238},
		{5, "Saree Nobes", "snobes4@google.com.au", "China", "0:31", 1398},
	}

	//Act
	result := AverageDestination(destination, TicketsTests)

	//Assert
	assert.Equal(t, resultadoEsperado, result)
}
