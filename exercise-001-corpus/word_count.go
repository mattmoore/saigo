package main

import (
	"fmt"
	"github.com/enova/saigo/exercise-001-corpus/corpus"
)

func main() {
	analysis := corpus.Analyze("Are you serious? I knew you were.")

	for _, word := range analysis {
		fmt.Println(word.Count, word.Word)
	}
}
