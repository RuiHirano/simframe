package engine

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/RuiHirano/simframe/api"
	"github.com/RuiHirano/simframe/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type SimulatorAPI struct {
	Address string
	Client api.SimulatorServiceClient
}

func NewSimulatorAPI(address string) *SimulatorAPI {
	sa := &SimulatorAPI{
		Address: address,
	}
	sa.SetUp()
	return sa
}

func (sa *SimulatorAPI) SetUp() {
	conn, err := grpc.Dial(sa.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	sa.Client = api.NewSimulatorServiceClient(conn)
}

func (sa *SimulatorAPI) RunSimulator() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := sa.Client.RunSimulator(ctx, &api.RunSimulatorRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get status: %s", r.GetStatus())
}

func (sa *SimulatorAPI) GetNeighborAgents(agents []app.IAgent) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := sa.Client.GetNeighborAgents(ctx, &api.GetNeighborAgentsRequest{
		Agents: []*api.Agent{},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get status: %s", r.GetStatus())
}

type SimulatorService struct{
	Port int
	RunSimulatorHandler func()
	GetNeighborAgentsHandler func()
	api.UnimplementedSimulatorServiceServer
}

func NewSimulatorService(port int, rsh func(), gnah func()) *SimulatorService{
   ms := &SimulatorService{
	   Port: port,
	   RunSimulatorHandler: rsh,
	   GetNeighborAgentsHandler: rsh,
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

func (es *SimulatorService) RunSimulator(ctx context.Context, request *api.RunSimulatorRequest) (*api.RunSimulatorResponse, error) {
   fmt.Printf("getRequest %v\n", request)

   es.RunSimulatorHandler()

   response := &api.RunSimulatorResponse{
	   Status: api.Status_OK,
   }
   return response, nil
}

func (es *SimulatorService) GetNeighborAgents(ctx context.Context, request *api.GetNeighborAgentsRequest) (*api.GetNeighborAgentsResponse, error) {
	fmt.Printf("getRequest %v\n", request)
 
	es.GetNeighborAgentsHandler()
 
	response := &api.GetNeighborAgentsResponse{
		Status: api.Status_OK,
		Agents: []*api.Agent{},
	}
	return response, nil
 }


type EngineAPI struct {
	Address string
	Client api.EngineServiceClient
}

func NewEngineAPI(address string) *EngineAPI {
	ea := &EngineAPI{
		Address: address,
	}
	ea.SetUp()
	return ea
}

func (ea *EngineAPI) SetUp() {
	conn, err := grpc.Dial(ea.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ea.Client = api.NewEngineServiceClient(conn)
}

func (ea *EngineAPI) RegisterSimulator(sim ISimulator) (app.IArea, app.IClock, []app.IAgent){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := ea.Client.RegisterSimulator(ctx, &api.RegisterSimulatorRequest{
		Simulator: &api.Simulator{
			Id: sim.GetID(),
			ServerAddress: sim.GetServerAddress(),
			Neighbors: toSimulators(sim.GetNeighbors()),
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get status: %s", r.GetStatus())
	return toIArea(r.GetArea()), toIClock(r.GetClock()), toIAgents(r.GetAgents())
}


type EngineService struct{
	RegisterSimulatorHandler func() (app.IArea, app.IClock, []app.IAgent)
	Port int
	api.UnimplementedEngineServiceServer
}

func NewEngineService(port int, rsh func() (app.IArea, app.IClock, []app.IAgent)) *EngineService{
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

   area, clock, agents := es.RegisterSimulatorHandler()

   response := &api.RegisterSimulatorResponse{
	   Status: api.Status_OK,
	   Area: toArea(area),
	   Clock: toClock(clock),
	   Agents: toAgents(agents),
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

func toIAgent(agent *api.Agent) *app.Agent{ // TODO to IAgent
	return app.NewAgent(agent.GetId(), agent.GetName(), toIPosition(agent.GetPosition()))
}

func toIAgents(agents []*api.Agent) []app.IAgent{
	appAgents := []app.IAgent{}
	for _, ag := range agents{
		appAgents = append(appAgents, toIAgent(ag))
	}
	return appAgents
}

func toSimulator(sim ISimulator) *api.Simulator{
	return &api.Simulator{
		Id: sim.GetID(),
		ServerAddress: sim.GetServerAddress(),
		Neighbors: toSimulators(sim.GetNeighbors()),
	}
}

func toSimulators(sims []ISimulator) []*api.Simulator{
	apiSims := []*api.Simulator{}
	for _, ag := range sims{
		apiSims = append(apiSims, toSimulator(ag))
	}
	return apiSims
}

