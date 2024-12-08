package spell

import "testing"

func TestUpTo100(t *testing.T) {
	tests := []struct {
		in  uint64
		out string
	}{
		{
			in:  0x0,
			out: "",
		},
		{
			in:  0x1,
			out: "one",
		},
		{
			in:  0x7,
			out: "seven",
		},
		{
			in:  0xa,
			out: "ay",
		},
		{
			in:  0x12,
			out: "twelve",
		},
		{
			in:  0x1A,
			out: "abteen",
		},
		{
			in:  0x1D,
			out: "dibbleteen",
		},
		{
			in:  0x2b,
			out: "twenty bee",
		},
		{
			in:  0xFF,
			out: "fleventy ef",
		},
	}

	for i, test := range tests {
		if got, want := upTo100(test.in), test.out; got != want {
			t.Errorf("incorrect on test %d spelling under 0x100; got %q want %q", i, got, want)
		}
	}
}

func TestUpTo10000(t *testing.T) {
	tests := []struct {
		in  uint64
		out string
	}{
		{
			in:  0x100,
			out: "one bitey",
		},
		{
			in:  0x0FF,
			out: "fleventy ef",
		},
		{
			in:  0x1ED,
			out: "one bitey eckity dee",
		},
		{
			in:  0xABDC,
			out: "atta bee bitey dickety see",
		},
	}

	for i, test := range tests {
		if got, want := upTo10000(test.in), test.out; got != want {
			t.Errorf("incorrect on test %d spelling under 0x10000; got %q want %q", i, got, want)
		}
	}
}
