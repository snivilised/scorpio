package i18n

// TODO: Should be updated to use url of the implementing project,
// so should not be left as arcadia.
const ScorpioSourceID = "github.com/snivilised/scorpio"

type scorpioTemplData struct{}

func (td scorpioTemplData) SourceID() string {
	return ScorpioSourceID
}
