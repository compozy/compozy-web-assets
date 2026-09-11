// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a330252ae2e073c594d17cff523ad3f20b28f012712733aa6686e8cbce58fc7e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "be00def6dbce32143b84ace452ff899f5acca33b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
