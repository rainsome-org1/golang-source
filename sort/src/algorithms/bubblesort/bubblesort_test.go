package bubblesort

import "testing"

func TBS1(t *testing.T) {
	values := []int{34, 65, 1, 35465, 23}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 23 || values[2] != 34 ||
	    values[3] != 65 || values[4] != 35465 {
		t.Error("BubbleSort() failes. Got", values,
				"Expected: 1 23 34 65 35465")
	}
}

func TBS2(t *testing.T) {
	values := []int{34, 65, 1, 23, 23}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 23 || values[2] != 23 ||
	    values[3] != 34 || values[4] != 65 {
		t.Error("BubbleSort() failes. Got", values,
			"Expected: 1 23 34 65 35465")
	}
}

func TBS3(t *testing.T) {
	values := []int{34}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failes. Got", values,
			"Expected: 1 23 34 65 35465")
	}
}
