package main

import "fmt"

type ProA struct {
	key1, key2, v1 string
}

type ProB struct {
	key1, key2, v2 string
}

type MergedAB struct {
	key1, key2, v1, v2 string
}

func equals(proA *ProA, proB *ProB) bool {
	if proA != nil && proB != nil {
		if proA.key1 == proB.key1 && proA.key2 == proB.key2 {
			return true
		}
	}
	return false
}

func main() {

	aArray := [3]ProA{}
	aArray[0] = ProA{
		key1: "0",
		key2: "0",
		v1:   "proA_0",
	}
	aArray[1] = ProA{
		key1: "1",
		key2: "1",
		v1:   "proA_1",
	}
	aArray[2] = ProA{
		key1: "2",
		key2: "2",
		v1:   "proA_2",
	}
	bArray := [3]ProB{}
	bArray[0] = ProB{
		key1: "0",
		key2: "0",
		v2:   "proB_0",
	}
	bArray[1] = ProB{
		key1: "1",
		key2: "1",
		v2:   "proB_1",
	}
	bArray[2] = ProB{
		key1: "22",
		key2: "22",
		v2:   "proB_2",
	}
	abArray := [9]MergedAB{}
	num := 0
	for i := 0; i < len(aArray); i++ {
		for j := 0; j < len(bArray); j++ {
			isEquals := equals(&aArray[i], &bArray[j])
			if isEquals {
				abArray[num].key1 = aArray[i].key1
				abArray[num].key2 = aArray[i].key2
				abArray[num].v1 = aArray[i].v1
				abArray[num].v2 = bArray[i].v2
				num++
			}
		}
	}
	fmt.Println(abArray)

}
