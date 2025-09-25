package files

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"reflect"
	"time"
)

type Directory struct {
	parent *Directory
	path   PathParts
}

func (d *Directory) Path() string {
	return d.path.String()
}

// OpenDirectory constructs a Directory struct.
// To create a new directory in the file system, see CreateDirectory
func OpenDirectory(path string) (*Directory, error) {
	if !fs.ValidPath(path) {
		return nil, errors.New(path + " is not a valid path")
	}
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New(path + " was found, but is not a directory")
	}

	pathParts := PathPartsFromString(path)
	return &Directory{parent: nil, path: pathParts}, nil
}

// MakeDirectory creates a new directory in the file system with the path specified,
// and returns a matching Directory struct.
func MakeDirectory(path string) (*Directory, error) {
	if !fs.ValidPath(path) {
		return nil, errors.New(path + " is not a valid path")
	}
	err := os.MkdirAll(path, fs.ModeDir)
	if err != nil {
		return nil, err
	}
	return OpenDirectory(path)
}

// MakeDir creates a new directory in the file system with the path specified.
// All directories necessary will be created.
func MakeDir(path string) error {
	if !fs.ValidPath(path) {
		return errors.New(path + " is not a valid path")
	}
	return os.MkdirAll(path, fs.ModeDir)
}

func (d *Directory) Type() reflect.Type {
	return reflect.TypeOf(d)
}

func (d *Directory) Mode() fs.FileMode {
	if d.path == nil || len(d.path) == 0 {
		return fs.ModeIrregular
	}
	return fs.ModeDir
}

func (d *Directory) ModTime() time.Time {
	os.OpenRoot()
}

func (d *Directory) IsDir() bool {
	//TODO implement me
	panic("implement me")
}

func (d *Directory) Sys() any {
	//TODO implement me
	panic("implement me")
}

func (d *Directory) Glob(pattern string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Directory) Sub(dir string) (fs.FS, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Directory) Stat(name string) (fs.FileInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Directory) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Directory) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

// Parent returns this Directory's the parent Directory.
func (d *Directory) Parent() *Directory {
	if d.parent == nil {
		d.parent = new(Directory)
		d.parent.path = d.path[:len(d.path)-1]
	}
	return d.parent
}

// Root recursively retrieves this Directory's root Directory.
func (d *Directory) Root() *Directory {
	if d.parent == nil && (d.path == nil || len(d.path) == 0) {
		return d
	}
	return d.Parent().Root()
}

func (d *Directory) RootDir() string {
	if d.path != nil && len(d.path) > 1 {
		return d.path[0]
	}
	if d.parent == nil {
		return d.Name()
	}
	return d.parent.RootDir()
}

// Name returns the name of this Directory.
func (d *Directory) Name() string {
	return d.path[len(d.path)-1]
}

// List returns an array of each os.DirEntry in this Directory.
func (d *Directory) List() []os.DirEntry {
	dir, err := os.ReadDir(d.path.String())
	if err != nil {
		log.Println(err)
	}
	return dir
}

// ListDirEntries returns an array of each os.DirEntry in this Directory,
// sorted alphabetically.
// If an error occurs while reading the directory, ListDirEntries returns the entries
// it was able to read before the error, along with the error.
func (d *Directory) ListDirEntries() ([]os.DirEntry, error) {
	return os.ReadDir(d.path.String())
}

// ListEntries returns an array of each Entry in this Directory, generating
// them from the backing os.DirEntry instances.
func (d *Directory) ListEntries() []Entry {
	dir := d.List()
	entries := make([]Entry, 0, len(dir))
	for _, entry := range dir {
		entries = append(entries, EntryFromDirEntry(entry))
	}
	return entries
}

// Length returns the number of immediate Entries in this Directory (non-recursively).
// See Size for the size in bytes.
func (d *Directory) Length() int {
	return len(d.List())
}

// DescendantCount returns the number of Entries in this Directory,
// recursively traversing subdirectories.
// See Size for the size in bytes.
func (d *Directory) DescendantCount() int {
	return len(d.List())
}

// Size returns the size in bytes of everything contained in this Directory (recursively).
// See Length for the number of immediate entries.
func (d *Directory) Size() int64 {
	var size int64
	err := fs.WalkDir(d, ".", func(path string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil && info != nil {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		log.Println(err)
	}
	return size
}

func (d *Directory) SizeChildren() (os.FileInfo, error) {}

func (d *Directory) SubDirs() []os.DirEntry {
	var dirs []os.DirEntry
	list := d.List()
	for _, dir := range list {
		if dir.IsDir() {
			dirs = append(dirs, dir)
		}
	}
	return dirs
}

func (d *Directory) SubDirectories() []Directory {
	var dirs []Directory
	list := d.List()
	subPath := append(d.path, d.Name())
	for _, dir := range list {
		if dir.IsDir() {
			dirs = append(dirs, Directory{d, subPath})
		}
	}
	return dirs
}
