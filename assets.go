// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9787c985885382f54ba579ebd6267ce31b8e9de9f6d85f114811fc87dd3ed24a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f63227d8473496b942cb83a42653315f2cef4ddf"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
