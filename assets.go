// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "fc2a12068064ad051ca47cb2fb4205402bd8cc241bbcab2a2471fd903f4c7ebd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0b69195b69b1b4dbf96d340eeffb562b711c4d41"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
