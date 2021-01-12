package script

import "github.com/traefik/yaegi/interp"

func GoScript() string {
	i := interp.New(interp.Options{})

	_, err := i.Eval("package foo\nfunc Bar(s string) string { return s + \"-Foo\" }")
	if err != nil {
		panic(err)
	}

	v, err := i.Eval("foo.Bar")
	if err != nil {
		panic(err)
	}

	bar := v.Interface().(func(string) string)

	r := bar("Kung")
	return r
}
