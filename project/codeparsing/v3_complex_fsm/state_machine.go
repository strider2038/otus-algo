package v3_complex_fsm

import (
	"github.com/strider2038/otus-algo/project/codeparsing/code"
)

const (
	initialState = ">" // специальный символ для обозначения начального состояния автомата
	finalState   = "<" // специальный символ для обозначения конечного состояния автомата
)

// result - результат работы парсера (конечного автомата).
type result struct {
	start       int
	end         int
	keywordType code.KeywordType
	subType     code.StandardType
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
	start         bool                 // текущая позиция помечается как начало целевого отрезка
	isCharIgnored bool                 // символ не добавляется в результат при этом переходе
}

// stateMachine - конечный автомат для разбора строки.
type stateMachine struct {
	root *state
}

// newStateMachine - создает на основе человеко-читаемой конфигурации паттерна
// конечный автомат с деревом переходов.
func newStateMachine(nodes pattern) *stateMachine {
	states := make([]state, len(nodes)+2)
	states[len(states)-1].isFinal = true
	indices := make(map[string]int, len(nodes)+2)

	index := 1
	for key := range nodes {
		if key == initialState {
			indices[key] = 0
		} else {
			indices[key] = index
			index++
		}
	}
	indices[finalState] = len(states) - 1

	for key, nodeTransitions := range nodes {
		transitions := make([]stateTransition, len(nodeTransitions))
		for j, transition := range nodeTransitions {
			transitions[j].condition = transition.condition
			transitions[j].start = transition.start
			transitions[j].isCharIgnored = transition.isCharIgnored
			transitions[j].modifyResult = transition.modifyResult
			transitions[j].target = &states[indices[transition.target]]
		}

		states[indices[key]].transitions = transitions
	}

	return &stateMachine{root: &states[0]}
}

// Start - инициализирует парсер с пустым состоянием.
func (m *stateMachine) Start() *blockParser {
	p := &blockParser{root: m.root, current: m.root}

	return p
}

// blockParser - часть конечного автомата, хранящая текущее состояние разбора строки.
type blockParser struct {
	root    *state
	current *state
	result  result
}

// Handle - принимает индекс и следующий символ. Если символ успешно принят (конечный автомат
// совершил переход в следующее состояние), то возвращается true, иначе false.
// При получении false обработку следует прервать.
// После прерывания необходимо проверить перешел ли автомат в конечное состояние
// вызовом метода IsFinished.
func (p *blockParser) Handle(index int, char rune) bool {
	// автомат еще не начал свою работу
	if p.current == p.root {
		p.result.start = index
	}

	for _, transition := range p.current.transitions {
		if transition.condition.Matches(char) {
			p.handleTransition(index, char, transition)

			return true
		}
	}

	return false
}

// IsFinished - возвращает признак перешел ли автомат в конечное состояние.
// Если перешел, то парсинг успешно завершен и из автомата можно извлечь ключевые слова
// с помощью метода Get.
func (p *blockParser) IsFinished() bool {
	return p.current != nil && p.current.isFinal
}

// Get - извлекает из текста ключевые слова исходя из состояния автомата.
// Следует вызывать только после проверки достижения конечного состояния IsFinished.
func (p *blockParser) Get(text []rune) []code.Keyword {
	return []code.Keyword{{
		Value:        string(text[p.result.start : p.result.end+1]),
		Type:         p.result.keywordType,
		StandardType: p.result.subType,
	}}
}

func (p *blockParser) handleTransition(index int, char rune, transition stateTransition) {
	p.current = transition.target
	if transition.modifyResult != nil {
		transition.modifyResult(&p.result)
	}
	if char == 0 || transition.isCharIgnored {
		return
	}
	if transition.start {
		p.result.start = index
	}
	p.result.end = index
}
