// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c9cdf405e9d3306458f37fd034ba550197fde0c2706b6dd4dc7787e5bc1d05b6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "107890a045055d8bb8975936789e81d6a63321f8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
