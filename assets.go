// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "835e889edbfc5b707408e5015368a1134fa4c6a9b8b58927d42b64b20c2a7a15"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "854fa2d3611f621a707396db66e64c0003cd9e04"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
