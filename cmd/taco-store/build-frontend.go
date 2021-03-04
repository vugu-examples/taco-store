// +build ignore

package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/vugu/vugu/distutil"
)

func main() {

	var err error
	start := time.Now()

	// find and copy wasm_exec.js
	distutil.MustCopyFile(distutil.MustWasmExecJsPath(), "dist/wasm_exec.js")

	// run go generate
	cmd := exec.Command("go", "generate")
	cmd.Dir, err = filepath.Abs("./ui")
	distutil.Must(err)
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s", out)
	distutil.Must(err)

	log.Printf("build-frontend.go complete in %v", time.Since(start))
}
