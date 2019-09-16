# Release notary

[![Build Status](https://cloud.drone.io/api/badges/commitsar-app/release-notary/status.svg)](https://cloud.drone.io/commitsar-app/release-notary)

Release Notary builds release notes using [Conventional Commit](https://www.conventionalcommits.org/) standard and then publishes it to Github. Release notes are appended to any text you already have in your release and therefore will not affect important announcements etc. Currently only Github is supported, but Gitlab is on the way as well.

Heavily inspired by https://github.com/graphql/graphql-js/releases but usable as a standalone app.

Expected output is [HERE](./expected-output.md)

:warning: **Currently experimental.** :warning:

### Table of contents

1. [Usage](#usage)

## Usage

#### Required variables

- GITHUB_TOKEN
- GITHUB_REPOSITORY (in the format of `owner/repository` e.g. `commitsar-app/release-notary`)

#### Using Github actions

Checkout git in order to get commits and master branch

```
- name: Check out code into the Go module directory
  uses: actions/checkout@v1
```

Run the Release Notary action. Github token needs to be explicitly added so that Release Notary can use it. See https://help.github.com/en/articles/virtual-environments-for-github-actions

```
- name: Release Notary Action
  uses: commitsar-app/release-notary@v0.2.0 (substitute for current version)
  env:
      GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

#### Using binary

Download and run: **(Substitute v0.0.2 for current version)**

```yml
- curl -L -O https://github.com/commitsar-app/release-notary/releases/download/v0.0.2/release-notary_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf release-notary_v0.0.2_Linux_x86_64.tar.gz
- export GITHUB_TOKEN=yourtoken
- export GITHUB_REPOSITORY=owner/repo
- ./release-notary publish
```

Gitlab release documentation: https://gitlab.com/help/api/releases/index.md
