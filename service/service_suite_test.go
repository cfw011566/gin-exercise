package service

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVideoService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Video Service Test Suite")
}
