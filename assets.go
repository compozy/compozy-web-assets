// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4bed4e835cf0d3f6f2f1579c974c93dc27a154c290d1e845fa60d619036420fe"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8be5b224f09b0e5aad31abeb7a52edb4b89b933c"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
