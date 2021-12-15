package lists

import (
	"testing"
	
	"github.com/stretchr/testify/suite"
)

type ListSuite struct {
	suite.Suite

	strList List[string]
	intList List[int]
}

func (s *ListSuite) SetupTest(){
	s.strList = List[string]{"str1", "str2"}
	s.intList = List[int]{23, 45}
}

func (s *ListSuite) TestIsGeneric() {
	s.Require().Equal("str1", s.strList[0])
	s.Require().Equal("str2", s.strList[1])

	s.Require().Equal(23, s.intList[0])
	s.Require().Equal(45, s.intList[1])
}

func (s *ListSuite) TestForEachStrings() {
	allStrs := ""

	s.strList.ForEach(func(s string){
		allStrs += s + "|"
	})

	s.Require().Equal("str1|str2|", allStrs)
}

func (s *ListSuite) TestForEachInts(){
	sum := 0

	s.intList.ForEach(func(i int){
		sum += i
	})

	s.Require().Equal(68, sum)
}

func (s *ListSuite) TestForEachIsChainable(){
	sum := 0

	s.intList.
	ForEach(func(i int){
		sum += i
	}).
	ForEach(func(i int){
		sum += 2
	})

	s.Require().Equal(72, sum)
}

func (s *ListSuite) TestMapStrings() {
	s.strList.Map(func(s string) string {
		return "mapped: " + s
	})

	s.Require().Equal(List[string]{"mapped: str1", "mapped: str2"}, s.strList)
}

func (s *ListSuite) TestMapInts() {
	s.intList.Map(func(i int) int {
		return i * 2
	})

	s.Require().Equal(List[int]{46, 90}, s.intList)
}

func (s *ListSuite) TestMapIsChainable(){
	s.intList.
	Map(func(i int) int {
		return i * 2
	}).
	Map(func(i int) int {
		return i * 100
	})
	
	s.Require().Equal(List[int]{4600, 9000}, s.intList)
}

func (s *ListSuite) TestSortStrings(){
	mediumStrList := List[string]{"z", "w", "a", "b"}

	mediumStrList.Sort(func(i, j int) bool {
		return mediumStrList[i] < mediumStrList[j]
	})

	s.Require().Equal(List[string]{"a", "b", "w", "z"}, mediumStrList)
}

func (s *ListSuite) TestSortInts(){
	mediumIntList := List[int]{2, 8, 21, 93, 2013, 92, 0}

	mediumIntList.Sort(func(i, j int) bool {
		return mediumIntList[i] < mediumIntList[j]
	})

	s.Require().Equal(List[int]{0, 2, 8, 21, 92, 93, 2013}, mediumIntList)
}

func (s *ListSuite) TestSortIsChainable(){
	mediumStrList := List[string]{"z", "w", "a", "b"}

	mediumStrList.Sort(func(i, j int) bool {
		return mediumStrList[i] < mediumStrList[j]
	}).
	Map(func(elem string) string{
		return elem + elem
	})

	s.Require().Equal(List[string]{"aa", "bb", "ww", "zz"}, mediumStrList)
}

func (s *ListSuite) TestLen(){
	s.Require().Equal(2, s.strList.Len())
}

func (s *ListSuite) TestReduceSumInts(){
	mediumIntList := List[int]{1,2,3,4,5}
	
	reducedVal := mediumIntList.Reduce(func (prev, curr int) int{
		return prev + curr
	}, 0)

	s.Require().Equal(15, reducedVal)
}


func (s *ListSuite) TestReduceFlattenSlice(){
	intSliceList := List[[]int]{[]int{1,2}, []int{3,4}, []int{5, 6}}
	
	reducedVal := intSliceList.Reduce(func (prev, curr []int) []int{
		return append(prev, curr...)
	}, nil)

	s.Require().Equal([]int{1,2,3,4,5,6}, reducedVal)
}

func (s *ListSuite) TestReduceNoElems(){
	mediumIntList := List[int]{}
	
	reducedVal := mediumIntList.Reduce(func (prev, curr int) int{
		return prev + curr
	}, 23)

	s.Require().Equal(23, reducedVal)
}


func TestList(t *testing.T) {
	suite.Run(t, new(ListSuite))
}
