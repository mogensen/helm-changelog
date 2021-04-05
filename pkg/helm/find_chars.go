package helm

import (
	"os"
	"path/filepath"
)

// FindCharts locates all "Chart.yaml" files in the current directory, and all sub-directories
func FindCharts(chartSearchDir string) ([]string, error) {

	fileList := []string{}
	err := filepath.Walk(chartSearchDir, func(path string, f os.FileInfo, err error) error {

		fileName := filepath.Base(path)
		if fileName == "Chart.yaml" {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}
