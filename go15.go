// +build go1.5
// +build !go1.4

package gotool

import (
	"go/build"
	"path/filepath"
	"runtime"
)

var gorootSrcPkg = filepath.Join(runtime.GOROOT(), "src", "pkg")

func shouldIgnoreImport(p *build.Package) bool {
	return p == nil || len(p.InvalidGoFiles) == 0
}
