// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "430e8646a161d593b0f8eef532c66632d48f713ad8dd57aba68b03b9fbfc2d5c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "27daadb355da28939f1956b6c44d5368920f8acf"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
