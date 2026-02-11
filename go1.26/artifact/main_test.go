package main

import (
	"path"
	"testing"
)

// START OMIT
func TestReport(t *testing.T) {
	reportFile := path.Join(t.ArtifactDir(), "report.csv") // HL
	t.Logf("report at: %q", reportFile)

	// Test report generation
}

// END OMIT
