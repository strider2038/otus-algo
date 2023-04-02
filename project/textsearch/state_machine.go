package textsearch

import (
	"strings"

	"github.com/strider2038/otus-algo/project/textsearch/code"
)

const (
	initialState = ">" // специальный символ для обозначения начального состояния автомата
	finalState   = "<" // специальный символ для обозначения конечного состояния автомата
)

// result - результат работы парсера (конечного автомата).
type result struct {
	value   strings.Builder
	subType code.StandardType
}

// state - состояние конечного автомата.
type state struct {
	transitions []stateTransition
	isFinal     bool
}

// stateTransition - параметры перехода в следующее состояние.
type stateTransition struct {
	condition matcher // условие для осуществления перехода
	target    *state  // следующее состояние автомата

	modifyResult  func(result *result) // замыкание для модификации результата
	isCharIgnored bool                 // символ не добавляется в результат при этом переходе
	replacement   rune                 // символ для замены
}

// stateMachine - конечный автомат для разбора строки.
type stateMachine struct {
	keywordType code.KeywordType
	root        *state
}

// newStateMachine - создает на основе человеко-читаемой конфигурации паттерна
// конечный автомат с деревом переходов.
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

	for key, nodeTransitions := range p.nodes {
		transitions := make([]stateTransition, len(nodeTransitions))
		for j, transition := range nodeTransitions {
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

// Start - инициализирует парсер с пустым состоянием.
func (m *stateMachine) Start() *blockParser {
	return &blockParser{keywordType: m.keywordType, root: m.root}
}

// blockParser - часть конечного автомата, хранящая текущее состояние разбора строки.
type blockParser struct {
	keywordType code.KeywordType
	root        *state
	current     *state
	result      result
}

// Handle - принимает следующий символ. Если символ успешно принят (конечный автомат
// совершил переход в следующее состояние), то возвращается true, иначе false.
// При получении false обработку следует прервать.
// После прерывания необходимо проверить перешел ли автомат в конечное состояние
// вызовом метода IsFinished.
func (p *blockParser) Handle(char rune) bool {
	// автомат еще не начал свою работу, состояние пустое
	if p.current == nil {
		for _, transition := range p.root.transitions {
			if transition.condition.Matches(char) {
				p.handleTransition(char, transition)

				return true
			}
		}

		return false
	}

	for _, transition := range p.current.transitions {
		if transition.condition.Matches(char) {
			p.handleTransition(char, transition)

			return true
		}
	}

	return false
}

// Finish - после передачи последнего символа строки необходимо вызвать этот метод
// для того, чтобы проверить на достижение конечного состояния.
func (p *blockParser) Finish() {
	if p.current != nil && !p.current.isFinal {
		p.Handle(0)
	}
}

// IsFinished - возвращает признак перешел ли автомат в конечное состояние.
// Если перешел, то парсинг успешно завершен и из автомата можно извлечь ключевые слова
// с помощью метода Get.
func (p *blockParser) IsFinished() bool {
	return p.current != nil && p.current.isFinal
}

// Get - извлекает из автомата ключевые слова. Следует вызывать только после проверки
// достижения конечного состояния IsFinished.
func (p *blockParser) Get() []code.Keyword {
	return []code.Keyword{{
		Value:        p.result.value.String(),
		Type:         p.keywordType,
		StandardType: p.result.subType,
	}}
}

func (p *blockParser) handleTransition(char rune, transition stateTransition) {
	p.current = transition.target
	if transition.modifyResult != nil {
		transition.modifyResult(&p.result)
	}
	if char == 0 || transition.isCharIgnored {
		return
	}
	if transition.replacement > 0 {
		p.result.value.WriteRune(transition.replacement)
	} else {
		p.result.value.WriteRune(char)
	}
}
