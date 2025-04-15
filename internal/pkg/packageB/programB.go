package packageB

import (
	"log/slog"
	"github.com/ThisaraWeerakoon/dynamic-go-logger/pkg/loggerfactory"
	"github.com/ThisaraWeerakoon/dynamic-go-logger/internal/pkg/packageC"
)

var (
	componentName = "packageB"
)	

type B struct{
	field_1 string
	field_2 int
	logger  *slog.Logger
}

func NewB(field_1 string, field_2 int) *B {
	b := &B{field_1: field_1, field_2: field_2}
	b.logger = loggerfactory.GetLogger(componentName, b)
	return b
}
func (b *B) UpdateLogger() {
	b.logger = loggerfactory.GetLogger(componentName, b)
}

func (b *B) ShowLogs() {
	b.logger.Debug("Debug message", "field_1", b.field_1, "field_2", b.field_2, "package", componentName)
	b.logger.Info("Info message", "field_1", b.field_1, "field_2", b.field_2, "package", componentName)
	b.logger.Warn("Warn message", "field_1", b.field_1, "field_2", b.field_2, "package", componentName)
	b.logger.Error("Error message", "field_1", b.field_1, "field_2", b.field_2, "package", componentName)
}

func (b *B) InitiatePackageC() {
	c := packageC.NewC(b.field_1, b.field_2)
	c.ShowLogs()
}

