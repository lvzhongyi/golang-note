package main
//字符串
func main() {
	var s1 = "abcdefg"
	println(s1)
	s2 := []byte(s1)
	s2[0] = 't'
	println(string(s2))

	s3 := "abcedfg"
	s3 = "z" + s3[1:]
	println(s3)

	s4 :=`hello
	world`
	println(s4)

}
