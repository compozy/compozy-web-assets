// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1444980648d8ead4f3aa432ce5bc365be38585ea1d75fe315a78a37257b87e9c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b0eff1326a0962a0494c0f7d3953dff844b5795e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
