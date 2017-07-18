### A minimal example of annotating test suites in Ginkgo

[Ginkgo][] suite files can be used to centralize bootstrapping for directories which
contain multiple test files. One (ab)use is to annotate the owner for the tests
contained within a single directory. Adding this information to the suite file has
the advantage of making test attribution to test authors who will automatically have
their tests annotated rather than having to remember to provide annotation at the
`Describe` container level.

### Running the tests

1. Clone this repository
1. Ensure you have [go][] installed
1. Install [direnv][]
1. Drop into `src/github.com/calebamiles/example-e2e-tests-by-sig/test`
1. Run `ginkgo -r` to see two simple examples of annotating your test suites

[Ginkgo]: https://github.com/onsi/ginkgo
[go]: https://golang.org/
[direnv]: https://direnv.net/
