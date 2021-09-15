package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID        uuid.UUID
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Active    bool
	Metadata  map[string]interface{}
	// tableName struct{}
}

func (m *ModelBase) BeforeUpdate(ctx context.Context) (context.Context, error) {
	m.UpdatedAt = time.Now()
	return ctx, nil
}
