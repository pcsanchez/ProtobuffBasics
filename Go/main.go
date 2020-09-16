package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pcsanchez/ProtobuffBasics/Go/src/complex/complexpb"
	"github.com/pcsanchez/ProtobuffBasics/Go/src/enum/enumpb"
	"github.com/pcsanchez/ProtobuffBasics/Go/src/simple/simplepb"
)

func main() {
	sm := doSimple()
	// readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second Message",
			},
			{
				Id:   3,
				Name: "Third Message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	fmt.Println(em.GetDayOfTheWeek())
}

func jsonDemo(pb proto.Message) {
	smAsString := toJSON(pb)
	fmt.Println(smAsString)
	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Succesfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Coulnd't unmarshal the JSON into the pb stduct", err)
	}
}

func readAndWriteDemo(pb proto.Message) {
	writeToFile("simple.bin", pb)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read Content:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protobuf struct")
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(sm.GetId())

	return &sm
}
