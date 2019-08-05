package mime

import (
	"os"
	"testing"
)

type TestData struct {
	filePath        string
	expectedSubject string
	expectedContent string
}

func TestParser(t *testing.T) {
	tests := []TestData{
		TestData{
			filePath:        "./testdata/phpmailer-quoted-printable.txt",
			expectedSubject: "Reminder from example.com",
			expectedContent: "Тест",
		},
		TestData{
			filePath:        "./testdata/geary-plain.txt",
			expectedSubject: "some test",
			expectedContent: "test",
		},
		TestData{
			filePath:        "./testdata/gmail-base64.txt",
			expectedSubject: "тест от gmail",
			expectedContent: "тест от gmail",
		},
	}

	for _, test := range tests {
		test := test // capture range variable

		t.Run(test.filePath, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(test.filePath)
			if err != nil {
				t.Errorf("Failed to open file: %s: %s", test.filePath, err)
				return
			}
			defer f.Close()

			subject, content, err := ExtractContentFromEmail(f)
			if err != nil {
				t.Error(err)
				return
			}

			if subject != test.expectedSubject {
				t.Errorf("Expected subject `%s`, but got `%s`", test.expectedSubject, subject)
				return
			}

			if content != test.expectedContent {
				t.Errorf("Expected content `%s`, but got `%s`", test.expectedContent, content)
				return
			}
		})
	}
}
