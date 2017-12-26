// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package fold

func FoldString(l []string, apply func(a string, b string) string) (res string) {
	res = l[0]
	for _, v := range l[1:] {
		res = apply(res, v)
	}
	return res
}

func FoldBytes(l []Bytes, apply func(a Bytes, b Bytes) Bytes) (res Bytes) {
	res = l[0]
	for _, v := range l[1:] {
		res = apply(res, v)
	}
	return res
}

func FoldUser(l []User, apply func(a User, b User) User) (res User) {
	res = l[0]
	for _, v := range l[1:] {
		res = apply(res, v)
	}
	return res
}
