// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "430e8646a161d593b0f8eef532c66632d48f713ad8dd57aba68b03b9fbfc2d5c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5a9d801d363d9e80369295005d8999a1aeb0e0db"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
