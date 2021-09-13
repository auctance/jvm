package classpath
import "os"
import "path/filepath"
import "strings"
func newWildcardEntry(path string) CompositeEntry{
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error{}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}