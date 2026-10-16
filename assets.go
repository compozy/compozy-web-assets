// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f37a9f257b6ea16a0babfec20a60e61d399793aba6edb1d902515ae50aecb267"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ffd666bdb77f4ee43ab9fc9dd1b273332ac0e0ad"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
