package cronjob

import (
	"time"

	ticketservice "github.com/webomindapps-dev/coolaid-backend/internal/domain/ticket"
)

func RegisterAll(ticketSvc *ticketservice.Service) *Scheduler {
	loc, _ := time.LoadLocation("Asia/Kolkata")

	scheduler := NewScheduler(loc)

	// Run every day at 12:00 PM
	ticketCleanup := &TicketCleanupJob{
		TicketService: ticketSvc,
	}

	scheduler.Add("0 0 12 * * *", ticketCleanup.Run)

	return scheduler
}
