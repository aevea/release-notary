# Contributing

By participating to this project, you agree to abide our [code of conduct](/CODE_OF_CONDUCT.md).

There are also some shared packages in the namespace:

- For Git logic https://github.com/aevea/git
- For Commit logic https://github.com/aevea/quoad

# Pre-Requisites

- `make`
- [Go 1.13+](https://golang.org/doc/install)
- [Git](https://git-scm.com/)

# Tests

PRs should include tests if they are fixing a bug or creating a new feature. This way we can be optimistic that master can be used on other projects.

Tests should be run before submitting a PR by:

- sh \$(cd testdata && ./setup_test_repos.sh)
- make ci

Worst case the CI will catch any errors. It's faster to run it locally however.

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we sare using Conventional Commits. These are enforced by the CI.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Using /testdata

The testdata folder includes a bunch of [Git bundles](https://git-scm.com/docs/git-bundle).

These allow for `git clone` like a normal repository and in the project they include actual testcases. It's a great approach for creating easy-to-test and reproduce repositories that don't need to be hosted anywhere.

**To create a new bundle**

- make a new folder
- `git init` the new folder
- add some commits and optionally tags (depends what you are testing)
- `git bundle create nameoffolder.bundle --all`
- copy this bundle to the root testdata
- add the clone command to ./testdata/setup_test_repos.sh
