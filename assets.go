// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "db50298aed6d1a78024a2e66ba79b89e36802c089262edc8fc3d2c22336ae605"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2549c6e0f0f844d5e62b2d40929f7401ede02e95"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
