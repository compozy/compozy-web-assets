// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "430e8646a161d593b0f8eef532c66632d48f713ad8dd57aba68b03b9fbfc2d5c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "69536f7a153f641710771532e6f29ac6480868ce"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
