package safeslice

type UpdateFunc func(interface{}) interface{}

type SafeSlice interface {
	Append(interface{})
	At(index int) interface{}
	Close() []interface{}
	Delete(index int)
	Len() int
	Update(index int, updateFunc UpdateFunc)
}

type safeSlice chan command

type action int

type command struct {
	action  action
	value   interface{}
	index   int
	updater UpdateFunc
	result  chan interface{}
	data    chan []interface{}
}

type getItemResult struct {
	data interface{}
	ok   bool
}

const (
	app action = iota
	remove
	getItem
	length
	update
	end
)

func New() SafeSlice {
	slice := make(safeSlice)
	slice.run()
	return slice
}

func (slice safeSlice) Append(value interface{}) {
	slice <- command{action: app, value: value}
}

func (slice safeSlice) At(index int) interface{} {
	reply := make(chan interface{})
	slice <- command{action: getItem, index: index, result: reply}
	result := (<-reply).(getItemResult)
	if result.ok {
		return result.data
	} else {
		return nil
	}
}

func (slice safeSlice) Close() []interface{} {
	data := make(chan []interface{})
	slice <- command{action: end, data: data}
	return <-data
}

func (slice safeSlice) Delete(index int) {
	slice <- command{action: remove, index: index}
}

func (slice safeSlice) Len() int {
	reply := make(chan interface{})
	slice <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (slice safeSlice) Update(index int, updateFunc UpdateFunc) {
	slice <- command{action: update, index: index, updater: updateFunc}
}

func (slice safeSlice) run() {
	go func() {
		data := make([]interface{}, 0)
		for {
			command := <-slice
			switch command.action {
			case app:
				data = append(data, command.value)
			case getItem:
				if command.index >= 0 && command.index < len(data) {
					command.result <- getItemResult{data[command.index], true}
				} else {
					command.result <- getItemResult{nil, false}
				}
			case remove:
				if command.index >= 0 && command.index < len(data) {
					copy(data[command.index:], data[command.index+1:])
					data[len(data)-1] = nil
					data = data[:len(data)-1]
				}
			case update:
				if command.index >= 0 && command.index < len(data) {
					data[command.index] = command.updater(data[command.index])
				}
			case length:
				command.result <- len(data)
			case end:
				command.data <- data
				close(slice)
				return
			}
		}
	}()
}
