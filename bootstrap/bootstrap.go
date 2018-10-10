package bootstrap

import (
	"log"

	"github.com/juliotorresmoreno/parse-server/models"

	"github.com/juliotorresmoreno/parse-server/db"
)

func Inicialize() {
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Sync2(models.User{})
}
