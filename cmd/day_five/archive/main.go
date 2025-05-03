package main

import (
	"fmt"
	"sync"
)

/*
### Anthropic Software Engineering Interview Prompts

These are sample prompts based on reported candidate experiences during live 55-minute software engineering interviews at Anthropic. Focus is on system design and practical implementation, not algorithm puzzles.

---

### 1. Design a File Store System

**Prompt**:

Design and implement an in-memory file storage system with the following functionalities:

- `set(key, value)`: Store a key-value pair.
- `get(key)`: Retrieve the value associated with a key.
- `filter(predicate)`: Return all key-value pairs that satisfy a given condition.
- `backup()`: Create a backup of the current state.
- `restore()`: Restore the system to the last backup state.

**Considerations**:

- Data persistence and recovery
- Efficient data retrieval and storage
- Scalability for large datasets
*/


type File struct {
	contents string // this is the raw-string
	metadata map[string]string // this is metadata about the file update times, user, creation time, etc etc
}

type FileSystem struct {
	l sync.Mutex
	FileMap map[string]*File
	BackUpFileMap map[string]*File
}

func NewFileSystem() *FileSystem {
	filesystemMap := make(map[string]*File)
	filesystem := FileSystem {
		l: sync.Mutex{}, // this is explicit but it's also automatic
		FileMap: filesystemMap,
	}
	return &filesystem
}


func (f *FileSystem) set(key string, value *File) {
	f.l.Lock()
	defer f.l.Unlock()
	f.FileMap[key] = value
}

func (f *FileSystem) get(key string) *File {
	f.l.Lock()
	defer f.l.Unlock()

  val, ok := f.FileMap[key]
	if !ok {
		fmt.Println("key not found")
		return nil
	}

	return val
}

func (f *FileSystem) filter(predicate func(string, *File) bool) map[string]*File{
	f.l.Lock()
	defer f.l.Unlock()

	var resultMap = make(map[string]*File)

	for k, v := range f.FileMap {
		if predicate(k, v) {
			resultMap[k] = v
		}
	}

	return resultMap
}

func (f *FileSystem) backup() {
	f.l.Lock()
	defer f.l.Unlock()

	var tmpFileSystem = make(map[string]*File)

	for k, v := range f.FileMap {
		var metametadata = make(map[string]string)
		for km, vm := range v.metadata {
			metametadata[km] = vm
		}

		tmpFileSystem[k] = &File{
			contents: v.contents,
			metadata: metametadata,
		}
	}

	f.BackUpFileMap = tmpFileSystem
}

func (f *FileSystem) restore(){
	f.l.Lock()
	defer f.l.Unlock()

	var tmpFileSystem = make(map[string]*File)

	for k, v := range f.BackUpFileMap {
		var tmpMetadata = make(map[string]string)
		for km, vm := range v.metadata {
			tmpMetadata[km] = vm
		}
		tmpFileSystem[k] = &File {
			contents: v.contents,
			metadata: tmpMetadata,

		}
		f.FileMap = tmpFileSystem
	}

}





func main() {}
