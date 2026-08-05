// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ad540a15d7040d0720a339d7e80cb510a80eb6d2de44ae5bfb8b8229fceec25f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5b81479b6fd104b42b3a72aa2decd7b215a42dfd"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
