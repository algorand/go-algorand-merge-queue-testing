// Copyright (C) 2019-2023 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package logic

import (
	"strings"
	"testing"

	"github.com/algorand/go-algorand/test/partitiontest"
	"github.com/stretchr/testify/require"
)

func TestGetSourceMap(t *testing.T) {
	partitiontest.PartitionTest(t)
	t.Parallel()
	a := require.New(t)

	sourceNames := []string{"test.teal"}
	offsetToLine := map[int]int{
		1: 1,
		2: 2,
		5: 3,
	}
	actualSourceMap := GetSourceMap(sourceNames, offsetToLine)

	a.Equal(sourceMapVersion, actualSourceMap.Version)
	a.Equal(sourceNames, actualSourceMap.Sources)
	a.Equal([]string{}, actualSourceMap.Names)

	// Check encoding for `mappings` field.
	splitMappings := strings.Split(actualSourceMap.Mappings, ";")
	prevLine := 0
	for pc := range splitMappings {
		if line, ok := offsetToLine[pc]; ok {
			a.Equal(MakeSourceMapLine(0, 0, line-prevLine, 0), splitMappings[pc])
			prevLine = line
		} else {
			a.Equal("", splitMappings[pc])
		}
	}
}

func TestVLQ(t *testing.T) {
	partitiontest.PartitionTest(t)
	t.Parallel()
	a := require.New(t)

	a.Equal("AAAA", MakeSourceMapLine(0, 0, 0, 0))
	a.Equal("AACA", MakeSourceMapLine(0, 0, 1, 0))
	a.Equal("AAEA", MakeSourceMapLine(0, 0, 2, 0))
	a.Equal("AAgBA", MakeSourceMapLine(0, 0, 16, 0))
	a.Equal("AAggBA", MakeSourceMapLine(0, 0, 512, 0))
	a.Equal("ADggBD", MakeSourceMapLine(0, -1, 512, -1))
}
