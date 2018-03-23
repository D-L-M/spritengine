package spritengine

// Direction constants
const (
	DirLeft       = -1
	DirRight      = 1
	DirStationary = 0
)

// Event constants
const (
	EventFloorCollision = 0
	EventDropOffLevel   = 1
	EventFreefall       = 2
)

// Collision edges
const (
	EdgeTop         = "top"
	EdgeTopLeft     = "top_left"
	EdgeTopRight    = "top_right"
	EdgeBottom      = "bottom"
	EdgeBottomLeft  = "bottom_left"
	EdgeBottomRight = "bottom_right"
	EdgeLeft        = "left"
	EdgeRight       = "right"
	EdgeNone        = "none"
)
