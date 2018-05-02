package util

// ParseCommandArgs 处理参数
func ParseCommandArgs(args []string) (string, []string) {
	var name string
	if len(args) > 0 {
		name = args[0]
		return name, args[1:]
	}
	return "", args
}
