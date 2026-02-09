package ticketservice

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

// CleanupExpiredTickets contains ONLY business logic
func (s *Service) CleanupExpiredTickets(ctx context.Context) (err error) {
	oplog.Info(ctx, "ticket cleanup service started")

	tx, err := s.DB.SqlDB.BeginTx(ctx, nil)
	if err != nil {
		oplog.Error(ctx, err)
		return err
	}

	// rollback safety
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				oplog.Error(ctx, rbErr)
			}
		}
	}()

	qtx := s.DB.Queries.WithTx(tx)

	// 1. Find & delete old pending tickets
	tickets, err := qtx.DeletePendingTicketsOlderThan2Days(ctx)
	if err != nil {
		oplog.Error(ctx, err)
		return err
	}

	// 2. Delete ticket items
	for _, ticket := range tickets {
		if _, err = qtx.DeleteTicketItemsByTicketId(ctx, ticket.ID); err != nil {
			oplog.Error(ctx, err)
			return err
		}
	}

	// 3. Commit transaction
	if err = tx.Commit(); err != nil {
		oplog.Error(ctx, err)
		return err
	}

	oplog.Info(ctx, "ticket cleanup completed", "deleted_count=", len(tickets))
	return nil
}
