package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"flag"

	"github.com/gin-gonic/gin"
	log "github.com/mgutz/logxi/v1"
)

func main() {
	var data []byte

	cmdPokemon := flag.String("pokemon", os.Getenv("POKEMOM"), "Pokemon name to use")
	flag.Parse()

	if len(*cmdPokemon) > 0 {
		log.Warn("Preselected Pokemon", *cmdPokemon)
		data, _ = pokedata(fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s/", *cmdPokemon))
	}
	if data == nil {
		host, _ := os.Hostname()
		data, _ = pokedata(pokenum(host))
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/resources", "./resources")
	router.Static("/css", "./resources/css")
	router.Static("/images", "./resources/images")
	router.Static("/favicon.ico", "./resources/images/favicon.ico")
	router.StaticFile("/", "./resources/index.html")
	router.GET("/whoami", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, string(data))
	})
	router.Run(":8080")
}

func pokedata(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, err
}

func pokenum(name string) string {
	return fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%d/", strtoint(name)%811)
}

func strtoint(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		num += int(s[i])
	}
	return num
}
