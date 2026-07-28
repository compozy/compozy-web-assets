// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "905ad583f77c024b6467828fd457bb193be34492e0451977c2f9f2456eacd6e7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "87195b69966e38ff9fedbca1f4a114a1a72a7a6f"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
