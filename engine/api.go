package engine

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/RuiHirano/simframe/api"
	"github.com/RuiHirano/simframe/app"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type SimulatorAPI struct {
	Port int
	Client api.SimulatorServiceClient
}

func NewSimulatorAPI(port int) *SimulatorAPI {
	sa := &SimulatorAPI{
		Port: port,
	}
	//sa.SetUp()
	return sa
}

func (sa *SimulatorAPI) SetUp() {
	color.Green("Connecting to simulator...\n")
	conn, err := grpc.Dial(fmt.Sprintf(":%d", sa.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	sa.Client = api.NewSimulatorServiceClient(conn)
	color.Green("Success connecting to simulator.\n")
}

func (sa *SimulatorAPI) SetUpSimulator(area app.IArea, clock app.IClock, agents []app.IAgent, neighbors []*Neighbor) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := sa.Client.SetUpSimulator(ctx, &api.SetUpSimulatorRequest{
		Area: toArea(area),
		Clock: toClock(clock),
		Agents: toAgents(agents),
		Neighbors: toNeighbors(neighbors),
	})
	if err != nil {
		log.Fatalf("could not set up: %v", err)
	}
	log.Printf("Get status: %s", r.GetStatus())
}

func (sa *SimulatorAPI) RunSimulator(senderId string ) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := sa.Client.RunSimulator(ctx, &api.RunSimulatorRequest{
		SenderId: senderId,
	})
	if err != nil {
		log.Fatalf("could not run: %v", err)
	}
	log.Printf("Get status: %s", r.GetStatus())
}

func (sa *SimulatorAPI) GetNeighborAgents(senderId string, agents []app.IAgent) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := sa.Client.GetNeighborAgents(ctx, &api.GetNeighborAgentsRequest{
		SenderId: senderId,
		Agents: []*api.Agent{},
	})
	if err != nil {
		log.Fatalf("could not get neighbor agents: %v", err)
	}
}

type SimulatorService struct{
	App app.IApp
	Port int
	SetUpSimulatorHandler func(app.IArea, app.IClock, []app.IAgent, []*Neighbor)
	RunSimulatorHandler func()
	GetNeighborAgentsHandler func(id string, agents []app.IAgent)
	api.UnimplementedSimulatorServiceServer
}

func NewSimulatorService(ap app.IApp, port int, ssh func(app.IArea, app.IClock, []app.IAgent, []*Neighbor), rsh func(), gnah func(id string, agents []app.IAgent)) *SimulatorService{
   ms := &SimulatorService{
	   App: ap,
	   Port: port,
	   SetUpSimulatorHandler: ssh,
	   RunSimulatorHandler: rsh,
	   GetNeighborAgentsHandler: gnah,
   }
   return ms
}

func (es *SimulatorService)  Serve() {
	fmt.Printf("Serve Simulator\n")

	server := grpc.NewServer()
	api.RegisterSimulatorServiceServer(server, es)

	reflection.Register(server)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", es.Port))
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is %d\n", es.Port)
	server.Serve(lis)
}

func (es *SimulatorService) SetUpSimulator(ctx context.Context, request *api.SetUpSimulatorRequest) (*api.SetUpSimulatorResponse, error) {
 
	es.SetUpSimulatorHandler(toIArea(request.GetArea()), toIClock(request.GetClock()),toIAgents(request.GetAgents(), es.App.GetScenarios()[0].GetModelMap()),toINeighbors(request.GetNeighbors()))
 
	response := &api.SetUpSimulatorResponse{
		Status: api.Status_OK,
	}
	return response, nil
 }

func (es *SimulatorService) RunSimulator(ctx context.Context, request *api.RunSimulatorRequest) (*api.RunSimulatorResponse, error) {

   es.RunSimulatorHandler()

   response := &api.RunSimulatorResponse{
	   Status: api.Status_OK,
   }
   return response, nil
}

func (es *SimulatorService) GetNeighborAgents(ctx context.Context, request *api.GetNeighborAgentsRequest) (*api.GetNeighborAgentsResponse, error) {
 
	es.GetNeighborAgentsHandler(request.GetSenderId(), toIAgents(request.GetAgents(), es.App.GetScenarios()[0].GetModelMap()))
 
	response := &api.GetNeighborAgentsResponse{
		Status: api.Status_OK,
		Agents: []*api.Agent{},
	}
	return response, nil
 }


type EngineAPI struct {
	Port int
	App app.IApp
	Client api.EngineServiceClient
}

func NewEngineAPI(port int, ap app.IApp) *EngineAPI {
	ea := &EngineAPI{
		App: ap,
		Port: port,
	}
	//ea.SetUp()
	return ea
}

