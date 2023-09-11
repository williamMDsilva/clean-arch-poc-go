package main

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/williamMDsilva/clean-arch-poc-go/adapter/api/docs"
	"github.com/williamMDsilva/clean-arch-poc-go/adapter/postgres"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title Clean GO CLI
// @version 1.0.0
// @contact.name William Moreira da Silva
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	defer fmt.Println("Bye!")

	postgres.RunMigrations()

	log.Printf("Run migration %v", ctx)
}
