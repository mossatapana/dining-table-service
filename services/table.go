package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type TableServiceImpl struct {
	Tables  []bool
	Booking map[int]string
}

func NewTableService() *TableServiceImpl {
	return &TableServiceImpl{}
}

func (t TableServiceImpl) Create(ctx context.Context, noOfTables int) error {
	if len(t.Tables) > 0 {
		return errors.New("table has been created")
	}
	for i := 0; i < noOfTables; i++ {
		t.Tables[i] = false
	}
	return nil
}

func (t TableServiceImpl) Reserve(ctx context.Context, noOfCus int) error {
	return nil
}

func generateBookingID(ctx context.Context) (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", nil
	}

	return uuid.String(), nil
}
