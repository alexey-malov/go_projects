package safemap

type SafeMap interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) interface{}
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

type UpdateFunc func(interface{}, bool) interface{}

type safeMap chan command

type command struct {
	action  action
	key     string
	value   interface{}
	updater UpdateFunc
	result  chan interface{}
	data    chan map[string]interface{}
}

type findResult struct {
	value interface{}
	found bool
}

type action int

const (
	insert action = iota
	remove
	find
	length
	update
	end
)

func (sm safeMap) Insert(key string, value interface{}) {
	sm <- command{action: insert, key: key, value: value}
}

func (sm safeMap) Delete(key string) {
	sm <- command{action: remove, key: key}
}

func (sm safeMap) Find(key string) (interface{}, bool) {
	reply := make(chan interface{})
	sm <- command{action: find, key: key, result: reply}
	result := (<-reply).(findResult)
	return result.value, result.found
}

func (sm safeMap) Len() int {
	reply := make(chan interface{})
	sm <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (sm safeMap) Update(key string, updater UpdateFunc) {
	sm <- command{action: update, key: key, updater: updater}
}

func (sm safeMap) Close() map[string]interface{} {
	data := make(chan map[string]interface{})
	sm <- command{action: end, data: data}
	return <-data
}

func New() SafeMap {
	sm := make(safeMap)
	sm.run()
	return sm
}

func (sm safeMap) run() {
	go func() {
		data := make(map[string]interface{})
		for {
			command := <-sm
			switch command.action {
			case insert:
				key := command.key
				data[key] = command.value
			case remove:
				delete(data, command.key)
			case find:
				result := findResult{}
				result.value, result.found = data[command.key]
				command.result <- result
			case length:
				command.result <- len(data)
			case update:
				value, found := data[command.key]
				data[command.key] = command.updater(value, found)
			case end:
				close(sm)

			}
		}
	}()
}

func inc(value interface{}, found bool) interface{} {
	if !found {
		return "Леха любит пирожки"
	}
	intValue, ok := value.(int)
	if ok {
		return intValue + 1
	}
	return value
}
