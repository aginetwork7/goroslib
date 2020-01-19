// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type CameraInfo struct {
	Header          std_msgs.Header
	Height          msg.Uint32
	Width           msg.Uint32
	DistortionModel msg.String
	D               []msg.Float64
	K               [9]msg.Float64
	R               [9]msg.Float64
	P               [12]msg.Float64
	BinningX        msg.Uint32
	BinningY        msg.Uint32
	Roi             RegionOfInterest
}
