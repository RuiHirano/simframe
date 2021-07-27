package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type IdsData struct{
	TargetIds []string
	GetIds []string
	Channel chan string
}

type ISyncronizer interface {
	RegisterIds(key string, ids []string)
	Wait(key string)
	AddId(key string, id string)
	Clean(key string)
}

type Syncronizer struct {
	IdsMap map[string]*IdsData
}

func NewSyncronizer() *Syncronizer {

	sy := &Syncronizer{
		IdsMap: map[string]*IdsData{},
	}

	return sy
}
func (sy *Syncronizer) RegisterIds(key string, ids []string) {
	fmt.Printf("register ids: %v\n", ids)
	sy.IdsMap[key] = &IdsData{
		TargetIds: ids,
		GetIds: []string{},
		Channel: make(chan string, 10),
	}
}

func (sy *Syncronizer) Wait(key string) {
	var eg errgroup.Group
	//mu.Lock()
	ch := sy.IdsMap[key].Channel
	//mu.Unlock()
	eg.Go(func() error {
		for {
			select {
			case id, _ := <-ch:
				fmt.Printf("get id: %s\n", id)
				sy.IdsMap[key].GetIds = append(sy.IdsMap[key].GetIds, id)
				if sy.finishWait(sy.IdsMap[key].GetIds, sy.IdsMap[key].TargetIds) {
					fmt.Printf("get all ids\n")
					return nil
				}
			}
		}
	})
	
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

func (sy *Syncronizer) AddId(key string, id string) {
	idsData, ok := sy.IdsMap[key]
	if ok {
		ch := idsData.Channel
		ch <- id
	}
	
}

func (sy *Syncronizer) Clean(key string) {
	delete(sy.IdsMap, key)
}

func (sy *Syncronizer) finishWait(ids []string, targetIds []string) bool {
	for _, tid := range targetIds {
		isExist := false
		for _, id := range ids{
			if id == tid{
				isExist = true
			}
		}
		if isExist == false {
			return false
		}
	}
	return true
}

func main(){
	syn := NewSyncronizer()
	syn.RegisterIds("test", []string{"1", "2", "3"})
	syn.RegisterIds("test2", []string{"1"})
	go func(){
		time.Sleep(2* time.Second)
		syn.AddId("test", "1")
		time.Sleep(2* time.Second)
		syn.AddId("test", "2")
		time.Sleep(2* time.Second)
		syn.AddId("test", "3")
		time.Sleep(2* time.Second)
		syn.AddId("test2", "1")
	}()
	syn.Wait("test")
	syn.Wait("test2")
	fmt.Printf("finished\n")


}
