package main_test

import (
	. "chi_contact_sample"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Context("initially", func() {
		It("has 0 contacts", func() {})
	})
	Context("when a new contact is added", func() {
		Context("the contacts", func() {
			It("has 1 more unique contact than it had earlier", func() {})
		})
	})
})
