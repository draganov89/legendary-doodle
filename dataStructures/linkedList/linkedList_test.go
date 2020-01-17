package linkedList

import "testing"

func TestString(t *testing.T){
	ll := ConstructLinkedList()
	ll.AddFirst(7)
	ll.AddFirst(17)
	ll.AddFirst(77)

	expected := "77 17 7"
	if ll.String() == expected{
		t.Logf("Success! Expected %v, got %v",expected, ll.String())
	} else {
		t.Errorf("Fail! Expected %v, got %v",expected, ll.String())
	}
}

func TestString_ZeroItems(t *testing.T){
	ll := ConstructLinkedList()

	expected := ""
	if ll.String() == expected{
		t.Logf("Success! Expected \"%v\", got \"%v\"",expected, ll.String())
	} else {
		t.Errorf("Fail! Expected \"%v\", got \"%v\"",expected, ll.String())
	}
}

func TestCount_SingleElement(t *testing.T){
	ll := ConstructLinkedList()

	ll.AddFirst(7)

	if ll.Count() == 1 {
		t.Logf("Success! Expected %v, got %v", 1, ll.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", 1, ll.Count())
	}
}

func TestCount_ManyElement(t *testing.T){
	ll := ConstructLinkedList()

	ll.AddFirst(7)
	ll.AddFirst(17)
	ll.AddFirst(77)

	expected := 3
	if ll.Count() == expected {
		t.Logf("Success! Expected %v, got %v", expected, ll.Count())
	} else {
		t.Errorf("Fail! Expected %v, got %v", expected, ll.Count())
	}
}

func TestRemoveFirst_Empty(t *testing.T){
	ll := ConstructLinkedList()
	removed, err := ll.RemoveFirst()

	if	err != nil {
		t.Logf("Success! Expected Error, got Error: %v", err)
	} else {
		t.Errorf("Fail! Expected Error, got %v", removed)
	}
}

func TestRemoveFirst_SingleElement(t *testing.T){
	ll := ConstructLinkedList()
	expected := 7
	ll.AddFirst(expected)
	removed, err := ll.RemoveFirst()

	if	err != nil {
		t.Errorf("Fail! Expected %v, got Error: %v", expected, err)
	} else {
		if removed == expected {
			t.Logf("Success! Expected %v, got Error: %v",expected, removed)
		} else {
			t.Errorf("Fail! Expected %v, got %v", expected, removed)
		}
	}
}

func TestRemoveFirst_MultipleElements(t *testing.T){
	ll := ConstructLinkedList()
	ll.AddFirst(12)
	ll.AddFirst(7)
	removed, err := ll.RemoveFirst()
	
	listremained := "12"
	listCount := 1
	expected := 7
	if	err != nil {
		t.Errorf("Fail! Expected %v, got Error: %v", expected, err)
	} else {
		if removed == expected && ll.String() == listremained && ll.Count() == listCount{
			t.Logf("Success! Expected %v, got Error: %v",expected, removed)
		} else {
			t.Errorf("Fail! Expected %v, got %v", expected, removed)
		}
	}
}

func TestRemoveLast_Empty(t *testing.T){
	ll := ConstructLinkedList()
	removed, err := ll.RemoveLast()

	if	err != nil {
		t.Logf("Success! Expected Error, got Error: %v", err)
	} else {
		t.Errorf("Fail! Expected Error, got %v", removed)
	}
}

func TestRemoveLast_SingleElement(t *testing.T){
	ll := ConstructLinkedList()
	expected := 17
	ll.AddFirst(expected)
	removed, err := ll.RemoveLast()

	if	err != nil {
		t.Errorf("Fail! Expected %v, got Error: %v", expected, err)
	} else {
		if removed == expected {
			t.Logf("Success! Expected %v, got Error: %v",expected, removed)
		} else {
			t.Errorf("Fail! Expected %v, got %v", expected, removed)
		}
	}
}

func TestRemoveLast_MultipleElements(t *testing.T){
	ll := ConstructLinkedList()
	ll.AddFirst(12)
	ll.AddFirst(7) // 7 12
	removed, err := ll.RemoveLast()
	
	expected := 12
	listremained := "7"
	listCount := 1
	if	err != nil {
		t.Errorf("Fail! Expected %v, got Error: %v", expected, err)
	} else {
		if removed == expected && ll.String() == listremained && ll.Count() == listCount {
			t.Logf("Success! Expected %v, got Error: %v",expected, removed)
		} else {
			t.Errorf("Fail! Expected %v, got %v", expected, removed)
		}
	}
}

// func TestContains(t *testing.T) {
// 	ll := ConstructLinkedList()
// 	ll.AddFirst(7)
// 	ll.AddFirst(17)
// 	ll.AddFirst(77)

// 	contained := ll.Contains(17)
// 	expected := true
// 	if expected == contained {
// 		t.Logf("Success! Expected %v, got %v",expected, contained)
// 	} else {
// 		t.Errorf("Fail! Expected %v, got %v", expected, contained)
// 	}
// }