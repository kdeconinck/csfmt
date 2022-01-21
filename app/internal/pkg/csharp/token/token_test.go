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

// Quality assurance: Verify (and measure the performance) of the public API of the token package.
package token_test

import (
	"testing"

	"github.com/kdeconinck/csfmt/internal/pkg/csharp/token"
)

// PRIVATE: The scenario's to verify (and measure the performance) of the Kind.String function.
var tbKindScenarios = [...]struct {
	description    string
	tokenKind      token.Kind
	expectedString string
}{
	{
		description:    "Get the string representation of the `EOF` token.",
		tokenKind:      token.EndOfFile,
		expectedString: "End of file",
	},

	{
		description:    "Get the string representation of the `UNKNOWN` token.",
		tokenKind:      token.Unknown,
		expectedString: "Unknown",
	},
}

// UT: Verify that the Kind.String function is implemented correctly.
func TestKindString(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// EXECUTION.
	for _, scenario := range tbKindScenarios {
		scenario := scenario // NOTE: Ensure that the t.Run function has the correct value when it's being executed.

		t.Run(scenario.description, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// ACT.
			result := scenario.tokenKind.String()

			// ASSERT.
			if result != scenario.expectedString {
				t.Fatalf("\n\n"+
					"UT Name:  %v\n"+
					"\033[32mExpected: %v\033[0m\n"+
					"\033[31mActual:   %v\033[0m\n\n",
					scenario.description, scenario.expectedString, result)
			}
		})
	}
}

// PRIVATE: Holds the result of the Kind.String function, calculated in the Kind.String benchmark.
//          This variable is required to avoid compiler optimizations.
var benchmarkKindStringOutput string

// BENCHMARK: Measure the performance of the Kind.String function.
func BenchmarkKindString(b *testing.B) {
	// EXECUTION.
	for _, scenario := range tbKindScenarios {
		scenario := scenario // NOTE: Ensure that the b.Run function has the correct value when it's being executed.

		b.Run(scenario.description, func(b *testing.B) {
			// RUN.
			for n := 0; n < b.N; n++ {
				// ACT.
				benchmarkKindStringOutput = scenario.tokenKind.String()
			}
		})
	}
}
