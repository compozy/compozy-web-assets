// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3af5c855c42e9a38dd265125076ec4a3de63cbe9ffdeb65396cc002437b39181"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6dab45937317aa61e1b52e25b0d508d35023eb2e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
