package state

import (
	"github.com/qts0312/ChaosRPC/pkg/failure"
	"sync"
)

var GlobalChaosState = NewChaosState()

type ChaosState struct {
	testIdMap map[int]bool
	sync.Mutex
}

func NewChaosState() *ChaosState {
	return &ChaosState{
		testIdMap: make(map[int]bool),
	}
}

func (s *ChaosState) Update(testId int, errorCode int) int {
	if testId == -1 {
		s.testIdMap = make(map[int]bool)
		return failure.ErrorNone
	}

	if errorCode == failure.ErrorNone {
		return failure.ErrorNone
	} else {
		s.Lock()
		defer s.Unlock()
		if _, exists := s.testIdMap[testId]; exists {
			return failure.ErrorNone
		} else {
			s.testIdMap[testId] = true
			return errorCode
		}
	}
}
