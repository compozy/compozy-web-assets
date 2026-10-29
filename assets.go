// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f81392634fb30ef927a209369e580c1d90bd8a744baa0af888c047019f782f06"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "61792c15693f592098767f59638689a020c2b50a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
