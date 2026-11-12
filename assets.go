// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "75d954711e92886712d0021821b29eb68a4e72a8bed46c8f7127db6b96787868"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "86fcd83d41b0af342ab86b0c2addd9fc987a5e42"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
