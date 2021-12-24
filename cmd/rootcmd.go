package cmd

import (
	"fmt"
	"myproject/config"
	"myproject/pet"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "{name} {behavior} [--config]",
	Short:   "return result based on behavior",
	Long:    `return result based on name and behavior`,
	Example: "c1 eat",
	Args:    cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := run(args)
		if err != nil {
			fmt.Println("command receive error:", err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run(args []string) error {
	name := args[0]
	form := ""
	for _, v := range config.Get().Pets {
		if name == v.Name {
			form = v.Form
		}
	}
	behavior := args[1]
	if form == "" {
		return fmt.Errorf("this name is not in the pets")
	}
	i := pet.GetInterface(form)
	flag := os.Getenv("SHOUT")
	switch behavior {
	case "eat":
		fmt.Println(i.Eat(flag))
	case "speak":
		fmt.Println(i.Speak(flag))
	case "move":
		fmt.Println(i.Move(flag))
	}
	return nil
}
