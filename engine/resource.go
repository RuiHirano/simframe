package engine

import (
	ap "github.com/RuiHirano/simframe/app"
)

type IResource interface {
	Run()
}

type Resource struct {
	App ap.IApp
}

func NewResource(app ap.IApp) *Resource {

	resource := &Resource{
		App: app,
	}

	return resource
}

func (rsc *Resource) CreateResource() {
	
}