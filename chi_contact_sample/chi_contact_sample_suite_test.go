package chi_contact_sample_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChiContactSample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChiContactSample Suite")
}
