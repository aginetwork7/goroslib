// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type ImageMarker struct {
	Header        std_msgs.Header
	Ns            msg.String
	Id            msg.Int32
	Type          msg.Int32
	Action        msg.Int32
	Position      geometry_msgs.Point
	Scale         msg.Float32
	OutlineColor  std_msgs.ColorRGBA
	Filled        msg.Uint8
	FillColor     std_msgs.ColorRGBA
	Lifetime      msg.Duration
	Points        []geometry_msgs.Point
	OutlineColors []std_msgs.ColorRGBA
}
