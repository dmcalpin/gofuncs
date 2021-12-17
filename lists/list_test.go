package lists

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type listSuite struct {
	suite.Suite

	strList List[string]
	intList List[int]
}

func (s *listSuite) SetupTest() {
	s.strList = List[string]{"str1", "str2"}
	s.intList = List[int]{23, 45}
}

func (s *listSuite) TestIsGeneric() {
	s.Require().Equal("str1", s.strList[0])
	s.Require().Equal("str2", s.strList[1])

	s.Require().Equal(23, s.intList[0])
	s.Require().Equal(45, s.intList[1])
}

func (s *listSuite) TestForEachStrings() {
	allStrs := ""

	s.strList.ForEach(func(s string) {
		allStrs += s + "|"
	})

	s.Require().Equal("str1|str2|", allStrs)
}

func (s *listSuite) TestForEachInts() {
	sum := 0

	s.intList.ForEach(func(i int) {
		sum += i
	})

	s.Require().Equal(68, sum)
}

func (s *listSuite) TestForEachIsChainable() {
	sum := 0

	s.intList.
		ForEach(func(i int) {
			sum += i
		}).
		ForEach(func(i int) {
			sum += 2
		})

	s.Require().Equal(72, sum)
}

func (s *listSuite) TestMapStrings() {
	updatedList := s.strList.Map(func(s string) string {
		return "mapped: " + s
	})

	s.Require().Equal(List[string]{"mapped: str1", "mapped: str2"}, updatedList)
}

func (s *listSuite) TestMapInts() {
	updatedList := s.intList.Map(func(i int) int {
		return i * 2
	})

	s.Require().Equal(List[int]{46, 90}, updatedList)
}

func (s *listSuite) TestMapIsChainable() {
	updatedList := s.intList.
		Map(func(i int) int {
			return i * 2
		}).
		Map(func(i int) int {
			return i * 100
		})

	s.Require().Equal(List[int]{4600, 9000}, updatedList)
}

func (s *listSuite) TestSortStrings() {
	strList := List[string]{"z", "w", "a", "b"}

	ascList := strList.Sort(func(a, b string) bool {
		return a < b
	})

	s.Require().Equal(List[string]{"a", "b", "w", "z"}, ascList)

	descList := strList.Sort(func(a, b string) bool {
		return a > b
	})

	s.Require().Equal(List[string]{"z", "w", "b", "a"}, descList)
}

func (s *listSuite) TestSortInts() {
	intList := List[int]{2, 8, 21, 93, 2013, 92, 0}

	updatedList := intList.Sort(func(a, b int) bool {
		return a < b
	})

	s.Require().Equal(List[int]{0, 2, 8, 21, 92, 93, 2013}, updatedList)
}

func (s *listSuite) TestSortIsChainable() {
	strList := List[string]{"z", "w", "a", "b"}

	updatedList := strList.
		Sort(func(a, b string) bool {
			return a < b
		}).
		Map(func(elem string) string {
			return elem + elem
		})

	s.Require().Equal(List[string]{"aa", "bb", "ww", "zz"}, updatedList)
}

func (s *listSuite) TestLen() {
	s.Require().Equal(2, s.strList.Len())
}

func (s *listSuite) TestReduceSumInts() {
	intList := List[int]{1, 2, 3, 4, 5}

	reducedVal := intList.Reduce(func(prev, curr int) int {
		return prev + curr
	}, 0)

	s.Require().Equal(15, reducedVal)
}

func (s *listSuite) TestReduceFlattenSlice() {
	intSliceList := List[[]int]{[]int{1, 2}, []int{3, 4}, []int{5, 6}}

	reducedVal := intSliceList.Reduce(func(prev, curr []int) []int {
		return append(prev, curr...)
	}, nil)

	s.Require().Equal([]int{1, 2, 3, 4, 5, 6}, reducedVal)
}

func (s *listSuite) TestReduceNoElems() {
	intList := List[int]{}

	reducedVal := intList.Reduce(func(prev, curr int) int {
		return prev + curr
	}, 23)

	s.Require().Equal(23, reducedVal)
}

func (s *listSuite) TestFilterShortStrings() {
	strList := List[string]{"a", "dog", "walks", "to", "the", "park"}

	updatedList := strList.Filter(func(word string) bool {
		return len(word) > 3
	})

	s.Require().Equal(List[string]{"walks", "park"}, updatedList)
}

func (s *listSuite) TestFill() {
	strList := List[string]{"a", "dog", "walks", "to", "the", "park"}

	updatedList := strList.Fill("[redacted]", 3, 6)

	s.Require().Equal(
		List[string]{"a", "dog", "walks", "[redacted]", "[redacted]", "[redacted]"},
		updatedList,
	)
}

func (s *listSuite) TestReverse() {
	s.Require().Equal(
		List[string]{},
		List[string]{}.Reverse(),
	)

	s.Require().Equal(
		List[string]{"a"},
		List[string]{"a"}.Reverse(),
	)

	s.Require().Equal(
		List[string]{"b", "a"},
		List[string]{"a", "b"}.Reverse(),
	)

	s.Require().Equal(
		List[int]{3, 2, 1},
		List[int]{1, 2, 3}.Reverse(),
	)

	s.Require().Equal(
		List[float64]{5.0, 4.0, 3.0, 2.0, 1.0},
		List[float64]{1.0, 2.0, 3.0, 4.0, 5.0}.Reverse(),
	)
}

func TestList(t *testing.T) {
	suite.Run(t, new(listSuite))
}
