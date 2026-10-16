// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dc8d9353c64e4f73267da9270ea836c718f10348ff27726da7d6d7c989ef9d87"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "73bed527cfad6eb9a0e21f003e9f1f2cc40d7873"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
