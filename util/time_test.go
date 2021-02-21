package util

import "testing"

// TestStringToTime
func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2021-02-21T20:19:05")

	// convertedTime.Year : ano
	if convertedTime.Year() != 2021 {
		t.Error("Espera que o ano seja 2021")
	}

	// convertedTime.Month : mes
	if convertedTime.Month() != 02 {
		t.Error("Espera que o mes seja 02")
	}

	// convertedTime.Hour : hora
	if convertedTime.Hour() != 10 {
		t.Error("Espera que a hora seja 10")
	}
}
