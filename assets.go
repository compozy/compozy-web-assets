// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "da3dcbd449f04e3199433ebeb71827c5305894dc3129c63bad53a56881c6745d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ff60d4ea131c3a5a0eb10e09fc036418459c6072"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
