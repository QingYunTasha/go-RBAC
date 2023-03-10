package usecasedomain

import (
	"context"
	RepoDomain "go-authorization/domain/repository"
)

type PermissionUsecase interface {
	GetAll(ctx context.Context) ([]RepoDomain.Permission, error)
	GetByResource(ctx context.Context, resourceName string) ([]RepoDomain.Permission, error)
	GetByRole(ctx context.Context, permission *RepoDomain.Role) ([]RepoDomain.Permission, error)
	Create(ctx context.Context, permission *RepoDomain.Permission) error
	Delete(ctx context.Context, resourceName string, action string) error
}
