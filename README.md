# Release notary

[![Build Status](https://cloud.drone.io/api/badges/commitsar-app/release-notary/status.svg)](https://cloud.drone.io/commitsar-app/release-notary)

Release Notary builds release notes using [Conventional Commit](https://www.conventionalcommits.org/) standard and then publishes it to Github. Release notes are appended to any text you already have in your release and therefore will not affect important announcements etc.

Currently supported providers are: `[Github, Gitlab]`.

Heavily inspired by https://github.com/graphql/graphql-js/releases, but usable as a standalone app.

Expected output is [HERE](./expected-output.md)

:warning: **Currently experimental. Please report any issues** :warning:

### Table of contents

1. [Setup](#setup)
   - [Github](#github)
   - [Gitlab](#gitlab)
2. [Usage](#usage)

## Setup

### Github

**Variables:**

| Name              | Example value                                                                                                                          | Required |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| GITHUB_TOKEN      | token (provided in Github Action or [HERE](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)) | true     |
| GITHUB_REPOSITORY | commitsar-app/commitsar                                                                                                                | true     |

In Github actions GITHUB_REPOSITORY is provided in the correct format. Does not need to be explicitly added.

In Github actions GITHUB_TOKEN is provided with the correct access rights, elsewhere it needs to be generated and added to the pipeline. Please see the [usage](#usage) section.

### Gitlab

**Variables:**

In Gitlab CI all the values are provided except for `GITLAB_TOKEN`. Documentation provided [here](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html).

| Name          | Example value             | Required |
| ------------- | ------------------------- | -------- |
| GITLAB_TOKEN  | token                     | true     |
| CI_API_V4_URL | https://gitlab.com/api/v4 | true     |
| CI_COMMIT_TAG | v0.0.4                    | true     |
| CI_PROJECT_ID | 1234                      | true     |

## Usage

#### Using Github actions

Should be run only on tags, example is [HERE](https://github.com/commitsar-app/commitsar/blob/master/.github/workflows/release.yml):

```yml
on:
  push:
    tags:
      - "v*"
```

Checkout git in order to get commits and master branch

```yml
- name: Check out code into the Go module directory
  uses: actions/checkout@v1
```

Run the Release Notary action. Github token needs to be explicitly added so that Release Notary can use it. See https://help.github.com/en/articles/virtual-environments-for-github-actions

```yml
- name: Release Notary Action
  uses: commitsar-app/release-notary@v0.2.0 (substitute for current version)
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

#### Gitlab CI

```yml
release:
  stage: release
  variables:
    GITLAB_TOKEN: $RELEASE_TOKEN
  image: commitsar/release-notary
  script:
    - release-notary publish
  only:
    - tags
```

#### Using binary

Download and run: **(Substitute v0.0.2 for current version)**

```yml
- curl -L -O https://github.com/commitsar-app/release-notary/releases/download/v0.0.2/release-notary_v0.0.2_Linux_x86_64.tar.gz
- tar -xzf release-notary_v0.0.2_Linux_x86_64.tar.gz
# Set up any required variables
- export GITHUB_TOKEN=yourtoken
- export GITHUB_REPOSITORY=owner/repo
- ./release-notary publish
```
