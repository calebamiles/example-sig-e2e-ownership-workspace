package sig_network_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSigNetwork(t *testing.T) {
	RegisterFailHandler(Fail)

	// even without fancy structs annotating SIG ownership is possible today
	RunSpecs(t, "SIG Network")
}
