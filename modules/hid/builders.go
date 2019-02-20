package hid

import (
	"github.com/bettercap/bettercap/network"
)

type FrameBuilder interface {
	BuildFrames([]*Command) error
}

var FrameBuilders = map[network.HIDType]FrameBuilder{
	network.HIDTypeLogitech: LogitechBuilder{},
	network.HIDTypeAmazon:   AmazonBuilder{},
}
