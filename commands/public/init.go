package public

import "regexp"

var (
	regularizer *regexp.Regexp
	Logfile     string
)

func init() {
	regularizer = regexp.MustCompile(`(?i)<script.*(</script[^>]*>)?`)
}
