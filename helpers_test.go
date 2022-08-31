package rem

import "testing"

type cleanPathTest struct {
	inURL       string
	expectedURL string
}

func TestCleanPath(t *testing.T) {
	tests := []cleanPathTest{
		{ "test", "/test/" },
		{ "test/", "/test/" },
		{ "test//", "/test/" },
		{ "test///", "/test/" },
		{ "/test", "/test/" },
		{ "//test", "/test/" },
		{ "///test", "/test/" },
		{ "/test/", "/test/" },
		{ "//test//", "/test/" },
		{ "test/abc", "/test/abc/" },
		{ "//test/abc", "/test/abc/" },
		{ "/test/abc/", "/test/abc/" },
		{ "/test/abc//", "/test/abc/" },
		{ "///test/abc///", "/test/abc/" },
		{ "///test__/__abc///", "/test__/__abc/" },
	}

	for testId, test := range tests {
		outURL := cleanPath(test.inURL)
		if test.expectedURL != outURL {
			t.Errorf("TestId: %v\tIn URL: %v\tReceived URL: %v\tExpected URL: %v", testId, test.inURL, outURL, test.expectedURL)
		}
	}
}