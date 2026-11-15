// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "22342343c54ad2976d5677ae143f4a4406ae914d91dcbfc766bd3cf0e2cd5f63"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b8556d349e49cbe8ba95129092f25e1ef6c0da2a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
