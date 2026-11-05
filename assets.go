// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e55583a5c0f7eac70407b3fe920998fe93e2c29cc43b59a131b1e0a1b9016d85"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "03911a784660310c3d50f8a4655f0503303f7cf5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
