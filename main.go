package main

import (
	"flag"
	"log"
	"myproject/cmd"
	"myproject/config"
)

func main() {
	configFile := flag.String("config", "config/config.yaml", "Path to the yaml configuration file")
	flag.Parse()
	if *configFile != "" {
		c, err := config.LoadFromFile(*configFile)
		if err != nil {
			log.Fatal(err)
		}
		config.Set(c)
	} else {
		log.Println("No configuration file")
		config.Set(config.NewConfig())
	}
	cmd.Execute()
}
