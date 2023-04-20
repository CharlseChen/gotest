package main

import "fmt"

type Listener interface {
	OnTeacherComming()
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

type StuZhang3 struct {
	Badthing string
}

func (s *StuZhang3) OnTeacherComming() {
	fmt.Println("zhang3 停止", s.Badthing)
}

func (s *StuZhang3) DoBadthing() {
	fmt.Println("zhang3 正在", s.Badthing)
}

type StuZhang4 struct {
	Badthing string
}

func (s *StuZhang4) OnTeacherComming() {
	fmt.Println("赵4 停止 ", s.Badthing)
}

func (s *StuZhang4) DoBadthing() {
	fmt.Println("赵4 正在", s.Badthing)
}

type ClassMonitor struct {
	listenerList []Listener
}

func (m *ClassMonitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range m.listenerList {
		if listener == l {
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
			break
		}
	}
}

func (m *ClassMonitor) Notify() {
	for _, listener := range m.listenerList {
		listener.OnTeacherComming()
	}
}

func main() {
	s1 := &StuZhang3{
		Badthing: "抄作业",
	}
	s2 := &StuZhang4{
		Badthing: "玩王者荣耀",
	}
	fmt.Println("上课了，但是老师没有来，学生们都在忙自己的事...")
	s1.DoBadthing()
	s2.DoBadthing()

	monitor := new(ClassMonitor)
	monitor.AddListener(s1)
	monitor.AddListener(s2)

	monitor.Notify()
}
