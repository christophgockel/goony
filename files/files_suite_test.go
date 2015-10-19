package files_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFiles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Files Suite")
}
