// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "879b6c6910234cc8062ba5187d903c022b9854680b258994ced89121615bc178"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "84297aa6f68e762bde2c24bbe43af40050ce3253"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
