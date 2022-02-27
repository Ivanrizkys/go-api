package repository

import (
	"context"
	"database/sql"

	"github.com/Ivanrizkys/go-api/api/model/domain"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
