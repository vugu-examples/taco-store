package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// NewFrontendHandler returns a new FrontendHandler.
func NewFrontendHandler(baseDir string) *FrontendHandler {
	return &FrontendHandler{
		BaseDir: baseDir,
	}
}

// FrontendHandler deals with serving up the thing needed for Vugu.
type FrontendHandler struct {
	BaseDir string // absolute path of project
	DevMode bool   // true if we should be automatically rebuilding the frontend (only on local)
}

// ServeHTTP implements http.Handler but only writes a response for files we serve.
func (h *FrontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	distDir := http.Dir(filepath.Join(h.BaseDir, "dist"))
	log.Printf("FrontendHandler starting with dist dir: %s", distDir)

	buildFrontend := func() (ok bool) {

		cmd := exec.Command("sh", "build-wasm.sh")

		if runtime.GOOS == "windows" {
			cmd = exec.Command("sh", "build-wasm.bat")
		}

		cmd.Env = append(os.Environ(), "GO111MODULE=auto")

		cmd.Dir = h.BaseDir
		b, err := cmd.CombinedOutput()
		log.Printf("go run build-frontend.go - err: %v\n%s", err, b)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(500)
			fmt.Fprintf(w, "build-frontend.go error: %v\n%s", err, b)
			return false
		}
		return true
	}
	ext := path.Ext(r.URL.Path)

	if ext != "" {
		f, err := distDir.Open(r.URL.Path)
		if os.IsNotExist(err) {
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("File error: %v", err), 500)
			return
		}
		defer f.Close()

		st, err := f.Stat()
		if err != nil {
			http.Error(w, fmt.Sprintf("File stat error: %v", err), 500)
			return
		}
		_ = st

		// manually handle some mime types
		switch ext {
		case ".wasm":
			w.Header().Set("Content-Type", "application/wasm")

		case ".css":
			w.Header().Set("Content-Type", "text/css")

		}
		http.ServeContent(w, r, r.URL.Path, st.ModTime(), f)
		return
	}

	if !buildFrontend() {
		return
	}
	h.serveIndex(w, r)
	return
}

func (h *FrontendHandler) serveIndex(w http.ResponseWriter, r *http.Request) {
	in := strings.NewReader(indexHTML)
	http.ServeContent(w, r, "tacostore", startupTime, in)
}

var startupTime = time.Now()

var indexHTML = `<!doctype html>
<html>
<head>
<meta charset="utf-8"/>
    <link rel="stylesheet" href="/css/main.css">
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
   <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<!-- styles -->
</head>
<body>
<div id="vugu_mount_point">
<img style="position: absolute; top: 50%; left: 50%;" src="https://cdnjs.cloudflare.com/ajax/libs/galleriffic/2.0.1/css/loader.gif">
</div>
<script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script> <!-- MS Edge polyfill -->
<script src="/wasm_exec.js"></script>
<!-- scripts -->
<script>
var wasmSupported = (typeof WebAssembly === "object");
if (wasmSupported) {
	if (!WebAssembly.instantiateStreaming) { // polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}
	var mainWasmReq = fetch("/index.wasm").then(function(res) {
		if (res.ok) {
			const go = new Go();
			WebAssembly.instantiateStreaming(res, go.importObject).then((result) => {
				go.run(result.instance);
			});		
		} else {
			res.text().then(function(txt) {
				var el = document.getElementById("vugu_mount_point");
				el.style = 'font-family: monospace; background: black; color: red; padding: 10px';
				el.innerText = txt;
			})
		}
	})
} else {
	document.getElementById("vugu_mount_point").innerHTML = 'This application requires WebAssembly support.  Please upgrade your browser.';
}
</script>
</body>
</html>`
