// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type InteractiveMarkerFeedback struct {
	Header          std_msgs.Header
	ClientId        msg.String
	MarkerName      msg.String
	ControlName     msg.String
	EventType       msg.Uint8
	Pose            geometry_msgs.Pose
	MenuEntryId     msg.Uint32
	MousePoint      geometry_msgs.Point
	MousePointValid msg.Bool
}
