// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "0eeb86338f2db06eaae43ab2769d1af36d204989d553cc6263bd40ef8e2278ab"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bfd56223bb2982fa756916f79031e079fd09fd2b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
