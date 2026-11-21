// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "688260ddf8acca0128b472e5c5d93542fd0957b8292c3f9fd00e224a993ac270"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9edd902f6e761ecd4f8849e36b075d454a8c7f92"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
