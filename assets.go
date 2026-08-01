// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8701fe10388013d8527a138038c69f2a4e6ecef02b104468cd0f3240b1d65db2"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "144d3250703bf02faec2ab64bc4260f1d4484fde"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
