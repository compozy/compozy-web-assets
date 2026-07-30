// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "633a5898f9ad2086f4cea7627df15f689e93d054ad3ae7f6d47a75f1388f5eec"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7285bf3cd8bab830e0f0879e0b25f3b58715c930"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
