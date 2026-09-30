// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5c17cb7a6c4b5e2f3c71a55986f9dcba9dd130255993618e88fca365622b2cfe"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a15f541f93e1da6d5d1678b22e2312101a1cda4f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
