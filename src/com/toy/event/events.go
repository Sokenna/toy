package event

var eventByname = make(map[string][]func(...interface{}))

func RegisterEvent(name string, callback func(...interface{})) {
	list := eventByname[name]
	list = append(list, callback)
	eventByname[name] = list
}

func CallEvent(name string, param ...interface{}) {
	list := eventByname[name]
	for _, callback := range list {
		callback(param...)
	}
}
