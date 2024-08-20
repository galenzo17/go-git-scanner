package main

import (
	"fmt"
	"log"
	"time"

	"go-git-scanner/src/utils"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

// Constantes para intervalos comunes
const (
	Every1Minute  = 1 * time.Minute
	Every5Minutes  = 5 * time.Minute
	Every30Minutes = 30 * time.Minute
	EveryHour      = time.Hour
	Every12Hours   = 12 * time.Hour
	EveryDay       = 24 * time.Hour
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error cargando archivo .env: %v", err)
    }
	// Ejemplo de uso: escaneo cada 5 minutos
	RunScannerWithInterval(Every1Minute)
}

func RunScannerWithInterval(interval time.Duration) {
	c := cron.New()
	
	// Convertir la duración a una expresión cron
	cronExpr := durationToCronExpr(interval)
	
	_, err := c.AddFunc(cronExpr, runScan)
	if err != nil {
		log.Fatal("Error al programar el cron job:", err)
	}

	log.Printf("Scanner programado para ejecutarse cada %v\n", interval)
	c.Start()
	select {} // Mantener el programa en ejecución
}

func durationToCronExpr(d time.Duration) string {
	switch {
	case d < time.Hour:
		minutes := int(d.Minutes())
		return fmt.Sprintf("*/%d * * * *", minutes)
	case d < 24*time.Hour:
		hours := int(d.Hours())
		return fmt.Sprintf("0 */%d * * *", hours)
	default:
		days := int(d.Hours() / 24)
		return fmt.Sprintf("0 0 */%d * *", days)
	}
}

func runScan() {
	log.Printf("Running first scann \n", )
	diff, err := utils.GetGitDiff()
	if err != nil {
		log.Println("Error al obtener el diff de Git:", err)
		return
	}

	err = utils.AppendToFile(diff)
	if err != nil {
		log.Println("Error al escribir en el archivo:", err)
		return
	}
	tete, err := utils.GetJiraTicketID()
	if err != nil {
		log.Println("Error al obtener branch:", err)
	}
	log.Println("el tete", tete)
	ticketID := "PYT-1180" // Reemplaza esto con el ID real del ticket
	output, err := utils.SendGetHttpRequest(ticketID)

	if err != nil {
		log.Println("Error al obtener información del ticket de Jira:", err)
		return
	}

	log.Printf("Output del ticket Jira %s: %s\n", ticketID, string(output))

	err = utils.SendHttpPost(diff)
	if err != nil {
		log.Println("Error al enviar la petición HTTP:", err)
	}
}