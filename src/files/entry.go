package files

import (
	"io/fs"
	"log"
	"os"
	"time"
)

type Entry struct {
	dirent os.DirEntry
	info   fs.FileInfo
}

func (e *Entry) Stat() (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (e *Entry) Read(bytes []byte) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (e *Entry) Close() error {
	//TODO implement me
	panic("implement me")
}

func (e *Entry) Mode() fs.FileMode {
	return e.info.Mode()
}

func (e *Entry) ModTime() time.Time {
	//TODO implement me
	panic("implement me")
}

func (e *Entry) Sys() any {
	//TODO implement me
	panic("implement me")
}

func (e *Entry) String() string {
	return e.dirent.Name()
}

func EntryFromDirEntry(entry os.DirEntry) Entry {
	info, err := entry.Info()
	if err != nil {
		panic(err)
	}
	return Entry{dirent: entry, info: info}
}

// Size returns the size (in bytes) of this Entry.
// Note that symbolic links will return the size of the link, not its
// target.
func (e *Entry) Size() int64 {
	info, err := e.dirent.Info()
	if err != nil {
		log.Println(err)
	}
	return info.Size()
}

// ========================
// DirEntry implementation
// ========================

// Name returns the name of the file (or subdirectory) described by the entry.
// This name is only the final element of the path (the base name), not the entire path.
// For example, Name would return "hello.go" not "home/gopher/hello.go".
func (e *Entry) Name() string {
	return e.dirent.Name()
}

// IsDir reports whether the entry describes a directory.
func (e *Entry) IsDir() bool {
	return e.dirent.IsDir()
}

// Type returns the type bits for the entry.
// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
func (e *Entry) Type() fs.FileMode {
	return e.dirent.Type()
}

// Info returns the FileInfo for the file or subdirectory described by the entry.
// The returned FileInfo may be from the time of the original directory read
// or from the time of the call to Info. If the file has been removed or renamed
// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
// If the entry denotes a symbolic link, Info reports the information about the link itself,
// not the link's target.
func (e *Entry) Info() (fs.FileInfo, error) {
	return e.dirent.Info()
}

func FileInfoToEntry(info fs.FileInfo) Entry {
	entry := Entry{}
	if info == nil {
		return entry
	}
	entry.dirent = fs.FileInfoToDirEntry(info)
	entry.info = info
	return entry
}

type Entries = []Entry
type DirEntries = []os.DirEntry
