package pkgs

import (
	"fmt"
	"go.mau.fi/whatsmeow"
	"time"
	"whatsappWiz/data"
)

func ValidateTime(client *whatsmeow.Client) {
	// Cargar el huso horario de Caracas
	loc, _ := time.LoadLocation("America/Caracas")

	// Ejecutar el programa indefinidamente
	for {
		// Obtener la hora actual en el huso horario local
		now := time.Now().In(loc)

		// Calcular el tiempo hasta la próxima ejecución a las 12pm
		next12pm := time.Date(now.Year(), now.Month(), now.Day(), 12, 1, 0, 0, loc)
		if now.After(next12pm) {
			next12pm = next12pm.Add(24 * time.Hour) // Avanzar al siguiente día si ya pasó la hora
		}
		duration12pm := next12pm.Sub(now)

		// Calcular el tiempo hasta la próxima ejecución a las 3pm
		next3pm := time.Date(now.Year(), now.Month(), now.Day(), 15, 11, 0, 0, loc)
		if now.After(next3pm) {
			next3pm = next3pm.Add(24 * time.Hour) // Avanzar al siguiente día si ya pasó la hora
		}
		duration3pm := next3pm.Sub(now)

		// Calcular el tiempo hasta la próxima ejecución a las 6pm
		next6pm := time.Date(now.Year(), now.Month(), now.Day(), 18, 30, 0, 0, loc)
		if now.After(next6pm) {
			next6pm = next6pm.Add(24 * time.Hour) // Avanzar al siguiente día si ya pasó la hora
		}
		duration6pm := next6pm.Sub(now)

		// Calcular el tiempo de espera más corto entre las tres horas
		minDuration := min(duration12pm, duration3pm, duration6pm)

		// Esperar hasta la próxima ejecución
		fmt.Printf("Próxima ejecución a las %s\n", now.Add(minDuration).Format("15:04:05"))
		SendMessageToLog(client, data.GroupLog().JID, data.GroupLog().Name,
			"⌚ *Próxima ejecución a las "+now.Add(
				minDuration).Format(
				"15:04:05")+"*")
		time.Sleep(minDuration)

		// Ejecutar la función enviarMensajes

		for _, group := range GetGroups() {
			SendMessage(client, group.JID, group.Name)
		}

		SendMessageToLog(client, data.GroupLog().JID, data.GroupLog().Name, "🚀 *Mensajes enviados*")
	}
}

// Función auxiliar para encontrar el tiempo mínimo entre tres duraciones
func min(durations ...time.Duration) time.Duration {
	minDuration := durations[0]
	for _, d := range durations {
		if d < minDuration {
			minDuration = d
		}
	}
	return minDuration
}