func (ea *EngineAPI) SetUp() {
	color.Green("Connecting to engine...\n")
	conn, err := grpc.Dial(fmt.Sprintf(":%d", ea.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	ea.Client = api.NewEngineServiceClient(conn)
	color.Green("Success connecting to engine.\n")
}

func (ea *EngineAPI) RegisterSimulator(simId string) (error, int){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := ea.Client.RegisterSimulator(ctx, &api.RegisterSimulatorRequest{
		Id: simId,
	})
	if err != nil {
		log.Fatalf("could not reuest [RegisterSimulator]: %v", err)
		return err, 0
	}
	return nil, int(r.GetPort())
}


type EngineService struct{
	RegisterSimulatorHandler func(id string) int
	Port int
	api.UnimplementedEngineServiceServer
}

func NewEngineService(port int, rsh func(id string) int) *EngineService{
   es := &EngineService{
	   Port: port,
	   RegisterSimulatorHandler: rsh,
   }
   return es
}

func (es *EngineService) Serve() {
	fmt.Printf("Serve Engine\n")
	server := grpc.NewServer()
	api.RegisterEngineServiceServer(server, es)

	reflection.Register(server)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", es.Port))
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("Serving! port is %d\n", es.Port)
	server.Serve(lis)
 }

func (es *EngineService) RegisterSimulator(ctx context.Context, request *api.RegisterSimulatorRequest) (*api.RegisterSimulatorResponse, error) {
   fmt.Printf("getRequest %v\n", request)

   port := es.RegisterSimulatorHandler(request.GetId())

   response := &api.RegisterSimulatorResponse{
	   Status: api.Status_OK,
	   Port: uint64(port),
   }
   return response, nil
}

func toArea(area app.IArea) *api.Area{
	return &api.Area{
		Space: &api.Space{
			MinX: float32(area.GetSpace().MinX),
			MaxX: float32(area.GetSpace().MaxX),
			MinY: float32(area.GetSpace().MinY),
			MaxY: float32(area.GetSpace().MaxY),
		},
	}
}

func toIArea(area *api.Area) app.IArea{
	return app.NewArea(
		&app.Space{
			MinX: float64(area.GetSpace().MinX),
			MaxX: float64(area.GetSpace().MaxX),
			MinY: float64(area.GetSpace().MinY),
			MaxY: float64(area.GetSpace().MaxY),
		},
	)
}

func toPosition(pos *app.Position) *api.Position{
	return &api.Position{
		X: float32(pos.X),
		Y: float32(pos.Y),
	}
}


func toIPosition(pos *api.Position) *app.Position{
	return &app.Position{
		X: float64(pos.X),
		Y: float64(pos.Y),
	}
}

func toClock(clock app.IClock) *api.Clock{
	return &api.Clock{
		Timestamp: clock.GetTimestamp(),
	}
}

func toIClock(clock *api.Clock) app.IClock{
	return app.NewClock(clock.GetTimestamp())
}

func toAgent(agent app.IAgent) *api.Agent{
	return &api.Agent{
		Id: agent.GetID(),
		Name: agent.GetName(),
		Position: toPosition(agent.GetPosition()),
	}
}

func toAgents(agents []app.IAgent) []*api.Agent{
	apiAgents := []*api.Agent{}
	for _, ag := range agents{
		apiAgents = append(apiAgents, toAgent(ag))
	}
	return apiAgents
}

func toIAgent(agent *api.Agent, modelMap app.ModelMap) app.IAgent{ // TODO convert Agent by model
	name := agent.GetName()
	modelFunc := modelMap[name]
	iAgent := modelFunc(agent.GetId(),toIPosition(agent.GetPosition()))
	return iAgent
}

func toIAgents(agents []*api.Agent, modelMap app.ModelMap) []app.IAgent{
	appAgents := []app.IAgent{}
	for _, ag := range agents{
		appAgents = append(appAgents, toIAgent(ag, modelMap))
	}
	return appAgents
}

func toSimulator(sim ISimulator) *api.Simulator{
	neighbors := []*api.Neighbor{}
	for _, nei := range sim.GetNeighbors(){
		neighbors = append(neighbors, &api.Neighbor{
			Id: nei.ID,
			Port: uint64(nei.Port),
		})
	}
	return &api.Simulator{
		Id: sim.GetID(),
		Port: uint64(sim.GetPort()),
		Neighbors: neighbors,
	}
}


func toSimulators(sims []ISimulator) []*api.Simulator{
	apiSims := []*api.Simulator{}
	for _, ag := range sims{
		apiSims = append(apiSims, toSimulator(ag))
	}
	return apiSims
}


func toINeighbor(sim *api.Neighbor) *Neighbor{
	return &Neighbor{
		ID: sim.GetId(),
		Port: int(sim.GetPort()),
	}
}

func toINeighbors(neighbors []*api.Neighbor) []*Neighbor{
	neis := []*Neighbor{}
	for _, nei := range neighbors{
		neis = append(neis, toINeighbor(nei))
	}
	return neis
}

func toNeighbor(nei *Neighbor) *api.Neighbor{
	return &api.Neighbor{
		Id: nei.ID,
		Port: uint64(nei.Port),
	}
}


func toNeighbors(neis []*Neighbor) []*api.Neighbor{
	apiNeis := []*api.Neighbor{}
	for _, nei := range neis{
		apiNeis = append(apiNeis, toNeighbor(nei))
	}
	return apiNeis
}