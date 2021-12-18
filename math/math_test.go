package math

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mathSuite struct {
	suite.Suite
}

func (s *mathSuite) TestAbs() {
	// don't even need the type declaration
	// go can infer the type
	s.Require().Equal(3, Abs(-3))
	s.Require().Equal(3, Abs(3))
	s.Require().Equal(10.8, Abs(-10.8))
	s.Require().Equal(int32(97), Abs(-'a'))
}

func (s *mathSuite) TestAcos() {
	s.Require().Equal(1.6709637479564563, Acos(-0.1))
	s.Require().Equal(float32(0.20033474), Acos(float32(.98)))
	s.Require().Equal(0.2003348423231197, Acos(.98))
	s.Require().Equal(0.0, Acos(1.0))
}

func (s *mathSuite) TestMax() {
	s.Require().Equal(1, Max(1, 0))
	s.Require().Equal(5.2, Max(3.0, 5.2))
}

func (s *mathSuite) TestMod() {
	s.Require().Equal(1, Mod(3, 2))
	s.Require().Equal(1, Mod(3.0, 2.0))
	s.Require().Equal(5, Mod(97, 23))
}

func TestMathSuite(t *testing.T) {
	suite.Run(t, new(mathSuite))
}
