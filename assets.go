// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "52fd5849e6963a51a0f6839f6c5fc07c5744f1664582294fb62170f42d8247f4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a30c2c408821f80a252ad1d55dffd4dd7a3c8bf3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
