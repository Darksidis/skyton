//go:build windows
// +build windows

package file

import "testing"

var uncTestPaths = []string{
	`C:\Ba*d\P|a?t<h>\Windows\Folder`,
	`C:\Windows\Folder`,
	`\\?\C:\Windows\Folder`,
	`\\?\UNC\server\share\Desktop`,
	`\\?\unC\server\share\Desktop\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`\\server\share\Desktop\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`C:\Desktop\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`C:\AbsoluteToRoot\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`\\server\share\Desktop`,
	`\\?\UNC\\share\folder\Desktop`,
	`\\server\share`,
}

var uncTestPathsResults = []string{
	`\\?\C:\Ba*d\P|a?t<h>\Windows\Folder`,
	`\\?\C:\Windows\Folder`,
	`\\?\C:\Windows\Folder`,
	`\\?\UNC\server\share\Desktop`,
	`\\?\unC\server\share\Desktop\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`\\?\UNC\server\share\Desktop\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`\\?\C:\Desktop\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`\\?\C:\AbsoluteToRoot\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path\Very Long path`,
	`\\?\UNC\server\share\Desktop`,
	`\\?\UNC\\share\folder\Desktop`,
	`\\?\UNC\server\share`,
}

// Test that UNC paths are converted.
func TestUncPaths(t *testing.T) {
	for i, p := range uncTestPaths {
		unc := UNCPath(p)
		if unc != uncTestPathsResults[i] {
			t.Fatalf("UNC test path\nInput:%s\nOutput:%s\nExpected:%s", p, unc, uncTestPathsResults[i])
		}
		// Test we don't add more.
		unc = UNCPath(unc)
		if unc != uncTestPathsResults[i] {
			t.Fatalf("UNC test path\nInput:%s\nOutput:%s\nExpected:%s", p, unc, uncTestPathsResults[i])
		}
	}
}
