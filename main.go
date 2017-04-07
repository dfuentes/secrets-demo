package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Clever/stealth/store"
)

var Environment string
var Service string

func init() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		log.Fatal("must set ENVIRONMENT")
	}
	Environment = env

	service := os.Getenv("SERVICE")
	if service == "" {
		log.Fatal("must set SERVICE")
	}
	Service = service
}

func main() {
	s := store.NewUnicredsStore()

	ids, err := s.List(getEnvironment(Environment), Service)
	if err != nil {
		log.Fatal(err)
	}

	script := generateScript(s, ids)
	fmt.Fprintf(os.Stdout, script)
}

func generateScript(s store.SecretStore, ids []store.SecretIdentifier) string {
	script := ""
	for _, id := range ids {
		secret, err := s.Read(id)
		if err != nil {
			log.Fatal(err)
		}
		script += fmt.Sprintf("export _SECRET_%s=%s", strings.ToUpper(id.Key), secret.Data)
	}
	return script
}

func getEnvironment(env string) store.Environment {
	if strings.ToLower(env) == "production" {
		return store.ProductionEnvironment
	}
	return store.DevelopmentEnvironment
}
