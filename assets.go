// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ef17a65be39c3b5962c48f3ecba76a3e4178914ff7881454d45484ba99b78665"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fec0e9b08631dd9e2d7eaa66fc7f76f5314ae5c0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
