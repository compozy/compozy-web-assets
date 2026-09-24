// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "84cd62f812f3cf404fadfae8d0db11d9b9a40c904a016bbb87057711eb5fd04d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "1f51928080d25616746376ab7272910edc7f6b31"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
