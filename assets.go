// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1ed37958b0fdf3331910f93743a390a556a0f884b636bc8a35b55225235fb965"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bf5fc3c385eb76a1c981ff73588d6518ff0a3c46"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
