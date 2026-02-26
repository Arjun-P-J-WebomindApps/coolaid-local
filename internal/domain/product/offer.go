package product

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//
// ============================================================
// 🔹 OFFER (Single Table Lifecycle)
// ============================================================
//

//
// 🔹 GET
//

func (s *Service) getOfferByPartNo(
	ctx context.Context,
	Q Queries,
	partNo string,
) (*ProductOffer, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.GetOfferByPartNo(ctx, partNo)
	if err != nil {
		return nil, ErrOfferNotFound
	}

	return mapOfferRowToModel(row), nil
}

// 🔹 CREATE
func parseDate(v string) (time.Time, error) {
	return time.Parse("2006-01-02", v)
}

func (s *Service) createOffer(
	ctx context.Context,
	Q Queries,
	partNo string,
	input CreateProductOfferInput,
) (*ProductOffer, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	now := time.Now()

	var startDate, endDate string

	if input.IsOfferActive {

		start, err := parseDate(input.StartDate)
		if err != nil {
			return nil, ErrInvalidOfferData
		}

		end, err := parseDate(input.EndDate)
		if err != nil {
			return nil, ErrInvalidOfferData
		}

		if !start.Before(end) {
			return nil, ErrInvalidOfferDateRange
		}

		startDate = input.StartDate
		endDate = input.EndDate

	} else {

		// 🔥 Default fallback dates when offer inactive
		// Use today's date for both
		today := now.Format("2006-01-02")
		startDate = today
		endDate = today
	}

	row, err := Q.CreateOffer(ctx, CreateProductOfferParams{
		ID:            uuid.NewString(),
		PartNo:        partNo,
		IsOfferActive: input.IsOfferActive,
		StartDate:     startDate,
		EndDate:       endDate,
		AcTrader:      input.AcTrader,
		MultiBrand:    input.MultiBrand,
		Autotrader:    input.Autotrader,
		AcWorkshop:    input.AcWorkshop,
		CreatedAt:     now,
		UpdatedAt:     now,
	})
	if err != nil {
		return nil, err
	}

	return mapOfferRowToModel(row), nil
}

//
// 🔹 UPDATE
//

func (s *Service) updateOffer(
	ctx context.Context,
	Q Queries,
	partNo string,
	input UpdateProductOfferInput,
) (*ProductOffer, error) {

	if partNo == "" {
		return nil, ErrInvalidInput
	}

	row, err := Q.UpdateOffer(ctx, UpdateProductOfferParams{
		PartNo:        partNo,
		IsOfferActive: input.IsOfferActive,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		AcTrader:      input.AcTrader,
		MultiBrand:    input.MultiBrand,
		Autotrader:    input.Autotrader,
		AcWorkshop:    input.AcWorkshop,
	})

	if err != nil {
		return nil, err
	}

	return mapOfferRowToModel(row), nil
}

//
// 🔹 DELETE
//

func (s *Service) deleteOffer(
	ctx context.Context,
	Q Queries,
	partNo string,
) error {

	if partNo == "" {
		return ErrInvalidInput
	}

	_, err := Q.GetOfferByPartNo(ctx, partNo)
	if err != nil {
		return ErrOfferNotFound
	}

	return Q.DeleteOffer(ctx, partNo)
}
