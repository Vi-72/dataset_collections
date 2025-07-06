package http

import (
	"dataset-collections/internal/core/application/usecases/commands"
	"dataset-collections/internal/core/application/usecases/queries"
	"dataset-collections/internal/pkg/errs"
)

type ApiHandler struct {
	startImportHandler    commands.StartImportCommandHandler
	listPopulationHandler queries.ListPopulationQueryHandler
}

func NewApiHandler(
	startImportHandler commands.StartImportCommandHandler,
	listPopulationHandler queries.ListPopulationQueryHandler,
) (*ApiHandler, error) {
	if startImportHandler == nil {
		return nil, errs.NewValueIsRequiredError("startImportHandler")
	}
	if listPopulationHandler == nil {
		return nil, errs.NewValueIsRequiredError("listPopulationHandler")
	}
	return &ApiHandler{
		startImportHandler:    startImportHandler,
		listPopulationHandler: listPopulationHandler,
	}, nil
}
