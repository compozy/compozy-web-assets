// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "53afe2fb1bf185e9c0eae99559c2f85af5c08d25822b11728f84072cf77cca6e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ed2561ce8464cb79bc416cfe532e65f7440fb0a7"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
