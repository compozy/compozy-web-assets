// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "79150de3c59fc0d96d52cf5abb66ad97d8a7ba35d8488a9d3c6124f7332d7002"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1c39d4aeddede40df831974f67b12174a98529f0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
