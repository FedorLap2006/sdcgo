package main

import (
	"flag"
	"fmt"
	"github.com/FedorLap2006/sdcgo"
	dotenv "github.com/joho/godotenv"
	"os"
)

var (
	GuildID = flag.String("guild", "", "ID of a guild for testing")
)

func init() {
	err := dotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	session := sdcgo.New(os.Getenv("SDC_TOKEN"))
	g, err := session.Guild(*GuildID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s\n", g.Name)
	fmt.Printf("Desc: %s\n", g.Desc)
	fmt.Printf("Status: %d\n", g.Status)
	place, err := session.GuildPlace(*GuildID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Place: %d\n", place)

	rates, err := session.GuildRates(*GuildID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Rates:")
	for u, v := range rates {
		fmt.Printf("\t%s left: %v", u, v)
	}
}