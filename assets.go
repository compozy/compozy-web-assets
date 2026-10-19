// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b93d01ffc5b3c53e53fb4f6e1a7632a10074d9bee200d810a4ad5ba0e260299d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "878bd89f6994b820db735146b1a1f9da93598686"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
