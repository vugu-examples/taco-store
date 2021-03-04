package main

import (
	"flag"
	"github.com/vugu-examples/taco-store/internal/handlers"
	"github.com/vugu-examples/taco-store/internal/memstore"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	baseDir := flag.String("base", "", "Base directory to look for files, default uses current directory")
	fullBaseDir, _ := filepath.Abs(*baseDir)

	mux := http.NewServeMux()
	hl := handlers.NewFrontendHandler(fullBaseDir)
	mem := memstore.NewTacoStore()

	hl2 := handlers.NewTacoStoreAPIHandler(mem)
	mux.Handle("/", hl)
	mux.Handle("/api/taco-list", hl2)

	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)
	log.Fatal(http.ListenAndServe(l, mux))
}
