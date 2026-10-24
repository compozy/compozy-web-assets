// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3cfea221ae2de9c63a3c89c64869c409981d5ea71b609b6ea2bd57d6cb8bd337"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e342ac705822d5c569acdfb756deed0f97461902"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
