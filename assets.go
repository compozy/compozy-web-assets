// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "94520f40d29097d92bf8ac84bf70eb2f0ba297dd0541245a3fe417fdc3dcfdb8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "af148a45c1bdc77f7fbab6db1aee3e30f480abf1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
