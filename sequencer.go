package qbcli

type Sequencer struct {
	sequence []QuestionHandler
	name     string
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
		if err != nil {
			return err
		}
	}
	return nil
}
