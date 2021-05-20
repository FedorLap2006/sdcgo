package main

import (
	"flag"
	"fmt"
	"github.com/FedorLap2006/sdcgo"
	dotenv "github.com/joho/godotenv"
	"os"
)

var (
	BotID = flag.String("bot", "", "ID of a bot for testing")
	Shards = flag.Int("shards", 1, "Number of shards")
	Guilds = flag.Int("guilds", 1, "Number of guilds")
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
	status, err := session.PostStats(*BotID, sdcgo.BotStats{
		Guilds: *Guilds,
		Shards: *Shards,
	})

	if err != nil {
		panic(err)
	} else if status != true {
		fmt.Println("Stats weren't posted")
	}
	fmt.Println("Stats successfully posted")
}
