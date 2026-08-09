// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e1863f5547bf29b2d475bdfd04aaf667a82130ce8a361a9b4e61b781bed36da5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "69db9ed7e13edcae2fcc0115da195021db89510f"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
