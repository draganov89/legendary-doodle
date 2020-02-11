package queue

import "testing"

func TestCount_Empty(t *testing.T) {
	qu := ConstructQueue()

	expected := 0
	actual := qu.Count()
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}
func TestCount_RemoveUntilEmpty(t *testing.T) {
	qu := ConstructQueue()
	qu.Enqueue(7)
	qu.Enqueue(77)

	_, ok := qu.Dequeue()
	_, ok = qu.Dequeue()

	if !ok {
		t.Errorf("Fail! Expected %v, got %v", true, ok)
		return 
	}

	expected := 0
	actual := qu.Count()
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}

func TestCount_Many(t *testing.T) {
	qu := ConstructQueue()
	qu.Enqueue(7)
	qu.Enqueue(77)

	expected := 2
	actual := qu.Count()
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}

func TestString(t *testing.T) {
	qu := ConstructQueue()
	qu.Enqueue(7)
	qu.Enqueue(41)
	qu.Enqueue(654)

	expected := "7 41 654"
	actual := qu.String()
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}

func TestString_Empty(t *testing.T) {
	qu := ConstructQueue()

	expected := ""
	actual := qu.String()
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}

func TestDequeue_Empty(t *testing.T){
	qu := ConstructQueue()
	actual, ok := qu.Dequeue()
	if ok {
		t.Errorf("Fail! Expected %v, got %v", false, ok)
		return 
	}


	if actual == nil {
		t.Logf("Success! Expected %v, got %v", nil, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", nil, actual)
	}
}

func TestEnqueDequeue_SingleValue(t *testing.T){
	qu := ConstructQueue()
	qu.Enqueue(77)

	actual, ok := qu.Dequeue()
	if !ok {
		t.Errorf("Fail! Expected %v, got %v", true, ok)
		return 
	}

	count := 0
	if qu.Count() == count {
		t.Logf("Success! Expected %v, got %v", count, qu.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", count, qu.Count())
	}
	
	expected := 77
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}

func TestEnqueDequeue_ManyValues(t *testing.T){
	qu := ConstructQueue()
	qu.Enqueue(12)
	qu.Enqueue(1)
	qu.Enqueue(4)

	actual, ok := qu.Dequeue()
	if !ok {
		t.Errorf("Fail! Expected %v, got %v", true, ok)
		return 
	}

	count := 2
	if qu.Count() == count {
		t.Logf("Success! Expected %v, got %v", count, qu.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", count, qu.Count())
	}
	
	expected := 12
	if actual == expected {
		t.Logf("Success! Expected %v, got %v", expected, actual)
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, actual)
	}
}




