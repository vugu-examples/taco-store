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
	log.Println(hl)
	mem := memstore.NewTacoStore()

	hl2 := handlers.NewTacoStoreAPIHandler(mem)
	mux.Handle("/", hl)
	mux.Handle("/api/taco-list", hl2)

	l := "127.0.0.1:8844"
	//log.Printf("Starting HTTP Server at %q", l)
	log.Fatal(http.ListenAndServe(l, mux))
	//	go func() {
	//	log.Fatal(http.ListenAndServe(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		r.URL.Scheme = "http"
	//		if r.URL.Host == "" {
	//			r.URL.Host = r.Host
	//		}
	//		ws.ServeHTTP(w, r)
	//	})))
	//}()
	//mem := memstore.NewTacoStore()

	//wc := devutil.NewWasmCompiler().SetDir(".")
	//mux := devutil.NewMux()
	////rh := http.RedirectHandler("http://example.org", 307)
	//mux.Exact("/api/taco-list",hl.GetTacoList())
	//mux.Match(devutil.NoFileExt, devutil.DefaultAutoReloadIndex.Replace(
	//	`<!-- styles -->`,
	//	`<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
	//		  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">`))
	//mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	//mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	//mux.Default(devutil.NewFileServer().SetDir("."))
	//
	//log.Fatal(http.ListenAndServe(l, mux))
}
