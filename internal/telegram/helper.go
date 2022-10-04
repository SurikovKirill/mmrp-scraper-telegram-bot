package telegram

import "fmt"

// Text is struct for message
type Text struct {
	date string
	data map[string]string
}

// ToString is method for making string from struct
func (t *Text) ToString() string {
	result := fmt.Sprintf("%s \n", t.date)
	for key, value := range t.data {
		result += fmt.Sprintf("%s: %s \n", key, value)
	}

	return result
}
