package wslice

type Ordered interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Reverse reverses the elements of a slice in place.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Contains checks if a slice contains a specific element.
func Contains[T comparable](s []T, elem T) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

// Sorts elements in place using the heap sort algorithm.
func Sort[T Ordered](s []T) {
	heapify(s)
	for end := len(s) - 1; end > 0; end-- {
		s[0], s[end] = s[end], s[0]
		down(s[:end], 0)
	}
}

func heapify[T Ordered](s []T) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		down(s, i)
	}
}

func down[T Ordered](s []T, i int) {
	for {
		l := 2*i + 1
		if l >= len(s) {
			break
		}
		j := l
		if r := l + 1; r < len(s) && s[r] > s[l] {
			j = r
		}
		if s[i] >= s[j] {
			break
		}
		s[i], s[j] = s[j], s[i]
		i = j
	}
}
