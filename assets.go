// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4a5534bd8308b621dca97bbb772b120eb6f10e7715ee87cdc5162f953e1b29d6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a7a89709d749eb1c803135afb0c0cf81d3952575"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
