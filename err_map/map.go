package err_map

var errMap = make(map[string]string)

func Register(uuid string, message string) {
	errMap[uuid] = message
}
