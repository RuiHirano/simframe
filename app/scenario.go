package app

//"reflect"

type IScenario interface {
	GetArea() IArea
	GetClock() IClock
	GetAgents() []IAgent
	GetModelMap() ModelMap
	RegisterModel(modelList []Model) 
}

type ModelMap map[string]Model
type Model func(id string, position *Position)IAgent

type Scenario struct {
	Area IArea
	Clock IClock
	Agents []IAgent
	ModelMap ModelMap
}

func NewScenario(agents []IAgent, area IArea, clock IClock) IScenario {

	scenario := &Scenario{
		Area: area,
		Clock: clock,
		Agents: agents,
		ModelMap: ModelMap{},
	}

	return scenario
}

func (sn *Scenario) GetArea() IArea{
	return sn.Area
}

func (sn *Scenario) GetClock() IClock{
	return sn.Clock
}

func (sn *Scenario) GetAgents() []IAgent{
	return sn.Agents
}

func (sn *Scenario) GetModelMap() ModelMap{
	return sn.ModelMap
}

func (sn *Scenario) RegisterModel(models []Model){
	for _, modelFunc := range models{
		md := modelFunc("", &Position{X: 0, Y: 0})
		//mdType := reflect.TypeOf(md).String()
		sn.ModelMap[md.GetName()] = modelFunc
	}
}