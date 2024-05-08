package main

import (
	"testing"
	"fmt"
	"time"
)

func TestStr(t *testing.T) {
	t.Error("this is a testing")
	t.Skipf("test")
	JsonTest()
	fmt.Println("target")
}

func TestValue(t *testing.T) {
	t.Skipf("test")
}

//func BenchmarkString1(b *testing.B) {
//	var buf strings.Builder
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		buf.WriteString("test")
//		buf.WriteString("1")
//		buf.WriteString("~~")
//	}
//	fmt.Errorf(buf.String())
//}
//
//func BenchmarkString2(b *testing.B) {
//	var s string
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		s = strings.Join([]string{s, "test", "1", "~~"}, "")
//	}
//	fmt.Errorf(s)
//}
//
//func BenchmarkString3(b *testing.B) {
//	var s string
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		s = fmt.Sprintf("%s%s%s%s", s, "test", "1", "~~")
//	}
//	fmt.Errorf(s)
//	//fmt.Sprintf(s)
//}
//
//func BenchmarkString4(b *testing.B) {
//	var s string
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		s += "test" + "1" + "~~"
//	}
//	fmt.Errorf(s)
//}
//
//func BenchmarkString5(b *testing.B) {
//	var buf bytes.Buffer
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		buf.WriteString("test")
//		buf.WriteString("1")
//		buf.WriteString("~~")
//		//_ = buf.String()
//	}
//	fmt.Errorf(buf.String())
//}

func TestTime(t *testing.T) {
	t.Log(int(time.Now().Weekday()))
}
