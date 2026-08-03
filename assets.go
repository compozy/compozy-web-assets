// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "be8e80172799eef46692b721460e762e3f90e9333dc67d7cabb62e0c492e3a8d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5ebd78dd8fa39db7badfa7b584a98c735d5f5446"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
