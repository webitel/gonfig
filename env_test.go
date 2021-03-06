package gonfig

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"strings"
)

var _ = Describe("EnvConfig", func() {
	var (
		err error
		cfg ReadableConfig
	)
	BeforeEach(func() {
		cfg = NewEnvConfig("")
		err = cfg.Load()
		Expect(err == nil)
	})
	It("Should load variables from environment", func() {
		Expect(len(cfg.All()) > 0).To(BeTrue())
		env := os.Environ()
		Expect(len(env) > 0).To(BeTrue())
		for _, kvpair := range env {
			pairs := strings.SplitN(kvpair, "=", 2)
			Expect(len(pairs) == 2).To(BeTrue())
			Expect(cfg.Get(pairs[0])).To(Equal(pairs[1]))
		}
	})
})
