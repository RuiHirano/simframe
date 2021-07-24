package engine

type ISyncronizer interface {
	Run()
}

type Syncronizer struct {
}

func NewSyncronizer() *Syncronizer {

	sy := &Syncronizer{}

	return sy
}

func (sy *Syncronizer) Run() {
}