package pkg

import (
	"encoding/json"
	pb "go-getting-started/gen/go/proto"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestProto(t *testing.T) {
	p1 := &pb.Product{
		Id:          1,
		Name:        "Product 1",
		Description: "Description 1",
		Price:       100,
	}
	p1Bytes, err := proto.Marshal(p1)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("p1Bytes: %v", p1Bytes)
	t.Logf("len(p1Bytes): %v", len(p1Bytes))

	jsonBytes, err := json.Marshal(p1)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Logf("jsonBytes: %v", jsonBytes)
	t.Logf("len(jsonBytes): %v", len(jsonBytes))
}
