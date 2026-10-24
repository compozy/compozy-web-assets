// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9787c985885382f54ba579ebd6267ce31b8e9de9f6d85f114811fc87dd3ed24a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "420ba8b9bdbd1ecb21e36e60d893057afbfca269"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
