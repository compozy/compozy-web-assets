// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "06bd7080e36de5901e1cf86055d5f04203a7ebaf3368b1f258ddd4b2c88fd268"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6a7ea3c99fd5bceb5e8a878123b594db5faa89d7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
