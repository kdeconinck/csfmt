// =====================================================================================================================
// = LICENSE:       Copyright (c) 2021 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package token contains the lexical tokens of the C# programming language.
package token

// Kind represents a single lexical token.
type Kind uint8

// CONSTANTS: Defines ALL the lexical tokens.
const (
	// EndOfFile is the lexical token which represents the end of a file.
	EndOfFile Kind = iota

	// Unknown is the `default` lexical token which represents anything which isn't known as another token.
	Unknown
)

// PRIVATE: An array which contains ALL the lexical tokens.
//          Since the function Kind.String uses this, the order of the elements in this array MUST match the order in
//          which the lexical tokens are defined.
var typeMap = [...]string{
	"End of file",
	"Unknown",
}

// String returns the string representation of kind.
func (kind Kind) String() string {
	return typeMap[kind]
}
