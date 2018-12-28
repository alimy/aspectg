package main

import (
	"github.com/alimy/aspectg/cmd"
	_ "github.com/alimy/aspectg/version"
)

func main() {
	// Setup root cli command of application
	cmd.Setup(
		"aspectg",                                // command name
		"AspectG auto build AOP code for golang", // command short describe
		"AspectG auto build AOP code for golang.(Inspiration from wire and AspectJ)", // command long describe
	)

	// Execute start application
	cmd.Execute()
}
