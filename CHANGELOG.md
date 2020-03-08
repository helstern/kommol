# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [v0.2.0] - 2020-03-08
### Added
- read the bucket's website configuration and serve the index file from a GCP bucket when the root file (`/`) is requested 
 - new cli option `log-level` which controls the level of logging. defaults to info
 - improve logging


## [v0.1.5] - 2019-10-17
### Changed
 - fix e2e tests
 - ignore the port from the hostname when using the static website bucket strategy
 - use read-only scope with GCP authorization   
    
## [v0.1.4] - 2019-10-14
### Added
 - end-to-end tests for haproxy using BATS
### Changed
 - adopt a more conventional golang project structure
 - fix handle relative HTTP requests
 
## [v0.1.3] - 2019-10-12
### Changed
- fix use actual values instead of env variables in travis deploy script section

## [v0.1.2] - 2019-10-12
### Added
- fix enable travis builds on release and deploy branches

## [v0.1.1] - 2019-10-12
### Changed
- fix travis build script

## [v0.1.0] - 2019-10-12
### Added
- very basic Google Cloud Storage reverse proxy with limited logging

[Unreleased]: https://github.com/helstern/kommol/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/helstern/kommol/compare/v0.1.5...v0.2.0
[v0.1.5]: https://github.com/helstern/kommol/compare/v0.1.4...v0.1.5
[v0.1.4]: https://github.com/helstern/kommol/compare/v0.1.3...v0.1.4
[v0.1.3]: https://github.com/helstern/kommol/compare/v0.1.2...v0.1.3
[v0.1.2]: https://github.com/helstern/kommol/compare/v0.1.1...v0.1.2
[v0.1.1]: https://github.com/helstern/kommol/compare/v0.1.0...v0.1.1
[v0.1.0]: https://github.com/helstern/kommol/compare/cbcc6ff...v0.1.0
