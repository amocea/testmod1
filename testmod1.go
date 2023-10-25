package testmod1

import (
	"errors"
	"fmt"
)

func Hi(str string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi %s!", str), nil
	case "pt":
		return fmt.Sprintf("Oi %s!", str), nil
	case "es":
		return fmt.Sprintf("Hola %s!", str), nil
	case "fr":
		return fmt.Sprintf("Bonjour %s!", str), nil
	default:
		return "", errors.New("unknown language")
	}
}
