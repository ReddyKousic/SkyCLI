package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ReddyKousic/skycli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: skycli <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  setapi   Set the API key")
		fmt.Println("  setloc   Set the locality")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "setapi":
		setAPICmd := flag.NewFlagSet("setapi", flag.ExitOnError)
		apiKey := setAPICmd.String("key", "", "WeatherAPI API key")
		setAPICmd.Parse(os.Args[2:])

		if *apiKey == "" {
			fmt.Println("Please provide an API key using the -key flag")
			setAPICmd.PrintDefaults()
			os.Exit(1)
		}

		err := cmd.SaveAPIKey(*apiKey)
		if err != nil {
			fmt.Printf("Error saving API key: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("API key saved successfully")

	case "setloc":
		setLocCmd := flag.NewFlagSet("setloc", flag.ExitOnError)
		locality := setLocCmd.String("loc", "", "Locality")
		setLocCmd.Parse(os.Args[2:])

		if *locality == "" {
			fmt.Println("Please provide a locality using the -loc flag")
			setLocCmd.PrintDefaults()
			os.Exit(1)
		}

		err := cmd.SaveLocality(*locality)
		if err != nil {
			fmt.Printf("Error saving locality: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Locality saved successfully")

	default:
		fmt.Printf("%q is not a valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
