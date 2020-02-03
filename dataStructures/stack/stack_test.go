package stack

import "testing"

func TestCount_Empty(t *testing.T) {
	st := ConstructStack()

	expected := 0
	if st.Count() == expected {
		t.Logf("Success! Expected %v, got %v", expected, st.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, st.Count())
	}
}

func TestCount_AddRemove(t *testing.T) {
	st := ConstructStack()
	st.Push(66)
	_, ok := st.Pull()

	expected := 0
	if ok && st.Count() == expected {
		t.Logf("Success! Expected %v, got %v", expected, st.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, st.Count())
	}
}

func TestPull_OneValue(t *testing.T) {
	st := ConstructStack()
	st.Push(66)
	actual, ok := st.Pull()

	expected := 66
	if ok && actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, st.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, st.Count())
	}
}

func TestPull_ManyValues(t *testing.T) {
	st := ConstructStack()
	st.Push(66)
	st.Push(2)
	actual, ok := st.Pull()

	expected := 2
	if ok && actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}

func TestString(t *testing.T) {
	st := ConstructStack()
	st.Push(36)
	st.Push(6)
	st.Push(12)
	actual := st.String()

	expected := "12 6 36"
	if expected == actual {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}


