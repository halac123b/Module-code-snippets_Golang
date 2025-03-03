package github

import (
	"github.com/trufflesecurity/trufflehog/v3/pkg/analyzer/config"
)

var _ analyzers.Analyzer = (*Analyzer)(nil)

type Analyzer struct {
	Cfg *config.Config
}
