package packageA

import (
	"log/slog"

	"github.com/ThisaraWeerakoon/dynamic-go-logger/pkg/loggerfactory"
)

const (
	componentName = "packageA"
)

type A struct {
	field_1 string
	field_2 int
	logger  *slog.Logger
}

func NewA(field_1 string, field_2 int) *A {
	a := &A{field_1: field_1, field_2: field_2}
	a.logger = loggerfactory.GetLogger(componentName, a)
	return a
}

func (a *A) UpdateLogger() {
	a.logger = loggerfactory.GetLogger(componentName, a)
}

func (a *A) ShowLogs() {
	a.logger.Debug("Debug message", "field_1", a.field_1, "field_2", a.field_2, "package", componentName)
	a.logger.Info("Info message", "field_1", a.field_1, "field_2", a.field_2, "package", componentName)
	a.logger.Warn("Warn message", "field_1", a.field_1, "field_2", a.field_2, "package", componentName)
	a.logger.Error("Error message", "field_1", a.field_1, "field_2", a.field_2, "package", componentName)
}

