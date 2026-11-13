// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3e1f5742927332269fa88ad57a104038112b4dded9e2e5eabb936791e739ebb7"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "85db4c5261f5da8b8213f99a565753a83d27204c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
