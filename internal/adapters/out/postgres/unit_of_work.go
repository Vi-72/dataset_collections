package postgres

import (
	"context"
	"dataset-collections/internal/adapters/out/postgres/importjobrepo"
	"dataset-collections/internal/adapters/out/postgres/populationrepo"
	"dataset-collections/internal/core/ports"
	"dataset-collections/internal/pkg/errs"
	"gorm.io/gorm"
)

var _ ports.UnitOfWork = &UnitOfWork{}

type UnitOfWork struct {
	tx                   *gorm.DB
	db                   *gorm.DB
	populationRepository ports.PopulationRepository
	importJobRepository  ports.ImportJobRepository
}

func NewUnitOfWork(db *gorm.DB) (ports.UnitOfWork, error) {
	if db == nil {
		return nil, errs.NewValueIsRequiredError("db")
	}

	uow := &UnitOfWork{db: db}

	popRepo, err := populationrepo.NewRepository(uow)
	if err != nil {
		return nil, err
	}
	uow.populationRepository = popRepo

	importJobRepo, err := importjobrepo.NewRepository(uow)
	if err != nil {
		return nil, err
	}
	uow.importJobRepository = importJobRepo

	return uow, nil
}

func (u *UnitOfWork) Tx() *gorm.DB {
	return u.tx
}

func (u *UnitOfWork) Db() *gorm.DB {
	return u.db
}

func (u *UnitOfWork) InTx() bool {
	return u.tx != nil
}

func (u *UnitOfWork) Begin(ctx context.Context) {
	u.tx = u.db.WithContext(ctx).Begin()
}

func (u *UnitOfWork) Rollback() error {
	if u.tx != nil {
		err := u.tx.Rollback().Error
		u.tx = nil
		return err
	}
	return nil
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	if u.tx == nil {
		return errs.NewValueIsRequiredError("cannot commit without transaction")
	}

	if err := u.tx.WithContext(ctx).Commit().Error; err != nil {
		return err
	}
	u.tx = nil
	return nil
}

// Геттеры репозиториев
func (u *UnitOfWork) PopulationRepository() ports.PopulationRepository {
	return u.populationRepository
}

func (u *UnitOfWork) ImportJobRepository() ports.ImportJobRepository {
	return u.importJobRepository
}
