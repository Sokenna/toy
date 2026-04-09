package greeting

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//return
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}
func randomFormat() string {
	forms := []string{
		"Hi,%v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return forms[rand.Intn(len(forms))]
}

func Hellos(names []string) (map[string]string, error) {
	greeting := map[string]string{}
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		greeting[name] = message
	}
	return greeting, nil
}
