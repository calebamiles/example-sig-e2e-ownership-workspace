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
1. Run `ginkgo -r --keepGoing` to see two simple examples of annotating your test suites

```
[1502236921] SIG Network - 1/1 specs 
------------------------------
• Failure [0.000 seconds]
Example [It] also fails 
/home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-network/example_test.go:10

  this test is flakey

  /home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-network/example_test.go:9
------------------------------


Summarizing 1 Failure:

[Fail] Example [It] also fails 
/home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-network/example_test.go:9

Ran 1 of 1 Specs in 0.000 seconds
FAIL! -- 0 Passed | 1 Failed | 0 Pending | 0 Skipped --- FAIL: TestSigNetwork (0.00s)
FAIL
[1502236921] {"SIG":"SIG Node","Owners":[{"name":"Dawn Chen","GitHub":"someGitHubHandle","Slack":"someSlackID"},{"name":"Derek Carr","GitHub":"someGitHubHandle","Slack":"someSlackID"}]} - 1/1 specs 
------------------------------
• Failure [0.000 seconds]
Example [It] fails 
/home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-node/example_test.go:11

  we didn't get around to writing this test

  /home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-node/example_test.go:10
------------------------------


Summarizing 1 Failure:

[Fail] Example [It] fails 
/home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-node/example_test.go:10

Ran 1 of 1 Specs in 0.000 seconds
FAIL! -- 0 Passed | 1 Failed | 0 Pending | 0 Skipped {
	"Owner": {
		"SIG": "SIG Node",
		"Owners": [
			{
				"name": "Dawn Chen",
				"GitHub": "someGitHubHandle",
				"Slack": "someSlackID"
			},
			{
				"name": "Derek Carr",
				"GitHub": "someGitHubHandle",
				"Slack": "someSlackID"
			}
		]
	},
	"test_failures": [
		"/home/caleb/workspace/example-e2e-by-sig/src/github.com/calebamiles/example-e2e-tests-by-sig/test/sig-node/example_test.go:10"
	]
}
--- FAIL: TestSigNode (0.00s)
FAIL

There were failures detected in the following suites:
	sig-network ./sig-network
	   sig-node ./sig-node

Ginkgo ran 2 suites in 572.163765ms
Test Suite Failed
```

[Ginkgo]: https://github.com/onsi/ginkgo
[go]: https://golang.org/
[direnv]: https://direnv.net/
