package terrain

import "app/common"

type Object interface {
	Location() common.Location
	SetLocation(common.Location)
	SetTerrain(terrain Terrain)
}
