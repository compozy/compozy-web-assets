// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e038703056f49f035e31e1b3dbbc0c9343ad0e66b154a1bc256dadfc5cc52311"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "31e35e1b838eefa28c67f8049f966c78bb82cfd1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
