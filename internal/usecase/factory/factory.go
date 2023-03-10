package factory

import (
	UsecaseDomain "go-authorization/domain/usecase"
	OrmFactory "go-authorization/internal/repository/database/factory"
	Usecase "go-authorization/internal/usecase"
)

type UsecaseRepository struct {
	Level      UsecaseDomain.LevelUsecase
	Permission UsecaseDomain.PermissionUsecase
	Resource   UsecaseDomain.ResourceUsecase
	Role       UsecaseDomain.RoleUsecase
	User       UsecaseDomain.UserUsecase
	Core       UsecaseDomain.CoreUsecase
}

func InitUsecaseRepository(orm *OrmFactory.OrmRepository) *UsecaseRepository {
	return &UsecaseRepository{
		Level:      Usecase.NewLevelUsecase(orm),
		Permission: Usecase.NewPermissionUsecase(orm),
		Resource:   Usecase.NewResourceUsecase(orm),
		Role:       Usecase.NewRoleUsecase(orm),
		User:       Usecase.NewUserUsecase(orm),
		Core:       Usecase.NewCoreUsecase(orm),
	}
}
