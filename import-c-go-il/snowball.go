// snowball stemmer
//
// Example:
//  stemmer = snowball.New("english")
//  fmt.Println(stemmer.stem("running")) // Will print "run"
package snowball

// START_IMPORT OMIT
import (
	"fmt"
	"runtime"
	"unsafe"
)
/*
#include <stdlib.h>
#include "libstemmer.h"
*/
import "C"


// Stemmer structure
type Stemmer struct {
	lang string
	stmr *C.struct_sb_stemmer
}

// END_IMPORT OMIT

const (
	Version = "0.1.2"
)

// START_NEW OMIT
// New creates a new stemmer for lang
func New(lang string) (*Stemmer, error) {
	stmr := &Stemmer{
		lang,
		C.sb_stemmer_new(C.CString(lang), nil),
	}

	if stmr.stmr == nil {
		return nil, fmt.Errorf("can't create stemmer for lang %s", lang)
	}

	runtime.SetFinalizer(stmr, free)  // Free C memory when GCed

	return stmr, nil
}

// free C resources
func free(stmr *Stemmer) {
	if stmr.stmr != nil {
		C.sb_stemmer_delete(stmr.stmr)
		stmr.stmr = nil
	}
}
// END_NEW OMIT

// Lang return the stemmer language
func (stmr *Stemmer) Lang() string {
	return stmr.lang
}

// START_STEM OMIT
// Stem returns them stem of word (e.g. running -> run)
func (stmr *Stemmer) Stem(word string) string {
	ptr := unsafe.Pointer(C.CString(word))
	defer C.free(ptr)

	w := (*C.sb_symbol)(ptr)
	res := unsafe.Pointer(C.sb_stemmer_stem(stmr.stmr, w, C.int(len(word))))
	size := C.sb_stemmer_length(stmr.stmr)

	buf := C.GoBytes(res, size)
	return string(buf)
}
// END_STEM OMIT

// START_LIST OMIT
// List returns the list of languages supported by snowball
func List() []string {
	names := []string{}

	// We don't need to free since sb_stemmer_list return pointer to static variable
	cp := uintptr(unsafe.Pointer(C.sb_stemmer_list()))
	size := unsafe.Sizeof(uintptr(0))

	for {
		name := C.GoString(*(**C.char)(unsafe.Pointer(cp)))
		if len(name) == 0 {
			break
		}
		names = append(names, name)
		cp += size
	}
	return names
}
// END_LIST OMIT
