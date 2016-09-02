package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCommonFolders(t *testing.T) {
	a := assert.New(t)

	list := CommonFolders("foo.json")
	a.Len(list, 4)
	a.Equal("foo.json", list[0])
	a.Equal(os.Getenv("HOME")+"/foo.json", list[1])
	a.Equal("/etc/foo.json", list[2])
	a.Equal("/usr/local/etc/foo.json", list[3])
}

func TestCommonFolderWithSubfolder(t *testing.T) {
	a := assert.New(t)

	list := CommonFoldersWithSubfolder("bar.yml", "app")
	a.Len(list, 5)
	a.Equal("bar.yml", list[0])
	a.Equal("app/bar.yml", list[1])
	a.Equal(os.Getenv("HOME")+"/app/bar.yml", list[2])
	a.Equal("/etc/app/bar.yml", list[3])
	a.Equal("/usr/local/etc/app/bar.yml", list[4])
}
