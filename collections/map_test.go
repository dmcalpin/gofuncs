package collections

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mapSuite struct {
	suite.Suite
}

func (s *mapSuite) TestKeys() {
	mp := Map[string, int]{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	keys := mp.Keys()
	sort.Strings(keys)
	s.Require().Equal([]string{"a", "b", "c"}, keys)
}

func (s *mapSuite) TestValues() {
	mp := Map[string, int]{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	values := mp.Values()
	sort.Ints(values)
	s.Require().Equal([]int{1, 2, 3}, values)
}

func (s *mapSuite) TestFilter() {
	mp := Map[string, int]{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	filteredMap := mp.Filter(func(key string, value int) bool {
		return value%2 != 0
	})

	s.Require().Equal(Map[string, int]{
		"a": 1,
		"c": 3,
	}, filteredMap)
}

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(mapSuite))
}
