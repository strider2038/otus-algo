package textsearch

import "strings"

type state struct {
	transitions []stateTransition
	isFinal     bool
}

type stateTransition struct {
	condition     matcher
	target        *state
	isCharIgnored bool // символ не добавляется в результат при этом переходе
	replacement   rune // символ для замены
}

type stateMachine struct {
	keywordType KeywordType
	root        *state

	current  *state
	keyword  strings.Builder
	onFinish func(keyword Keyword)
}

func (m *stateMachine) Handle(char rune) {
	// автомат еще не начал свою работу, состояние пустое
	if m.current == nil {
		for _, transition := range m.root.transitions {
			if transition.condition.Matches(char) {
				m.handleTransition(char, transition)

				break
			}
		}

		return
	}

	for _, transition := range m.current.transitions {
		if transition.condition.Matches(char) {
			m.handleTransition(char, transition)

			return
		}
	}

	// если условия не подошли, но текущее состояние является конечным, то совпадение найдено
	if m.current.isFinal {
		m.finishKeyword()
	}

	m.current = nil
	m.keyword = strings.Builder{}
}

func (m *stateMachine) handleTransition(char rune, transition stateTransition) {
	m.current = transition.target
	if transition.isCharIgnored {
		return
	}

	if transition.replacement > 0 {
		m.keyword.WriteRune(transition.replacement)
	} else {
		m.keyword.WriteRune(char)
	}
}

func (m *stateMachine) Finish() {
	if m.current != nil && m.current.isFinal {
		m.finishKeyword()
	}
}

func (m *stateMachine) finishKeyword() {
	m.onFinish(Keyword{
		Value: m.keyword.String(),
		Type:  m.keywordType,
	})
}

func newStateMachine(p pattern, onFinish func(keyword Keyword)) *stateMachine {
	states := make([]state, len(p.nodes)+1)
	indices := make(map[string]int, len(p.nodes)+1)

	index := 1
	for key := range p.nodes {
		if key == "" {
			indices[key] = 0
		} else {
			indices[key] = index
			index++
		}
	}

	for key, node := range p.nodes {
		transitions := make([]stateTransition, len(node.transitions))
		for j, transition := range node.transitions {
			transitions[j].condition = transition.condition
			transitions[j].isCharIgnored = transition.isCharIgnored
			transitions[j].replacement = transition.replacement
			transitions[j].target = &states[indices[transition.target]]
		}

		states[indices[key]].isFinal = node.isFinal
		states[indices[key]].transitions = transitions
	}

	return &stateMachine{
		keywordType: p.keywordType,
		root:        &states[0],
		onFinish:    onFinish,
	}
}
