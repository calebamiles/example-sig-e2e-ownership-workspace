package sig_node_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	"testing"
)

// this would live somewhere in framework/ for consistency
type SIGTag struct {
	SIGName string `json:"SIGName"`
	SIGOwners []string `json:"SIGOwners"`
}


func TestSigNode(t *testing.T) {
	sig := SIGTag{}
	sig.SIGName = "Kubernetes SIG Node"
	sig.SIGOwners = append(sig.SIGOwners, "Dawn Chen")
	sig.SIGOwners = append(sig.SIGOwners, "Derek Carr")


	sigBytes, _ := json.Marshal(sig)

	RegisterFailHandler(Fail)
	RunSpecs(t, string(sigBytes))
}
