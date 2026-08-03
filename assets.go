// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "94f99930ca32583f03ff954a11114f54ff6b6f4b228e1e2039e4ec6a4f7ba626"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cff6dae4834e056b6a0a484c43710d8341c2a47f"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
