// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "db469573bc99698751ab7398b3509f1d892d3150df0d660a6f72db4e16ac6f4f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "43a238496344c8277c7237fbbe143a3d8c5feeb0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
