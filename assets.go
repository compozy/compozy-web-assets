// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f1ff475e5437ba75f9ccc517f2c54315a14dffa6b8068ffce2ce9135cb2feb7f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "688aa99590f6f8a2633bdca035f2094ffa80b33e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
