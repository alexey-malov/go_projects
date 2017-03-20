package command

type Command func()

func NewMacroCommand(commands []Command) Command {
	return func() {
		for _, c := range commands {
			c()
		}
	}
}
