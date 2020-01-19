// Autogenerated with msg-import, do not edit.
package visualization_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Marker struct {
	Header                   std_msgs.Header
	Ns                       msg.String
	Id                       msg.Int32
	Type                     msg.Int32
	Action                   msg.Int32
	Pose                     geometry_msgs.Pose
	Scale                    geometry_msgs.Vector3
	Color                    std_msgs.ColorRGBA
	Lifetime                 msg.Duration
	FrameLocked              msg.Bool
	Points                   []geometry_msgs.Point
	Colors                   []std_msgs.ColorRGBA
	Text                     msg.String
	MeshResource             msg.String
	MeshUseEmbeddedMaterials msg.Bool
}
