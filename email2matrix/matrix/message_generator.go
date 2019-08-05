package matrix

import "fmt"

func GenerateMessage(
	subject string,
	body string,
	ignoreSubject, ignoreBody, skipMarkdown bool,
) string {
	if skipMarkdown {
		if ignoreBody || body == "" {
			if subject == "" {
				return ""
			}
			return fmt.Sprintf("%s", subject)
		}

		if ignoreSubject || subject == "" {
			return body
		}

		return fmt.Sprintf("%s\n\n%s", subject, body)
	}

	if ignoreBody || body == "" {
		if subject == "" {
			return ""
		}
		return fmt.Sprintf("# %s", subject)
	}

	if ignoreSubject || subject == "" {
		return body
	}

	return fmt.Sprintf("# %s\n\n%s", subject, body)
}
