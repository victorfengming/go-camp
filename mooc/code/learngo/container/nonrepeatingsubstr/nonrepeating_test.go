package nonrepeatingsubstr

import "testing"

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	s := "黑化肥发会会化肥发会黑化肥发会黑化肥发会黑化肥发会会会发"
	for i := 0; i < 10; i++ {
		s = s + s
	}
	res := lengthOfNonRepeatingSubStr(s)
	//res := lengthOfNonRepeatingSubStr("lengthOfNonRepeatingSubStr")
	print(res)
}

func BenchmarkSubstr(b *testing.B) {
	//b.StartTimer()
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 8
	b.Logf("len() =- %d", len(s))
	for i := 0; i < b.N; i++ {
		res := lengthOfNonRepeatingSubStr(s)
		if res != ans {
			b.Errorf("got %d %s %d", res, s, ans)
		}
	}
	/**
	goos: windows
	goarch: amd64
	pkg: learngo/container/nonrepeatingsubstr
	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	BenchmarkSubstr
	    nonrepeating_test.go:22: len() =- 2752512
	    nonrepeating_test.go:22: len() =- 2752512
	    nonrepeating_test.go:22: len() =- 2752512
	BenchmarkSubstr-12    	      50	  23185836 ns/op
	PASS

	Process finished with the exit code 0
	*/

	// 优化后 >>>
	/**
	goos: windows
	goarch: amd64
	pkg: learngo/container/nonrepeatingsubstr
	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
	BenchmarkSubstr
	    nonrepeating_test.go:22: len() =- 491520
	    nonrepeating_test.go:22: len() =- 491520
	    nonrepeating_test.go:22: len() =- 491520
	BenchmarkSubstr-12    	     596	   1741422 ns/op
	PASS

	Process finished with the exit code 0

	*/
}
