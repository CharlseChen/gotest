package main

import "fmt"

type FeatureOne struct {
}

func (f *FeatureOne) MA() {
	fmt.Println("MA")
}

type FeatureTwo struct{}

func (f *FeatureTwo) MB() {
	fmt.Println("MB")
}

type FeatureThree struct{}

func (f *FeatureThree) MC() {
	fmt.Println("MC")
}

type Facade struct {
	a *FeatureOne
	b *FeatureTwo
	c *FeatureThree
}

func (f *Facade) GA() {
	f.a.MA()
	f.b.MB()
}

func main() {
	f := Facade{
		a: new(FeatureOne),
		b: new(FeatureTwo),
		c: new(FeatureThree),
	}
	f.GA()
}
