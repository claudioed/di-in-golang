package pkg

import "github.com/tjarratt/babble"

type WordGenerator struct {
	dic *Dictionary
}

func (gen *WordGenerator) GenerateWord() string {
	return gen.dic.Word()
}

// New WordGenerator Producer
func NewWordGenerator(dic *Dictionary) *WordGenerator  {
	return &WordGenerator{dic:dic}
}

// New Dictionary Producer
func NewDictionary() *Dictionary  {
	return &Dictionary{}
}

// Encapsulate Babbler
type Dictionary struct {
}

// Babbler word generator
func (dic *Dictionary) Word() string {
	return babble.NewBabbler().Babble()
}