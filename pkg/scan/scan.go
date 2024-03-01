package scan

import (
	"fmt"
	"sincro/pkg/config"
)

func Start() {
	config, err := config.Read()
	if err != nil {
		fmt.Println("Config file 'sincro.json' not found. Please run 'sincro init'")
		return
	}

	fmt.Println(config)
}
