// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2bc86ddf03072f700d1866b29a396c7fbb659efded3c5d3c9523332a12965bfb"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "96f95549d40832240fe074847ee3324b980a251a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
