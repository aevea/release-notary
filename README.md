# Release notary

[![Build Status](https://cloud.drone.io/api/badges/commitsar-app/release-notary/status.svg)](https://cloud.drone.io/commitsar-app/release-notary)

Release notary is a tiny language agnostic app for automatically building changelogs. In the future it will also update GitLab and GitHub release notes. It is similar to https://goreleaser.com/ (which I highly recommend for Go projects) but doesn't handle builds and artifacts, just release notes.

Currently experimental.

### Table of contents

1. [Usage](#usage)

## Usage

Download and run: **(Substitute v0.0.2 for current version)**

```yml
- curl -L -O https://github.com/commitsar-app/release-notary/releases/download/v0.0.2/release-notary_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf release-notary_v0.0.2_Linux_x86_64.tar.gz
- ./release-notary log
```

Gitlab release documentation: https://gitlab.com/help/api/releases/index.md
