package util

import "time"

/*
	layout : formato do pacote time:
		ano-mes-diaMarcadoHora:Minuto:Segundo

	Mais detalhes em: https://golang.org/pkg/time/
*/
const layout = "2006-01-02T15:04:05"

// StringToTime :
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, value)
	return convertedTime
}
