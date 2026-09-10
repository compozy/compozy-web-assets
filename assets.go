// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cbacfa9d0e71903721f92d35a9be24e415d1c79fe7dd2ae78c197cb2e975d353"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0e820dc0e2e38d958ba06b57a941631599ef092b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
