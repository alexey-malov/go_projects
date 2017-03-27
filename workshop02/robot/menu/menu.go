package menu

import (
	"bufio"
	"fmt"
	"github.com/alexey-malov/go_projects/workshop02/robot/command"
)

type Menu struct {
	input        *bufio.Reader
	output       *bufio.Writer
	exit         bool
	items        map[string]Item
	currentMacro *macro
}

type macro struct {
	shortcut    string
	description string
	commands    []command.Command
}

type Item struct {
	shortcut    string
	description string
	command     command.Command
	composable  bool
}

func (m *Menu) AddItem(shortcut, description string, command command.Command) {
	m.addItemInternal(shortcut, description, command, true)
}

func (m *Menu) addItemInternal(shortcut, description string, command command.Command, composable bool) {
	_, ok := m.items[shortcut]
	if ok {
		fmt.Fprintf(m.output, "Command '%s' has been already registered\n", shortcut)
		m.output.Flush()
		return
	}
	m.items[shortcut] = Item{shortcut, description, command, composable}

}

func (m *Menu) Run() {
	for {
		cm := m.currentMacro
		if cm != nil {
			fmt.Fprintf(m.output, "macro:%s>", cm.shortcut)
		} else {
			fmt.Fprint(m.output, ">")
		}
		m.output.Flush()

		s, err := readLn(m.input)
		if err != nil {
			break
		}

		if !m.executeCommand(s) {
			break
		}
	}
}

func (m *Menu) executeCommand(word string) bool {
	m.exit = false
	item, ok := m.items[word]
	defer m.output.Flush()
	if !ok {
		fmt.Fprintln(m.output, "Unknown command")
	} else {
		cm := m.currentMacro
		if cm != nil && item.composable {
			cm.commands = append(cm.commands, item.command)
			fmt.Fprintf(m.output, "Action '%s' was recorded\n", word)
		} else {
			item.command()
		}
	}
	return !m.exit
}

func (m *Menu) ShowInstructions() {
	defer m.output.Flush()
	fmt.Fprintln(m.output, "Commands list:")
	for _, item := range m.items {
		fmt.Fprintf(m.output, "\t%v: %v\n", item.shortcut, item.description)
	}
}

func NewMenu(input *bufio.Reader, output *bufio.Writer) *Menu {
	m := &Menu{}
	m.input = input
	m.output = output
	m.items = make(map[string]Item)

	m.addItemInternal("help", "Show instructions", m.ShowInstructions, false)
	m.addItemInternal("exit", "Exit from this menu", m.Exit, false)
	m.addItemInternal("begin_macro", "Records a new macro command", m.BeginMacro, false)
	m.addItemInternal("end_macro", "Stops recording a macro command", m.EndMacro, false)

	return m
}

func (m *Menu) Exit() {
	m.exit = true
}

func (m *Menu) BeginMacro() {
	defer m.output.Flush()
	if m.currentMacro != nil {
		fmt.Fprintln(m.output, "Can't begin a new macro while another macro is being recorded. You must complete a current macro first")
		return
	}

	fmt.Fprint(m.output, "Enter a shortcut or empty string to cancel: ")
	m.output.Flush()
	shortcut, err := readLn(m.input)
	if err != nil || shortcut == "" {
		fmt.Fprintln(m.output, "Cancelled")
		return
	}
	_, ok := m.items[shortcut]
	if ok {
		fmt.Fprintf(m.output, "A command '%s' has been already registered\n", shortcut)
		return
	}

	fmt.Fprint(m.output, "Enter description or empty string to cancel: ")
	m.output.Flush()
	description, err := readLn(m.input)
	if err != nil || description == "" {
		fmt.Fprintln(m.output, "Cancelled")
		return
	}

	m.currentMacro = &macro{shortcut, description, []command.Command{}}
}

func (m *Menu) EndMacro() {
	defer m.output.Flush()
	cm := m.currentMacro
	if cm == nil {
		fmt.Fprintln(m.output, "Can't end a macro. Try to begin a new one")
		return
	}
	m.currentMacro = nil
	if len(cm.commands) == 0 {
		m.currentMacro = nil
		fmt.Fprintln(m.output, "A macro is empty. Ignoring it")
		return
	}
	m.AddItem(cm.shortcut, cm.description, command.NewMacroCommand(cm.commands))
	fmt.Fprintf(m.output, "A new macro '%s' was recorded\n", cm.shortcut)
	m.ShowInstructions()
}

func readLn(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
