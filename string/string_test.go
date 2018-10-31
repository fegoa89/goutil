package string

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

func TestUpperCamelCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"foo and bar", "FooAndBar"},
		{"foo.and.bar", "FooAndBar"},
		{"foo-and-bar", "FooAndBar"},
		{"foo_and_bar", "FooAndBar"},
		{"Foo and bar", "FooAndBar"},
	}

	for _, c := range cases {
		got := UpperCamelCase(c.in)
		if got != c.want {
			t.Errorf("UpperCamelCase(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestLowerCamelCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"foo and bar", "fooAndBar"},
		{"foo.and.bar", "fooAndBar"},
		{"foo-and-bar", "fooAndBar"},
		{"foo_and_bar", "fooAndBar"},
		{"Foo and bar", "fooAndBar"},
	}

	for _, c := range cases {
		got := LowerCamelCase(c.in)
		if got != c.want {
			t.Errorf("LowerCamelCase(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"fooAndBar", "foo_and_bar"},
		{"foo.AndBar", "foo_and_bar"},
		{"Foo-And-Bar", "foo_and_bar"},
		{"foo And Bar", "foo_and_bar"},
		{"你好", "你_好"},
	}

	for _, c := range cases {
		got := ToSnakeCase(c.in)
		if got != c.want {
			t.Errorf("ToSnakeCase(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"fooAndBar", "foo-and-bar"},
		{"foo.AndBar", "foo-and-bar"},
		{"foo-And-Bar", "foo-and-bar"},
		{"foo And Bar", "foo-and-bar"},
		{"你好", "你-好"},
	}

	for _, c := range cases {
		got := ToKebabCase(c.in)
		if got != c.want {
			t.Errorf("ToKebabCase(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestHumanize(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"fooAndBar", "foo and bar"},
		{" foo-and-bar", "foo and bar"},
		{"foo_and_bar", "foo and bar"},
	}

	for _, c := range cases {
		got := Humanize(c.in)
		if got != c.want {
			t.Errorf("Humanize(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	cases := []struct {
		in     string
		length int
		want   string
	}{
		{
			"En un lugar de la Mancha, de cuyo nombre no quiero acordarme",
			27,
			"En un lugar de la Mancha...",
		},
		{
			"King Gizzard & The Lizard Wizard",
			15,
			"King Gizzard...",
		},
	}

	for _, c := range cases {
		got := Truncate(c.in, c.length)
		if got != c.want {
			t.Errorf("Truncate(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsEmptyOrBlank(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{" ", true},
		{"my string", false},
	}

	for _, c := range cases {
		got := IsEmptyOrBlank(c.in)
		if got != c.want {
			t.Error("IsEmptyOrBlank(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{" ", true},
		{"my string", false},
	}

	for _, c := range cases {
		got := IsBlank(c.in)
		if got != c.want {
			t.Error("IsBlank(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{" ", false},
		{"my string", false},
	}

	for _, c := range cases {
		got := IsEmpty(c.in)
		if got != c.want {
			t.Error("IsEmpty(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDeleteWhitespace(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"my string", "mystring"},
		{" my string ", "mystring"},
	}

	for _, c := range cases {
		got := DeleteWhitespace(c.in)
		if got != c.want {
			t.Error("DeleteWhitespace(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
