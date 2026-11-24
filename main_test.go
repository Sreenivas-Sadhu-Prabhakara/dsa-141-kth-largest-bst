package main

import "testing"

func TestKthLargest(t *testing.T) {
	root := tn(5, tn(3, tn(2, nil, nil), tn(4, nil, nil)), tn(8, tn(6, nil, nil), tn(9, nil, nil)))
	wants := []int{9, 8, 6, 5, 4, 3, 2} // descending order
	for k := 1; k <= len(wants); k++ {
		got, ok := KthLargest(root, k)
		if !ok || got != wants[k-1] {
			t.Errorf("KthLargest(%d) = (%d,%v), want %d", k, got, ok, wants[k-1])
		}
	}
	if _, ok := KthLargest(root, 8); ok {
		t.Error("k beyond size should fail")
	}
	if _, ok := KthLargest(nil, 1); ok {
		t.Error("empty tree should fail")
	}
}
