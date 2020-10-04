// Package visualization_msgs contains message definitions (autogenerated).
package visualization_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
	"time"
)

const (
	CIRCLE           uint8 = 0
	POLYGON          uint8 = 3
	ADD              uint8 = 0
	REMOVE           uint8 = 1
	INHERIT          uint8 = 0
	FIXED            uint8 = 1
	VIEW_FACING      uint8 = 2
	NONE             uint8 = 0
	MENU             uint8 = 1
	BUTTON           uint8 = 2
	MOVE_AXIS        uint8 = 3
	MOVE_PLANE       uint8 = 4
	ROTATE_AXIS      uint8 = 5
	MOVE_ROTATE      uint8 = 6
	MOVE_3D          uint8 = 7
	ROTATE_3D        uint8 = 8
	MOVE_ROTATE_3D   uint8 = 9
	KEEP_ALIVE       uint8 = 0
	POSE_UPDATE      uint8 = 1
	MENU_SELECT      uint8 = 2
	BUTTON_CLICK     uint8 = 3
	MOUSE_DOWN       uint8 = 4
	MOUSE_UP         uint8 = 5
	UPDATE           uint8 = 1
	ARROW            uint8 = 0
	CUBE             uint8 = 1
	SPHERE           uint8 = 2
	CYLINDER         uint8 = 3
	CUBE_LIST        uint8 = 6
	SPHERE_LIST      uint8 = 7
	TEXT_VIEW_FACING uint8 = 9
	MESH_RESOURCE    uint8 = 10
	TRIANGLE_LIST    uint8 = 11
	MODIFY           uint8 = 0
	DELETE           uint8 = 2
	DELETEALL        uint8 = 3
	FEEDBACK         uint8 = 0
	ROSRUN           uint8 = 1
	ROSLAUNCH        uint8 = 2
)

type ImageMarker struct {
	msg.Package     `ros:"visualization_msgs"`
	msg.Definitions `ros:"uint8 CIRCLE=0,uint8 LINE_STRIP=1,uint8 LINE_LIST=2,uint8 POLYGON=3,uint8 POINTS=4,uint8 ADD=0,uint8 REMOVE=1"`
	Header          std_msgs.Header
	Ns              string
	Id              int32
	Type            int32
	Action          int32
	Position        geometry_msgs.Point
	Scale           float32
	OutlineColor    std_msgs.ColorRGBA
	Filled          uint8
	FillColor       std_msgs.ColorRGBA
	Lifetime        time.Duration
	Points          []geometry_msgs.Point
	OutlineColors   []std_msgs.ColorRGBA
}

type InteractiveMarker struct {
	msg.Package `ros:"visualization_msgs"`
	Header      std_msgs.Header
	Pose        geometry_msgs.Pose
	Name        string
	Description string
	Scale       float32
	MenuEntries []MenuEntry
	Controls    []InteractiveMarkerControl
}

type InteractiveMarkerControl struct {
	msg.Package                  `ros:"visualization_msgs"`
	msg.Definitions              `ros:"uint8 INHERIT=0,uint8 FIXED=1,uint8 VIEW_FACING=2,uint8 NONE=0,uint8 MENU=1,uint8 BUTTON=2,uint8 MOVE_AXIS=3,uint8 MOVE_PLANE=4,uint8 ROTATE_AXIS=5,uint8 MOVE_ROTATE=6,uint8 MOVE_3D=7,uint8 ROTATE_3D=8,uint8 MOVE_ROTATE_3D=9"`
	Name                         string
	Orientation                  geometry_msgs.Quaternion
	OrientationMode              uint8
	InteractionMode              uint8
	AlwaysVisible                bool
	Markers                      []Marker
	IndependentMarkerOrientation bool
	Description                  string
}

type InteractiveMarkerFeedback struct {
	msg.Package     `ros:"visualization_msgs"`
	msg.Definitions `ros:"uint8 KEEP_ALIVE=0,uint8 POSE_UPDATE=1,uint8 MENU_SELECT=2,uint8 BUTTON_CLICK=3,uint8 MOUSE_DOWN=4,uint8 MOUSE_UP=5"`
	Header          std_msgs.Header
	ClientId        string
	MarkerName      string
	ControlName     string
	EventType       uint8
	Pose            geometry_msgs.Pose
	MenuEntryId     uint32
	MousePoint      geometry_msgs.Point
	MousePointValid bool
}

type InteractiveMarkerInit struct {
	msg.Package `ros:"visualization_msgs"`
	ServerId    string
	SeqNum      uint64
	Markers     []InteractiveMarker
}

type InteractiveMarkerPose struct {
	msg.Package `ros:"visualization_msgs"`
	Header      std_msgs.Header
	Pose        geometry_msgs.Pose
	Name        string
}

type InteractiveMarkerUpdate struct {
	msg.Package     `ros:"visualization_msgs"`
	msg.Definitions `ros:"uint8 KEEP_ALIVE=0,uint8 UPDATE=1"`
	ServerId        string
	SeqNum          uint64
	Type            uint8
	Markers         []InteractiveMarker
	Poses           []InteractiveMarkerPose
	Erases          []string
}

type Marker struct {
	msg.Package              `ros:"visualization_msgs"`
	msg.Definitions          `ros:"uint8 ARROW=0,uint8 CUBE=1,uint8 SPHERE=2,uint8 CYLINDER=3,uint8 LINE_STRIP=4,uint8 LINE_LIST=5,uint8 CUBE_LIST=6,uint8 SPHERE_LIST=7,uint8 POINTS=8,uint8 TEXT_VIEW_FACING=9,uint8 MESH_RESOURCE=10,uint8 TRIANGLE_LIST=11,uint8 ADD=0,uint8 MODIFY=0,uint8 DELETE=2,uint8 DELETEALL=3"`
	Header                   std_msgs.Header
	Ns                       string
	Id                       int32
	Type                     int32
	Action                   int32
	Pose                     geometry_msgs.Pose
	Scale                    geometry_msgs.Vector3
	Color                    std_msgs.ColorRGBA
	Lifetime                 time.Duration
	FrameLocked              bool
	Points                   []geometry_msgs.Point
	Colors                   []std_msgs.ColorRGBA
	Text                     string
	MeshResource             string
	MeshUseEmbeddedMaterials bool
}

type MarkerArray struct {
	msg.Package `ros:"visualization_msgs"`
	Markers     []Marker
}

type MenuEntry struct {
	msg.Package     `ros:"visualization_msgs"`
	msg.Definitions `ros:"uint8 FEEDBACK=0,uint8 ROSRUN=1,uint8 ROSLAUNCH=2"`
	Id              uint32
	ParentId        uint32
	Title           string
	Command         string
	CommandType     uint8
}
