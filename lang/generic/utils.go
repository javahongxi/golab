package generic

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}
	return us
}

func Filter[T any](ts []T, f func(T) bool) []T {
	var us []T
	for _, t := range ts {
		if f(t) {
			us = append(us, t)
		}
	}
	return us
}

func Reduce[T, U any](ts []T, f func(U, T) U, initial U) U {
	result := initial
	for _, t := range ts {
		result = f(result, t)
	}
	return result
}

func Find[T any](ts []T, f func(T) bool) (T, bool) {
	for _, t := range ts {
		if f(t) {
			return t, true
		}
	}
	var zero T
	return zero, false
}

func Contains[T comparable](ts []T, t T) bool {
	for _, v := range ts {
		if v == t {
			return true
		}
	}
	return false
}

func Unique[T comparable](ts []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, t := range ts {
		if !seen[t] {
			seen[t] = true
			result = append(result, t)
		}
	}
	return result
}

func Reverse[T any](ts []T) []T {
	result := make([]T, len(ts))
	for i, t := range ts {
		result[len(ts)-1-i] = t
	}
	return result
}

func Chunk[T any](ts []T, size int) [][]T {
	var result [][]T
	for i := 0; i < len(ts); i += size {
		end := i + size
		if end > len(ts) {
			end = len(ts)
		}
		result = append(result, ts[i:end])
	}
	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
