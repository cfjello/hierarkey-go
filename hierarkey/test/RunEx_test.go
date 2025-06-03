package testing

import (
	"testing"

	"github.com/cfjello/hierarkey-go/hierarkey/pkg/examples"
)

func RunExamples(t *testing.T) {
	t.Run("Example A", func(t *testing.T) {
		examples.RunExampleA()
	})
	t.Run("Example B", func(t *testing.T) {
		examples.RunExampleB()
	})
}
