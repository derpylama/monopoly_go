package test

func Add(a int, b int) int {
	return a + b
}

type testStruct struct {
	IntField int
}

type testStruct2 struct {
	Struct1 testStruct
}

func NewTestStruct2(intField int) testStruct2 {
	return testStruct2{
		Struct1: testStruct{IntField: intField},
	}
}
