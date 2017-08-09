package sig_node_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"

	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

// this would live somewhere in framework/ for consistency
type SIGOwner struct {
	Name string `json:"name"`
	GitHubHandle string `json:"GitHub"`
	SlackID string `json:"Slack"`
}


type SIGTag struct {
	SIGName string `json:"SIG"`
	Owners []*SIGOwner `json:"Owners"`
}

type FailureReport struct {
	Owner SIGTag
	Failures []string `json:"test_failures"`
}

type SIGReporter struct {
	owner SIGTag
	failures []*types.SpecSummary
}


// noop
func (r *SIGReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {}


// noop
func (r *SIGReporter) BeforeSuiteDidRun(_  *types.SetupSummary) {}

// noop
func (r *SIGReporter) SpecWillRun(_  *types.SpecSummary) {}

func (r *SIGReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	switch specSummary.State {
	case types.SpecStateFailed:
		r.failures = append(r.failures, specSummary)
	case types.SpecStateTimedOut:
		r.failures = append(r.failures, specSummary)
	case types.SpecStatePanicked:
		r.failures = append(r.failures, specSummary)
	}
}



// noop
func (r *SIGReporter) AfterSuiteDidRun(_  *types.SetupSummary) {}

// process any failures from the suite
func (r *SIGReporter) SpecSuiteDidEnd(summary  *types.SuiteSummary) {

	// early exit if there's nothing to do
	if summary.SuiteSucceeded {
		return
	}

	report := &FailureReport{Owner: r.owner}

	for i := range r.failures {
		report.Failures = append(report.Failures, fmt.Sprintf("%s:%d", r.failures[i].Failure.Location.FileName, r.failures[i].Failure.Location.LineNumber))
	}


	reportBytes, _ := json.Marshal(report)
	prettyJSON := &bytes.Buffer{}

	json.Indent(prettyJSON, reportBytes, "", "\t")

	fmt.Println(prettyJSON.String())
}









func TestSigNode(t *testing.T) {
	sig := SIGTag{}
	sig.SIGName = "SIG Node"

	dawn := &SIGOwner{Name: "Dawn Chen", GitHubHandle: "someGitHubHandle", SlackID: "someSlackID"}
	derek := &SIGOwner{Name: "Derek Carr", GitHubHandle: "someGitHubHandle", SlackID: "someSlackID"}

	sig.Owners = append(sig.Owners, dawn)
	sig.Owners = append(sig.Owners, derek)


	sigBytes, _ := json.Marshal(sig)

	myReporter := &SIGReporter{owner: sig}
	RegisterFailHandler(Fail)

	reporters := []Reporter{myReporter}
	RunSpecsWithDefaultAndCustomReporters(t, string(sigBytes), reporters)
	//RunSpecs(t, string(sigBytes))
}
