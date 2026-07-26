# Compozy Web Assets

This repository contains the generated production web bundle consumed by
`github.com/compozy/compozy`.

The Compozy main repository intentionally does not version generated frontend
assets. Release preparation builds `web/dist`, publishes this module, and then
the Compozy CLI imports this module so `go install github.com/compozy/compozy@latest`
ships the complete single binary with the embedded UI.

Do not edit `dist/` by hand. Regenerate it from the Compozy main repository.

Initial seed: copied from `github.com/compozy/agh-web-assets@v0.0.99` for the
v0.3 migration cutover.
