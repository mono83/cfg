package file

import (
	"os"
	"runtime"
)

// CommonFolders returns list of common folder for configuration location
// including current folder, home folder and etc
func CommonFolders(filename string) []string {
	return []string{
		filename,
		getHomeFolder() + "/" + filename,
		"/etc/" + filename,
	}
}

// CommonFoldersWithSubfolder returns list of common folder for configuration location
// including current folder, home folder and etc
func CommonFoldersWithSubfolder(filename, subfolder string) []string {
	return []string{
		filename,
		subfolder + "/" + filename,
		getHomeFolder() + "/" + subfolder + "/" + filename,
		"/etc/" + subfolder + "/" + filename,
	}
}

func getHomeFolder() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
