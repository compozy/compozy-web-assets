// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c9cdf405e9d3306458f37fd034ba550197fde0c2706b6dd4dc7787e5bc1d05b6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c38ba0fac69aa657140fc578067d2c538b18ec10"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
