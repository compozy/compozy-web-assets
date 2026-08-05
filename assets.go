// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4bed4e835cf0d3f6f2f1579c974c93dc27a154c290d1e845fa60d619036420fe"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5d236a41f4c02968edad62de81f2913b244143d3"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
