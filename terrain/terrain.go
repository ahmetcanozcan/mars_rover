package terrain

import (
	"app/common"
)

//go:generate moq -out ../rover/terrain_moq_test.go . Terrain

var DefaultStartLocation = common.Location{X: 0, Y: 0}

type Terrain interface {
	GetName() string
	AddObject(Object)
	IsLocationAvailable(common.Location) bool
}

type terrainImpl struct {
	name    string
	width   int
	height  int
	objects []Object
}

func New(name string, width int, height int) Terrain {
	return &terrainImpl{name, width, height, make([]Object, 0)}
}

func (t *terrainImpl) GetName() string {
	return t.name
}

func (t *terrainImpl) AddObject(object Object) {
	l := object.Location()
	if !t.IsLocationAvailable(l) {
		object.SetLocation(DefaultStartLocation)
	}
	t.objects = append(t.objects, object)
	object.SetTerrain(t)
}

func (t *terrainImpl) IsLocationAvailable(l common.Location) bool {
	for _, object := range t.getObjects() {
		if object.Location().Equals(l) {
			return false
		}
	}
	return t.checkBoundary(l)
}

func (t *terrainImpl) checkBoundary(l common.Location) bool {
	return l.X >= 0 && l.X <= t.width && l.Y >= 0 && l.Y <= t.height
}

func (t *terrainImpl) getObjects() []Object {
	return t.objects
}
