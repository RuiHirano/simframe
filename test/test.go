package main

import (
	"fmt"
	"reflect"
)

type IScenario interface {
	GetId() string
}

type Scenario struct {
	Id string
}

func NewScenario(id string) IScenario {
	return &Scenario{
		Id:id,
	}
}

func (s *Scenario) GetId() string {

	return s.Id + s.GetId2()
}

func (s *Scenario) GetId2() string {
	return s.Id+"test"
}

type Scenario2 struct {
	Id string
}

func NewScenario2(id string) IScenario {
	return &Scenario2{
		Id:id,
	}
}

func (s *Scenario2) GetId() string {

	return s.Id + "100"
}

func main(){
	sn := NewScenario("43")
	P(sn)
	fmt.Printf(sn.GetId())
	var sn2 IScenario
	P(sn2)
	sn2 = sn
	P(sn2)
	sn3 := NewScenario2("43")
	fmt.Printf("sn3", sn3.GetId())

	Map := map[string]reflect.Type{}
	reflect.TypeOf(sn2)
	fmt.Printf(reflect.TypeOf(sn2).String())
	Map[reflect.TypeOf(sn2).String()] = reflect.TypeOf(sn2)
	fmt.Printf("%v", Map[reflect.TypeOf(sn2).String()])


}

func P(t interface{}) {
    fmt.Println(reflect.TypeOf(t))
}