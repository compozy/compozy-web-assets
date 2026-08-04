// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a75ff7aed1bc9a363a9fec1317fc32abb0f33dfbb6ad5d5b7233dcc8f926bc7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "170359ac10680cac7cb1d0723bfc21cae374a214"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
