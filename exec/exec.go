package exec

import "strconv"

var numFuncMap = map[int]func(){}

func Register(num int, f func()) {
	numFuncMap[num] = f
}

func Exec(num string) {
	i, _ := strconv.Atoi(num)
	if f, ok := numFuncMap[i]; ok {
		f()
		return
	}
	panic("exec func not register")
}
