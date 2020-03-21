package service

import (
	"context"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Status(ctx)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	// testing status
	ok := s == "ok"
	if !ok {
		t.Fatalf("expected service to be ok")
	}
}

func TestGet(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	time := time.Now()
	today := time.Format("02/01/2006")

	// testing today's date
	ok := today == d
	if !ok {
		t.Fatalf("expected dates to be equal")
	}
}

func TestValidate(t *testing.T) {
	srv, ctx := setup()

	b, err := srv.Validate(ctx, "31/12/2019")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	// testing that the date is valid
	if !b {
		t.Fatalf("date should be valid")
	}

	tests := []struct {
		input string
		want  bool
	}{
		{input: "31/31/2019", want: false},
		{input: "12/31/2019", want: false},
	}

	for _, tc := range tests {
		got, err := srv.Validate(ctx, tc.input)
		if err != nil {
			//t.Fatalf("Error: %s", err)
		}
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}

	/*
		// testing an invalid date
		b, err = srv.Validate(ctx, "31/31/2019")
		if b {
			t.Fatalf("date should be invalid")
		}
		// testing a USA date date
		b, err = srv.Validate(ctx, "12/31/2019")
		if b {
			t.Fatalf("USA date should be invalid")
		}
	*/
}

func setup() (srv HelloService, ctx context.Context) {
	return NewBasicHelloService(), context.Background()
}
