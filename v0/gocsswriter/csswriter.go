package gocsswriter

// GenerateCssFromString generate css from a html string
// given a html with utility classes generate the css
func GenerateCssFromString(Source string) string {
	return ""
}

// GenerateCssFromFile generate css from a html file
// given a html file with utility classes generate the css and save it to a file
func GenerateCssFromFile(sourcePath string, destPath string) error {
	return nil
}

// GenerateCssFromStringToDir generates css from a html string and saves it to a directory
// returns the path to the generated css file the file name is randomly generated
func GenerateCssFromStringToDir(source string, destDir string) (csspath string, err error) {
	panic("not implemented yet")
}

func GenerateCss() {
	
}

type CSSWriter struct{}
