package bmanalysis

import (
	"os"
	"testing"
)

func TestNotebookGeneration(t *testing.T) {
	s := CreateAnalysisTemplate()
	s.ProjectLists = []string{"proj_zedboard_1", "proj_zedboard_2"}

	f, err := os.Create("test.ipynb")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	if r, err := s.WritePython(); err != nil {
		t.Error(err)
	} else {
		f.WriteString(r)
	}
}
