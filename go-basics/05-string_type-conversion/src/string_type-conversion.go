package main

import "fmt"

func main() {

	/* Raw String Literal
	Back Quote (` `)로 둘러 싸인 문자열을 Raw String Literal 라고 한다
	이 안에 있는 문자열은 별도로 해석되지 않고 Raw String 그대로의 값을 갖는다.
	ex) 문자열 안에 \n이 NewLine으로 해석되지 않는다.
	Back Quote (`)은 복수 라인의 문자열을 표현할 때 자주 사용된다.
	*/
	rawLiternal := `가나다라
\n
	\n마바사`

	/* Interpreted String Literal
	이중인용부호(" ")로 둘러 싸인 문자열은 Interpreted String Literal 라고 한다.
	복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석된다.
	*/

	interLiternal := "가나다라\n" + "마바사"

	fmt.Println(rawLiternal)
	fmt.Println()
	fmt.Println(interLiternal)

	/* Type Conversion
	Go 에서는 타입간 변환을 명시적으로 지정해주어야 함!
	*/

	var a int = 100
	var ua uint = uint(a)
	var fa float32 = float32(a)
	fmt.Println(fa, ua)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes, str2)
}
