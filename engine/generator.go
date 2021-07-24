package engine

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

type Resource struct {
	ApiVersion string   `yaml:"apiVersion,omitempty"`
	Kind       string   `yaml:"kind,omitempty"`
	Metadata   Metadata `yaml:"metadata,omitempty"`
	Spec       Spec     `yaml:"spec,omitempty"`
}

type Spec struct {
	Containers []Container `yaml:"containers,omitempty"`
	Selector   Selector    `yaml:"selector,omitempty"`
	Ports      []Port      `yaml:"ports,omitempty"`
	Type       string      `yaml:"type,omitempty"`
}

type Container struct {
	Name            string `yaml:"name,omitempty"`
	Image           string `yaml:"image,omitempty"`
	ImagePullPolicy string `yaml:"imagePullPolicy,omitempty"`
	Stdin           bool   `yaml:"stdin,omitempty"`
	Tty             bool   `yaml:"tty,omitempty"`
	Env             []Env  `yaml:"env,omitempty"`
	Ports           []Port `yaml:"ports,omitempty"`
	Command         []string `yaml:"command,omitempty"`
}

type Env struct {
	Name  string `yaml:"name,omitempty"`
	Value string `yaml:"value,omitempty"`
}

type Selector struct {
	App         string `yaml:"app,omitempty"`
	MatchLabels Label  `yaml:"matchLabels,omitempty"`
}

type Port struct {
	Name          string `yaml:"name,omitempty"`
	Port          int    `yaml:"port,omitempty"`
	TargetPort    int    `yaml:"targetPort,omitempty"`
	NodePort      int    `yaml:"nodePort,omitempty"`
	ContainerPort int    `yaml:"containerPort,omitempty"`
}

type Metadata struct {
	Name   string `yaml:"name,omitempty"`
	Labels Label  `yaml:"labels,omitempty"`
}

type Label struct {
	App string `yaml:"app,omitempty"`
}

type Area struct {
	Id          int
	Control     []Coord
	Duplicate   []Coord
	NeighborIds []int // neighbor area
}

type Coord struct {
	Latitude  float64
	Longitude float64
}

type Option struct {
	FileName        string
	AreaCoords      []Coord
	DevideSquareNum int
	DuplicateRate   float64
}

func GenerateUUID()string{
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)	
	}
	uu := u.String()
	return uu
}

type IResourceGenerator interface {
	Apply()
}

type ResourceGenerator struct {
}

func NewResourceGenerator() *ResourceGenerator {

	rc := &ResourceGenerator{
	}

	return rc
}

func (bd *ResourceGenerator) Apply(id string, port int) error{
	fmt.Print("GenerateResource!")
	
	rsrcs := []Resource{
		bd.NewSimulatorService(id, port),
		bd.NewSimulator(id, port),
	}

	// create yaml
	fileName := "simulator" + id + ".yaml"
	for _, rsrc := range rsrcs {
		err := bd.WriteOnFile(fileName, rsrc)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	
	// apply yaml
	cmd := exec.Command("kubectl", "apply", "-f", fileName)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Command Start Error. %v\n", err)
		return err
	}

	// delete yaml
	if err := os.Remove(fileName); err != nil {
        fmt.Println(err)
		return err
    }
	return nil
}


func (bd *ResourceGenerator) AreaDivider(areaCoords []Coord, divideSquareNum int, duplicateRate float64) []Area {

	//neighbors := [][]int{}
	areas := []Area{}

	maxLat, maxLon, minLat, minLon := GetCoordRange(areaCoords)
	for i := 0; i < divideSquareNum; i++ { // 横方向
		leftlon := (maxLon-minLon)*float64(i)/float64(divideSquareNum) + minLon
		rightlon := (maxLon-minLon)*(float64(i)+1)/float64(divideSquareNum) + minLon

		for k := 0; k < divideSquareNum; k++ { // 縦方向
			bottomlat := (maxLat-minLat)*float64(k)/float64(divideSquareNum) + minLat
			toplat := (maxLat-minLat)*(float64(k)+1)/float64(divideSquareNum) + minLat
			id, _ := strconv.Atoi(strconv.Itoa(i+1) + strconv.Itoa(k+1))
			area := Area{
				Id: id,
				Control: []Coord{
					{Longitude: leftlon, Latitude: toplat},
					{Longitude: leftlon, Latitude: bottomlat},
					{Longitude: rightlon, Latitude: bottomlat},
					{Longitude: rightlon, Latitude: toplat},
				},
				Duplicate: []Coord{
					{Longitude: leftlon - (rightlon-leftlon)*duplicateRate, Latitude: toplat + (toplat-bottomlat)*duplicateRate},
					{Longitude: leftlon - (rightlon-leftlon)*duplicateRate, Latitude: bottomlat - (toplat-bottomlat)*duplicateRate},
					{Longitude: rightlon + (rightlon-leftlon)*duplicateRate, Latitude: bottomlat - (toplat-bottomlat)*duplicateRate},
					{Longitude: rightlon + (rightlon-leftlon)*duplicateRate, Latitude: toplat + (toplat-bottomlat)*duplicateRate},
				},
			}

			// add neighbors
			if i+1+1 <= divideSquareNum {
				id2, _ := strconv.Atoi(strconv.Itoa(i+1+1) + strconv.Itoa(k+1))
				area.NeighborIds = append(area.NeighborIds, id2)
			}
			if k+1+1 <= divideSquareNum {
				id3, _ := strconv.Atoi(strconv.Itoa(i+1) + strconv.Itoa(k+1+1))
				area.NeighborIds = append(area.NeighborIds, id3)
			}

			areas = append(areas, area)

		}
	}

	return areas

}

func GetCoordRange(coords []Coord) (float64, float64, float64, float64) {
	maxLon, maxLat := math.Inf(-1), math.Inf(-1)
	minLon, minLat := math.Inf(0), math.Inf(0)
	for _, coord := range coords {
		if coord.Latitude > maxLat {
			maxLat = coord.Latitude
		}
		if coord.Longitude > maxLon {
			maxLon = coord.Longitude
		}
		if coord.Latitude < minLat {
			minLat = coord.Latitude
		}
		if coord.Longitude < minLon {
			minLon = coord.Longitude
		}
	}
	return maxLat, maxLon, minLat, minLon
}


func (bd *ResourceGenerator) WriteOnFile(fileName string, data interface{}) error {
	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, string(buf))   
	fmt.Fprintln(file, string("---")) 
	return nil
}


// engine
func (bd *ResourceGenerator) NewSimulatorService(id string, port int) Resource {
	service := Resource{
		ApiVersion: "v1",
		Kind:       "Service",
		Metadata:   Metadata{Name: "simulator"+id},
		Spec: Spec{
			Selector: Selector{App: "simulator"+id},
			Ports: []Port{
				{
					Name:       "simulator"+id,
					Port:       port,
					TargetPort: port,
				},
			},
		},
	}
	return service
}

func (bd *ResourceGenerator) NewSimulator(id string, port int) Resource {
	master := Resource{
		ApiVersion: "v1",
		Kind:       "Pod",
		Metadata: Metadata{
			Name:   "simulator"+id,
			Labels: Label{App: "simulator"+id},
		},
		Spec: Spec{
			Containers: []Container{
				{
					Name:            "simulator"+id,
					Image:           fmt.Sprintf("simframe/%s:%s", "sample", "1.0.0"),
					ImagePullPolicy: "Never",
					Env: []Env{
					},
					Ports: []Port{{ContainerPort: port}},
					Command: []string{"go", "run", "main.go", "run", "simulator"},
				},
			},
		},
	}
	return master
}