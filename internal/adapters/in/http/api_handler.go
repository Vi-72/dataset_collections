package http

import (
	"dataset-collections/internal/core/application/usecases/commands"
	"dataset-collections/internal/core/application/usecases/queries"
	"dataset-collections/internal/pkg/errs"
)

type ApiHandler struct {
	startImportHandler        commands.StartImportCommandHandler
	listPopulationHandler     queries.ListPopulationQueryHandler
	getImportJobStatusHandler queries.GetImportJobStatusQueryHandler
}

func NewApiHandler(
	startImportHandler commands.StartImportCommandHandler,
	listPopulationHandler queries.ListPopulationQueryHandler,
	getImportJobStatusHandler queries.GetImportJobStatusQueryHandler,
) (*ApiHandler, error) {
	if startImportHandler == nil {
		return nil, errs.NewValueIsRequiredError("startImportHandler")
	}
	if listPopulationHandler == nil {
		return nil, errs.NewValueIsRequiredError("listPopulationHandler")
	}
	if getImportJobStatusHandler == nil {
		return nil, errs.NewValueIsRequiredError("getImportJobStatusHandler")
	}
	return &ApiHandler{
		startImportHandler:        startImportHandler,
		listPopulationHandler:     listPopulationHandler,
		getImportJobStatusHandler: getImportJobStatusHandler,
	}, nil
}
