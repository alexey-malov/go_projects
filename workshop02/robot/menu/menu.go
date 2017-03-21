package menu

import (
	"bufio"
	"fmt"
	"github.com/alexey-malov/go_projects/workshop02/robot/command"
	"io"
	"os"
)

type Menu struct {
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
		fmt.Printf("Command '%s' has been already registered\n", shortcut)
		return
	}
	m.items[shortcut] = Item{shortcut, description, command, composable}

}

func (m *Menu) Run(input *bufio.Reader) {
	for {
		cm := m.currentMacro
		if cm != nil {
			fmt.Printf("macro:%s>", cm.shortcut)
		} else {
			fmt.Print(">")
		}
		s, isPrefix, err := input.ReadLine()
		if err == io.EOF {
			break
		}
		if isPrefix {
			fmt.Println("Command is too long, try again")
			continue
		}
		if !m.executeCommand(string(s)) {
			break
		}
	}
}

func (m *Menu) executeCommand(word string) bool {
	m.exit = false
	item, ok := m.items[word]
	if !ok {
		fmt.Println("Unknown command")
	} else {
		cm := m.currentMacro
		if cm != nil && item.composable {
			cm.commands = append(cm.commands, item.command)
			fmt.Printf("Action '%s' was recorded\n", word)
		} else {
			item.command()
		}
	}
	return !m.exit
}

func (m *Menu) ShowInstructions() {
	fmt.Println("Commands list:")
	for _, item := range m.items {
		fmt.Printf("\t%v: %v\n", item.shortcut, item.description)
	}
}

func NewMenu() *Menu {
	m := &Menu{}
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
	if m.currentMacro != nil {
		fmt.Println("Can't begin a new macro while another macro is being recorded. You must complete a current macro first")
		return
	}

	input := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a shortcut or empty string to cancel: ")
	shortcut, err := readLn(input)
	if err != nil || shortcut == "" {
		fmt.Println("Cancelled")
		return
	}
	_, ok := m.items[shortcut]
	if ok {
		fmt.Printf("A command '%s' has been already registered\n", shortcut)
		return
	}

	fmt.Print("Enter description or empty string to cancel: ")
	description, err := readLn(input)
	if err != nil || description == "" {
		fmt.Println("Cancelled")
		return
	}

	m.currentMacro = &macro{shortcut, description, []command.Command{}}
}

func (m *Menu) EndMacro() {
	cm := m.currentMacro
	if cm == nil {
		fmt.Println("Can't end a macro. Try to begin a new one")
		return
	}
	m.currentMacro = nil
	if len(cm.commands) == 0 {
		m.currentMacro = nil
		fmt.Println("A macro is empty. Ignoring it")
		return
	}
	m.AddItem(cm.shortcut, cm.description, command.NewMacroCommand(cm.commands))
	fmt.Printf("A new macro '%s' was recorded\n", cm.shortcut)
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
