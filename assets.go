// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "6a501a22cc280341088fbbdf5fe908a01822c15a62d91d41cb89f516fd5cda03"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1c6f5d9ec4e09a179713ecaa196e97179a79651f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
