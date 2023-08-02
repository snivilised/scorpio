package command_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/i18n"
	"github.com/snivilised/arcadia/src/internal/helpers"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/snivilised/extendio/xfs/utils"

	ci18n "github.com/snivilised/cobrass/src/assistant/i18n"
)

var _ = Describe("WidgetCmd", Ordered, func() {
	var (
		repo     string
		l10nPath string
	)

	BeforeAll(func() {
		repo = helpers.Repo("../../..")
		l10nPath = helpers.Path(repo, "src/test/data/l10n")
		Expect(utils.FolderExists(l10nPath)).To(BeTrue(),
			fmt.Sprintf("💥 l10Path: '%v' does not exist", l10nPath),
		)
	})

	BeforeEach(func() {
		xi18n.ResetTx()
		err := xi18n.Use(func(uo *xi18n.UseOptions) {
			uo.From = xi18n.LoadFrom{
				Path: l10nPath,
				Sources: xi18n.TranslationFiles{
					i18n.ArcadiaSourceID: xi18n.TranslationSource{
						Name: "pixa",
					},

					ci18n.CobrassSourceID: xi18n.TranslationSource{
						Name: "cobrass",
					},
				},
			}
		})

		if err != nil {
			Fail(err.Error())
		}
	})

	When("specified flags are valid", func() {
		It("🧪 should: execute without error", func() {
			bootstrap := command.Bootstrap{
				Detector: &DetectorStub{},
			}
			tester := helpers.CommandTester{
				Args: []string{"widget", "-p", "P?<date>", "-t", "42"},
				Root: bootstrap.Root(),
			}
			_, err := tester.Execute()
			Expect(err).Error().To(BeNil(),
				"should pass validation due to all flag being valid",
			)
		})
	})

	When("specified flags are valid", func() {
		It("🧪 should: return error due to option validation failure", func() {
			bootstrap := command.Bootstrap{
				Detector: &DetectorStub{},
			}
			tester := helpers.CommandTester{
				Args: []string{"widget", "-p", "P?<date>", "-t", "99"},
				Root: bootstrap.Root(),
			}
			_, err := tester.Execute()
			Expect(err).Error().NotTo(BeNil(),
				"expected validation failure due to -t being within out of range",
			)
		})
	})
})
