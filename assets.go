// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e55583a5c0f7eac70407b3fe920998fe93e2c29cc43b59a131b1e0a1b9016d85"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3f616b82626020aba7f6d00c7a91965ed773a5fe"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
