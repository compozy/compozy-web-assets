// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "18411e3acbbf35768dd8859007e44b92cfcff2f34ed52110fce8bc65c10b571d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c43ff82e5c5b9cb8a809829bb5aae88c540e0051"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
