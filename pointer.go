package tools

func Pointer[T any](value T) *T {
	return &value
}

func PointerSlices[T any](list []T) []*T {
	if len(list) == 0 {
		return nil
	}
	result := make([]*T, len(list))
	for i := 0; i < len(list); i++ {
		result[i] = Pointer(list[i])
	}
	return result
}

func PointerMap[T any](m map[string]T) map[string]*T {
	if len(m) == 0 {
		return nil
	}
	result := make(map[string]*T)
	for k, v := range m {
		result[k] = Pointer(v)
	}
	return result
}
