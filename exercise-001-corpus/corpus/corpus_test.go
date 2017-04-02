package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert := assert.New(t)
	assert.Equal(len(result), 6)
	assert.Equal(result[0].Word, "you")
	assert.Equal(result[0].Count, 2)
}
