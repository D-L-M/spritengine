package spritengine

// Direction constants
const (
	DirLeft       = -1
	DirRight      = 1
	DirStationary = 0
)

// Event constants
const (
	EventFloorCollision = 0x00000001
	EventDropOffLevel   = 0x00000010
	EventFreefall       = 0x00000011
)

// Collision edges
const (
	EdgeTop    = "top"
	EdgeBottom = "bottom"
	EdgeLeft   = "left"
	EdgeRight  = "right"
	EdgeNone   = "none"
)
