package service

import (
	"context"

	"github.com/Ivanrizkys/go-api/api/model/web"
)

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
}
