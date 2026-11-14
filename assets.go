// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "19d94dfe7b23114929d39ab21c05c2ce6840c04a0af067768d8b708a1e763e79"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "926a2dd560de8075a04f4b76935667bde4c0c34d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
