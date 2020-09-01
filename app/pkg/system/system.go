package system

import "runtime"

// LsCPU 获取cpu 核数
func LsCPU() int {
	return runtime.NumCPU()
}

// RunFuncName 获取正在运行的函数名
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
