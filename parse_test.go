// Copyright 2020 H2O.ai, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wave

import "testing"

type parseBytesTestCase struct {
	input  string
	output uint64
}

func TestParseBytes(t *testing.T) {
	_, ok, no := Assert(t)
	units := []string{"e", "p", "t", "g", "m", "k"}
	values := []int{EXABYTE, PETABYTE, TERABYTE, GIGABYTE, MEGABYTE, KILOBYTE}
	for i, unit := range units {
		v := values[i]

		b, err := parseBytes("5" + unit)
		no(err)
		ok(uint64(5*v) == b, "int without B")

		b, err = parseBytes("0" + unit)
		no(err)
		ok(0 == b, "zero")

		b, err = parseBytes("5" + unit + "b")
		no(err)
		ok(uint64(5*v) == b, "int with B")

		b, err = parseBytes("5" + unit + "ib")
		no(err)
		ok(uint64(5*v) == b, "int with iB")

		b, err = parseBytes("4.2" + unit + "ib")
		no(err)
		ok(uint64(4.2*float64(v)) == b, "float values")
	}

	b, err := parseBytes("  \n\t5MB\t")
	ok(uint64(5*MEGABYTE) == b, "allows whitespace")
	no(err)

	_, err = parseBytes("5")
	ok(err != nil, "no unit")

	_, err = parseBytes("-5bb")
	ok(err != nil, "bad unit")

	_, err = parseBytes("-5mbb")
	ok(err != nil, "bad unit")

	_, err = parseBytes("-5mb")
	ok(err != nil, "negative values")
}
