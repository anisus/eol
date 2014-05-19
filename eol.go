// Package eol provides string constants for line terminators.
//
//
// Usage
//
// The package provides the convenience OS constant:
//   fmt.Sprint(w, "Text" + eol.OS )
// Eg. when targeting Windows you will have:
//   const OS = CRLF
package eol

// Unicode line terminators as defined in:
// http://unicode.org/standard/reports/tr13/tr13-5.html
const (
	LF   = "\n"     // Line Feed
	CRLF = "\r\n"   // Carriage Return + Line Feed
	VT   = "\v"     // Vertical Tab
	FF   = "\f"     // Form Feed
	CR   = "\r\n"   // Carriage Return
	NEL  = "\u0085" // Next Line
	LS   = "\u2028" // Line Separator
	PS   = "\u2029" // Paragraph Separator
)

// OS specific line terminators
const (
	Darwin  = LF
	FreeBSD = LF
	NetBSD  = LF
	OpenBSD = LF
	Plan9   = LF
	Windows = CRLF
)
