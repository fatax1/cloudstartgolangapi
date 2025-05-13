package main

import (
	"math/rand"
	"net/http"
	"time"

	"cloudgolangapi/data"

	"github.com/gin-gonic/gin"
)

var config Config
var theRandom *rand.Rand

// start är en enkel route som returnerar text.
func start(c *gin.Context) {
	c.Data(http.StatusOK, "application/text", []byte("Tjena"))
}

// enableCors hanterar CORS (Cross-Origin Resource Sharing).
func enableCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

// apiStats returnerar statistik om antalet spel och antal vinster.
func apiStats(c *gin.Context) {
	enableCors(c)
	totalGames, wins := data.Stats()
	c.JSON(http.StatusOK, gin.H{
		"totalGames": totalGames,
		"wins":       wins,
		"coolest":    "Abdil",
		"test":       "ny version av keel automatiskt",
	})
}

// apiPlay hanterar logiken för att spela spelet "Sten, Sax, Påse".
func apiPlay(c *gin.Context) {
	enableCors(c)
	yourSelection := c.Query("yourSelection")
	mySelection := randomizeSelection()
	winner := "Tie"

	// Logik för att bestämma vinnaren
	if yourSelection == "STONE" && mySelection == "SCISSOR" {
		winner = "You"
	}
	if yourSelection == "SCISSOR" && mySelection == "BAG" {
		winner = "You"
	}
	if yourSelection == "BAG" && mySelection == "STONE" {
		winner = "You"
	}
	if mySelection == "STONE" && yourSelection == "SCISSOR" {
		winner = "Computer"
	}
	if mySelection == "SCISSOR" && yourSelection == "BAG" {
		winner = "Computer"
	}
	if mySelection == "BAG" && yourSelection == "STONE" {
		winner = "Computer"
	}

	// Spara spelet i databasen
	err := data.SaveGame(yourSelection, mySelection, winner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the game"})
		return
	}

	// Svara med resultatet
	c.JSON(http.StatusOK, gin.H{
		"winner":            winner,
		"yourSelection":     yourSelection,
		"computerSelection": mySelection,
	})
}

// randomizeSelection genererar ett slumpmässigt val för datorn.
func randomizeSelection() string {
	val := theRandom.Intn(3) + 1
	switch val {
	case 1:
		return "STONE"
	case 2:
		return "SCISSOR"
	case 3:
		return "BAG"
	default:
		return "ERROR" // Förhindra ogiltiga resultat
	}
}

// main-funktionen startar applikationen och sätter upp alla rutter och databasanrop.
func main() {
	// Skapa en ny random generator baserat på systemets tid
	theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Läs konfigurationen från config (se till att readConfig är definierad korrekt i din kod)
	readConfig(&config)

	// Initiera databasen baserat på konfigurationen
	data.InitDatabase(
		config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port,
	)

	// Skapa en Gin-router
	router := gin.Default()
	router.GET("/", start)
	router.GET("/api/play", apiPlay)
	router.GET("/api/stats", apiStats)

	// Starta servern på port 8080
	router.Run(":8080")
}
