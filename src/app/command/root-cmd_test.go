package command_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive // ginkgo ok
	. "github.com/onsi/gomega"    //nolint:revive // gomega ok

	"github.com/snivilised/extendio/xfs/utils"
	"github.com/snivilised/scorpio/src/app/command"
	"github.com/snivilised/scorpio/src/internal/helpers"
)

var _ = Describe("RootCmd", Ordered, func() {
	var (
		repo     string
		l10nPath string
	)

	BeforeAll(func() {
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		Expect(utils.FolderExists(l10nPath)).To(BeTrue())
	})

	It("🧪 should: execute", func() {
		bootstrap := command.Bootstrap{
			Detector: &DetectorStub{},
		}
		tester := helpers.CommandTester{
			Args: []string{},
			Root: bootstrap.Root(),
		}
		_, err := tester.Execute()
		Expect(err).Error().To(BeNil())
	})
})
