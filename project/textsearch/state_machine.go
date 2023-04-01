package textsearch

import "strings"

const (
	initialState = ">"
	finalState   = "<"
)

type result struct {
	value   strings.Builder
	subType StandardType
}

type state struct {
	transitions []stateTransition
	isFinal     bool
}

type stateTransition struct {
	condition     matcher
	target        *state
	modifyResult  func(result *result)
	isCharIgnored bool // символ не добавляется в результат при этом переходе
	replacement   rune // символ для замены
}

type stateMachine struct {
	keywordType KeywordType
	root        *state

	current *state
	result  result
}

func (m *stateMachine) Reset() {
	m.current = nil
	m.result = result{}
}

func (m *stateMachine) Handle(char rune) bool {
	// автомат еще не начал свою работу, состояние пустое
	if m.current == nil {
		for _, transition := range m.root.transitions {
			if transition.condition.Matches(char) {
				m.handleTransition(char, transition)

				return true
			}
		}

		return false
	}

	for _, transition := range m.current.transitions {
		if transition.condition.Matches(char) {
			m.handleTransition(char, transition)

			return true
		}
	}

	return false
}

func (m *stateMachine) Finish() {
	if m.current != nil && !m.current.isFinal {
		m.Handle(0)
	}
}

func (m *stateMachine) IsFinished() bool {
	return m.current != nil && m.current.isFinal
}

func (m *stateMachine) Get() []Keyword {
	return []Keyword{{
		Value:        m.result.value.String(),
		Type:         m.keywordType,
		StandardType: m.result.subType,
	}}
}

func (m *stateMachine) handleTransition(char rune, transition stateTransition) {
	m.current = transition.target
	if transition.modifyResult != nil {
		transition.modifyResult(&m.result)
	}
	if char == 0 || transition.isCharIgnored {
		return
	}
	if transition.replacement > 0 {
		m.result.value.WriteRune(transition.replacement)
	} else {
		m.result.value.WriteRune(char)
	}
}

func newStateMachine(p pattern) *stateMachine {
	states := make([]state, len(p.nodes)+2)
	states[len(states)-1].isFinal = true
	indices := make(map[string]int, len(p.nodes)+2)

	index := 1
	for key := range p.nodes {
		if key == initialState {
			indices[key] = 0
		} else {
			indices[key] = index
			index++
		}
	}
	indices[finalState] = len(states) - 1

	for key, node := range p.nodes {
		transitions := make([]stateTransition, len(node.transitions))
		for j, transition := range node.transitions {
			transitions[j].condition = transition.condition
			transitions[j].isCharIgnored = transition.isCharIgnored
			transitions[j].replacement = transition.replacement
			transitions[j].modifyResult = transition.modifyResult
			transitions[j].target = &states[indices[transition.target]]
		}

		states[indices[key]].transitions = transitions
	}

	return &stateMachine{
		keywordType: p.keywordType,
		root:        &states[0],
	}
}
