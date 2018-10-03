package goutil

import "testing"

func TestCapitalize(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hallo", "Hallo"},
		{"hola que tal", "Hola que tal"},
		{"", ""},
	}

	for _, c := range cases {
		got := Capitalize(c.in)
		if got != c.want {
			t.Errorf("Capitalize(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCapitalizeTitle(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"foo and bar", "Foo and Bar"},
	}

	for _, c := range cases {
		got := CapitalizeTitle(c.in)
		if got != c.want {
			t.Errorf("CapitalizeTitle(%q) == %q, want %q", c.in, got, c.want)
		}
	}

	cases = []struct {
		in, want string
	}{
		{"milk and honey", "milk and honey"},
	}

	for _, c := range cases {
		got := CapitalizeTitle(c.in, []string{"honey", "milk"})
		if got != c.want {
			t.Errorf("CapitalizeTitle(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}