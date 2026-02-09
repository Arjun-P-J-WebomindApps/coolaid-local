package cronjob

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler(loc *time.Location) *Scheduler {
	return &Scheduler{
		cron: cron.New(
			cron.WithLocation(loc),
			cron.WithSeconds(), // optional but future-proof
		),
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("[CRON] Scheduler started")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Println("[CRON] Scheduler stopped")
}

func (s *Scheduler) Add(spec string, job func()) {
	if _, err := s.cron.AddFunc(spec, job); err != nil {
		log.Fatalf("[CRON] Failed to schedule job: %v", err)
	}
}
