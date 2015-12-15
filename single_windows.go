// +build windows

package single

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Filename returns an absolute filename, appropriate for the operating system
func (s *Single) Filename() string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s.lock", s.name))
}

// Lock tries to remove the lock file, if it exists.
// If the file is already open by another instance of the program,
// remove will fail and exit the program.
func (s *Single) Lock() {

	if err := os.Remove(s.Filename()); err != nil && !os.IsNotExist(err) {
		log.Fatal(ErrAlreadyRunning)
	}

	file, err := os.OpenFile(s.Filename(), os.O_EXCL|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	s.file = file
}

// Unlock closes and removes the lockfile.
func (s *Single) Unlock() {
	if err := s.file.Close(); err != nil {
		log.Print(err)
	}
	if err := os.Remove(s.Filename()); err != nil {
		log.Print(err)
	}
}
