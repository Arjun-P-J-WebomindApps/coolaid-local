package cronjob

import (
	"context"
	"log"

	ticketservice "github.com/webomindapps-dev/coolaid-backend/internal/domain/ticket"
)

type TicketCleanupJob struct {
	TicketService *ticketservice.Service
}

func (j *TicketCleanupJob) Run() {
	log.Println("[CRON] Ticket cleanup started")

	if err := j.TicketService.CleanupExpiredTickets(context.Background()); err != nil {
		log.Printf("[CRON] Ticket cleanup failed: %v", err)
		return
	}

	log.Println("[CRON] Ticket cleanup finished")
}
