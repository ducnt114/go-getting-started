package service

import (
	"context"
	"testing"
)

func TestCtxTimeout(t *testing.T) {
	s := &bookServiceImpl{}
	if err := s.GetData(context.Background()); err != nil {
		t.Error(err)
		t.Fail()
	}
}
