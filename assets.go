// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "aca636b026451008ba1c1c875b217ccd41b478c8bc2c428a1a650b88b1682a86"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8753ad9e9c046f2c40df2eaec34f4d8db153d96c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
