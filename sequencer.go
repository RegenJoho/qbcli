package qbcli

import "github.com/inancgumus/screen"

type Sequencer struct {
	sequence       []QuestionHandler
	name           string
	shallClear     bool
	shallClearEach bool
}

func (s *Sequencer) Clear(shallClearAfter, shallClearEachTime bool) *Sequencer {
	s.shallClear = shallClearAfter
	s.shallClearEach = shallClearEachTime
	return s
}

func NewSequencer(name string) *Sequencer {
	return &Sequencer{
		name:     name,
		sequence: make([]QuestionHandler, 0),
	}
}

func NewSequencerWithOptions(name string, handlers ...QuestionHandler) *Sequencer {
	return &Sequencer{
		name:     name,
		sequence: handlers,
	}
}

func (s *Sequencer) Add(handler QuestionHandler) *Sequencer {
	s.sequence = append(s.sequence, handler)
	return s
}

func (s *Sequencer) GetName() string {
	return s.name
}
func (s *Sequencer) GetOptions() []string {
	return []string{}
}

func (s *Sequencer) Ask() error {
	for _, v := range s.sequence {
		err := v.Ask()
		if s.shallClearEach {
			screen.Clear()
			screen.MoveTopLeft()
		}
		if err != nil {
			return err
		}
	}
	if s.shallClear {
		screen.Clear()
		screen.MoveTopLeft()
	}
	return nil
}
