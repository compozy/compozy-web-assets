// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5c4fa67fbd40e7efb493a46e2f5e117119a50a093b29c3ff2c7a4bc7d1c78976"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f78b6de68b3448f130dfdec93652447c3a689006"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
