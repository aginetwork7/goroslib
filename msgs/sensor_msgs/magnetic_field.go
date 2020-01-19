// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type MagneticField struct {
	Header                  std_msgs.Header
	MagneticField           geometry_msgs.Vector3
	MagneticFieldCovariance [9]msg.Float64
}
