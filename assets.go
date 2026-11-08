// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "db469573bc99698751ab7398b3509f1d892d3150df0d660a6f72db4e16ac6f4f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0b30658da20c3da3161a4fd4e06d0e29894c014f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
