package main

import (
	_ "../sample/matchers"
	"../sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("Coronavirus")

}
