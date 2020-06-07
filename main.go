package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "example.com/chethan/src/complex"
	enumpb "example.com/chethan/src/enum_example"
	simpleproto "example.com/chethan/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()

	// readAndWriteDemo(sm)
	jsonDemo(sm)

	//enum_example
	doEnum()

	doComplex()

}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "first message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {

	em := enumpb.EnumMessage{
		Id:  108,
		Day: enumpb.DayOfWeek_SUNDAY,
	}

	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smString := toJSON(sm)
	fmt.Println("converted to JSON...", smString)

	sm1 := &simpleproto.SimpleMessage{}
	fromJSON(smString, sm1)
	fmt.Println("successfully converted from JSON to pb struct....", sm1)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("cannot convert to JSON format", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("cannot unmarshall JSON to protobuf struct")
	}
}

func readAndWriteDemo(sm proto.Message) {

	writeToFile("simple.bin", sm)
	sm2 := &simpleproto.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println("Read the content...", sm2)
}
func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("cant serialize into bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("cant write to file", err)
		return err
	}

	fmt.Println("Data has been written...")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Cannot read from file", err)
		return err
	}
	err1 := proto.Unmarshal(in, pb)
	if err1 != nil {
		log.Fatalln("Couldn't put the bytes into protocol buffer struct", err)
		return err1
	}
	return nil
}

func doSimple() *simpleproto.SimpleMessage {
	sm := simpleproto.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Sample message",
		SampleList: []int32{1, 4, 7, 9},
	}
	// fmt.Println(sm)

	sm.Name = "i renamed you"
	// fmt.Println(sm)

	// fmt.Println("id is ", sm.GetId())

	return &sm

}
