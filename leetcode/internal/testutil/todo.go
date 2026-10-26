package testutil

// SkipTODO 应直接 defer 在测试函数中，仅跳过模板的 panic("TODO")。
// 其他 panic 原样抛出，避免掩盖越界、空指针等实现错误。
func SkipTODO(t interface {
	Helper()
	Skip(...any)
}) {
	if value := recover(); value != nil {
		if message, ok := value.(string); ok && message == "TODO" {
			t.Helper()
			t.Skip("尚未实现：请完成 solution.go 并移除 panic(\"TODO\")，然后重新运行测试")
			return
		}
		panic(value)
	}
}
