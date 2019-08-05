package matrix

import (
	"fmt"
	"testing"
)

type testData struct {
	subject string
	body    string

	ignoreSubject bool
	ignoreBody    bool
	skipMarkdown  bool

	expectedOutput string
}

func TestGenerate(t *testing.T) {
	tests := []testData{
		testData{
			subject:        "test",
			body:           "content",
			ignoreSubject:  false,
			ignoreBody:     false,
			skipMarkdown:   false,
			expectedOutput: "# test\n\ncontent",
		},
		testData{
			subject:        "test",
			body:           "",
			ignoreSubject:  false,
			ignoreBody:     false,
			skipMarkdown:   false,
			expectedOutput: "# test",
		},
		testData{
			subject:        "test",
			body:           "content",
			ignoreSubject:  false,
			ignoreBody:     true,
			skipMarkdown:   false,
			expectedOutput: "# test",
		},
		testData{
			subject:        "test",
			body:           "content",
			ignoreSubject:  false,
			ignoreBody:     true,
			skipMarkdown:   true,
			expectedOutput: "test",
		},
		testData{
			subject:        "test",
			body:           "content",
			ignoreSubject:  true,
			ignoreBody:     false,
			skipMarkdown:   false,
			expectedOutput: "content",
		},
		testData{
			subject:        "test",
			body:           "content",
			ignoreSubject:  true,
			ignoreBody:     false,
			skipMarkdown:   true,
			expectedOutput: "content",
		},
		testData{
			subject:        "test",
			body:           "content",
			ignoreSubject:  false,
			ignoreBody:     false,
			skipMarkdown:   true,
			expectedOutput: "test\n\ncontent",
		},
	}

	for idx, test := range tests {
		test := test // capture range variable

		t.Run(fmt.Sprintf("test-%d", idx), func(t *testing.T) {
			t.Parallel()

			output := GenerateMessage(test.subject, test.body, test.ignoreSubject, test.ignoreBody, test.skipMarkdown)

			if output != test.expectedOutput {
				t.Errorf("Expected output `%s`, but got `%s`", test.expectedOutput, output)
				return
			}
		})
	}
}
