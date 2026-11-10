// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0da5897527eca5ef01b18d1402e6a42dc6bf9e809b504b623f17c4fc51aacd4e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "52038d6521c869026b6bc0f80b401b67798cf46d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
