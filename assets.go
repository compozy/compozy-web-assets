// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "79a818cffac35722d2e62b06af401d7c1b718b81caaf79115afd89b95a445012"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f2ff994b029c1142961dde70ee97418e93d70169"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
