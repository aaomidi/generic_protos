package main

import (
	"fmt"

	"github.com/aaomidi/proto-test/message"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func KeyMetadataName(e protoreflect.Enum) string {
	value := e.Descriptor().Values().ByNumber(e.Number())

	if value == nil {
		panic("sad")
	}

	options := value.Options().(*descriptorpb.EnumValueOptions)
	if !proto.HasExtension(options, message.E_KeyMetadata_Name) {
		panic("sad")
	}

	return proto.GetExtension(options, message.E_KeyMetadata_Name).(string)
}

func Mapperize(t protoreflect.EnumType) map[string]protoreflect.EnumNumber {
	m := make(map[string]protoreflect.EnumNumber)
	values := t.Descriptor().Values()
	for i := 0; i < values.Len(); i++ {
		v := values.Get(i)
		if proto.HasExtension(v.Options(), message.E_KeyMetadata_Name) {
			m[KeyMetadataName(t.New(v.Number()))] = v.Number()
		}
	}
	return m
}

func AnotherKeyMaker[E protoreflect.Enum](val E) string{
	return KeyMetadataName(val)
}

func EnumFromString[E protoreflect.Enum](val string) (E, error) {
	var x E
	t := x.Type()
	m := Mapperize(t)

	number, ok := m[val]
	if !ok {
		return x, fmt.Errorf("not a real %s", x.Descriptor().FullName())
	}

	returnVal, ok := t.New(number).(E)
	if !ok {
		return x, fmt.Errorf("sad")
	}

	return returnVal, nil
}

func AnotherMapperize[E protoreflect.Enum]() {
	var x E
	t := x.Type()
	fmt.Println(Mapperize(t))
}

func main() {
	fmt.Println(Mapperize(message.Pet(0).Type()))

	AnotherMapperize[message.Pet]()

	fmt.Println((message.Pet_DOG))
	fmt.Println((message.Pet_DOG))

	possibleDog, err := EnumFromString[message.MediaPlayer]("notJelly")

	fmt.Println(err)
	spew.Dump(possibleDog)
}
