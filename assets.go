// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2858a907de56c07accc0a4d4d9c232cacc409ebaf9f6dd9f4221d104a250099e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f213aa2f332ff876d2a93ea25fbf602bdc906a2c"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
