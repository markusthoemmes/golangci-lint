package golinters

import (
	"github.com/gordonklaus/ineffassign/pkg/ineffassign"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

const ineffassignName = "ineffassign"

func NewIneffassign() *goanalysis.Linter {
	return goanalysis.NewLinter(
		ineffassignName,
		"Detects when assignments to existing variables are not used",
		[]*analysis.Analyzer{ineffassign.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
