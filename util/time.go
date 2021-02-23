package util

import "time"

/*
	layout : formato do pacote time:
		ano-mes-diaMarcadoHora:Minuto:Segundo

	Mais detalhes em: https://golang.org/pkg/time/

	No caso criamos uma const chamada layout no lugar de uma var, pois o layout será usado e não alterado.
*/
const layout = "2006-01-02T15:04:05"

// StringToTime : Função usada para conversão de string para o Time.
func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, value)
	return convertedTime
}
