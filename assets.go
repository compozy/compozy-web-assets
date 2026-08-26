// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "aca636b026451008ba1c1c875b217ccd41b478c8bc2c428a1a650b88b1682a86"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a67cf361965a81b7e5f109ffab7d54251b160be0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
