// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5c4fa67fbd40e7efb493a46e2f5e117119a50a093b29c3ff2c7a4bc7d1c78976"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "19fe37c64b179e4daba5db479d9066fd952d65cd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
