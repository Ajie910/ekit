package slice

import "reflect"

func IndexOf[T any](src []T, element T) (int, error) {
	for idx, t := range src {
		if reflect.DeepEqual(t, element) {
			return idx, nil
		}
	}
	return -1, nil
}
