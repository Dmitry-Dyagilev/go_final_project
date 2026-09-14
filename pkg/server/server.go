package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/Dmitry-Dyagilev/final-project/pkg/api"
)

func Run(port string) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	rootPath := filepath.Dir(exePath)
	webDir := filepath.Join(rootPath, "web")

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	api.Init()
	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}
	return nil
}
