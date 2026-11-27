// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "44dbae6af625be425aaa464a0bc4b3f2ed6855edb7851f9b8504a78fbec44584"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6ae1580efae4ddd24e556ad0e2d28e1d87ecbc0f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
