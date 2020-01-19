// Autogenerated with msg-import, do not edit.
package sensor_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/std_msgs"
)

type Image struct {
	Header      std_msgs.Header
	Height      msg.Uint32
	Width       msg.Uint32
	Encoding    msg.String
	IsBigendian msg.Uint8
	Step        msg.Uint32
	Data        []msg.Uint8
}
