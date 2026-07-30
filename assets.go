// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4f7175b892f80b0d08db8351c0f67ee08ea5a7e4d31d49d07d0713c20f5c845f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0fdbdf28b0d09ebb5f073beb71d984c05d97dce4"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
