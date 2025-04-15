package packageC

import (
	"log/slog"

	"github.com/ThisaraWeerakoon/dynamic-go-logger/pkg/loggerfactory"
)

const (
	componentName = "packageC"
)

type C struct {
	field_1 string
	field_2 int
	logger  *slog.Logger
}

func NewC(field_1 string, field_2 int) *C {
	c := &C{field_1: field_1, field_2: field_2}
	c.logger = loggerfactory.GetLogger(componentName, c)
	return c
}

func (c *C) UpdateLogger() {
	c.logger = loggerfactory.GetLogger(componentName, c)
}

func (c *C) ShowLogs() {
	c.logger.Debug("Debug message", "field_1", c.field_1, "field_2", c.field_2, "package", componentName)
	c.logger.Info("Info message", "field_1", c.field_1, "field_2", c.field_2, "package", componentName)
	c.logger.Warn("Warn message", "field_1", c.field_1, "field_2", c.field_2, "package", componentName)
	c.logger.Error("Error message", "field_1", c.field_1, "field_2", c.field_2, "package", componentName)
}