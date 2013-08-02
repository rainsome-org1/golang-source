package qsort

import "testing"

func TSS1(t *testing.T) {
	values := []int{34, 65, 1, 35465, 23}
	QuikSort(values)
	if values[0] != 1 || values[1] != 23 || values[2] != 34 ||
	    values[3] != 65 || values[4] != 35465 {
		t.Error("QuikSort() failes. Got", values,
				"Expected: 1 23 34 65 35465")
	}
}

func TSS2(t *testing.T) {
	values := []int{34, 65, 1, 23, 23}
	QuikSort(values)
	if values[0] != 1 || values[1] != 23 || values[2] != 23 ||
	    values[3] != 34 || values[4] != 65 {
		t.Error("QuikSort() failes. Got", values,
			"Expected: 1 23 34 65 35465")
	}
}

func TSS3(t *testing.T) {
	values := []int{34}
	QuikSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failes. Got", values,
			"Expected: 1 23 34 65 35465")
	}
}
