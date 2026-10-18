// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "066aa3cdf3617cc584045b9747ae6d128c5b404e9ec8999625954484fc55731d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "224dc50216513d44564f8b71298046a840b59cb8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
