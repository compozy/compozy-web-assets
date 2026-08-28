// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d22e66c1aa2ad499c2ded128e16b9cef3863ddd7ab37c1c3413fc344aae25051"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5ae61a9a4cf00b8580db9bae5d9bdd41d1f5c8be"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
