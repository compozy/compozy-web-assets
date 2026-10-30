// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "79a818cffac35722d2e62b06af401d7c1b718b81caaf79115afd89b95a445012"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0c189b7ee1f5b3d97098053b3433513dcc8ac724"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
