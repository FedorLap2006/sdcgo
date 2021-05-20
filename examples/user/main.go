
package main

import (
	"flag"
	"fmt"
	"github.com/FedorLap2006/sdcgo"
	dotenv "github.com/joho/godotenv"
	"os"
)

var (
	UserID = flag.String("user", "", "ID of an user for testing")
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
	rates, err := session.UserRates(*UserID)

	if err != nil {
		panic(err)
	}
	for guild, rate := range rates {
		fmt.Printf("%v on %s\n", rate, guild)
	}
}

