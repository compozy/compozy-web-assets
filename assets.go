// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "00766d52a4c2c932fb3e5d72975b0d08932e04e1ce70e28732c17878f9160353"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ae8bca8dff1ae35c9ed07eeacb836bbf27738113"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
