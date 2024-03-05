package ui

import "github.com/charmbracelet/huh"

func PromptText(
	title string,
	prompt string,
	validationFnc func(string) error,
) string {
	var response string

	input := huh.NewInput().
		Title(title).
		Prompt(prompt).
		Validate(validationFnc).
		Value(&response)

	input.Run()

	return response
}
