// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "fc26fdc6ea899047801003debf44d135fb5732cf17facfccdc601b431133e966"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "19b624863ab87cfb3c500c88f5033f072de956a0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
