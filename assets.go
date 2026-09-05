// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0a3138127a854ffdc4e41cdb9e52974ba9ea1e292f95abe1fce383e3597ce313"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0e6cff50d51574573deeb188464a09014417d46d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
