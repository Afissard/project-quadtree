package camera

import "github.com/Afissard/project-quadtree/configuration"

// Init met en place une caméra.
func (c *Camera) Init() {
	if configuration.Global.CameraMode == Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}
}
