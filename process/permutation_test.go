package process

import "testing"

func TestPermutationFunction(t *testing.T){
	testStr := "ab"
	permMaps := make(map[string]bool)

	Perm([]rune(testStr), func(a []rune){
		permMaps[string(a)] = true
	})

	if permMaps["ab"] == false {
		t.Error("Permutation failed")
	}

	if permMaps["ba"] == false {
		t.Error("Permutation failed")
	}
}