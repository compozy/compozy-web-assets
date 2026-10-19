// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "88ee3de348efa15e774b413af50df3b4a7cff65139aa54c8c523838b458e7aae"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a98ec6ca3d057d652b08de4ef6eaf39b700dabd3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
