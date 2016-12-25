package palette

// Though the import name is kelly22, access the vars below as palette.Kelly22
// and not kelly22.Kelly22.

import "image/color"

// kelly_colors = [ '2B3D26']

//Kelly22 is Kelly's 22 colours of maximum contrast.
var Kelly22 = []color.Color{
	color.RGBA{0xf2, 0xf3, 0xf4, 0xff}, // white
	color.RGBA{0x22, 0x22, 0x22, 0xff}, // black
	color.RGBA{0xf3, 0xc3, 0x00, 0xff}, // yellow
	color.RGBA{0x87, 0x56, 0x92, 0xff}, // purple
	color.RGBA{0xf3, 0x84, 0x00, 0xff}, // orange
	color.RGBA{0xa1, 0xca, 0xf1, 0xff}, // light blue
	color.RGBA{0xbe, 0x00, 0x32, 0xff}, // red
	color.RGBA{0xc2, 0xb2, 0x80, 0xff}, // buff
	color.RGBA{0x84, 0x84, 0x82, 0xff}, // grey
	color.RGBA{0x00, 0x88, 0x56, 0xff}, // green
	color.RGBA{0xe6, 0x8f, 0xac, 0xff}, // purplish pink
	color.RGBA{0x00, 0x67, 0xa5, 0xff}, // blue
	color.RGBA{0xf9, 0x93, 0x79, 0xff}, // yellowish pink
	color.RGBA{0x60, 0x4e, 0x97, 0xff}, // violet
	color.RGBA{0xf6, 0xa6, 0x00, 0xff}, // orange yellow
	color.RGBA{0xb3, 0x44, 0x6c, 0xff}, // purplish red
	color.RGBA{0xdc, 0xd3, 0x00, 0xff}, // greenish yellow
	color.RGBA{0x88, 0x2d, 0x17, 0xff}, // reddish brown
	color.RGBA{0x8d, 0xb6, 0x00, 0xff}, // yellow green
	color.RGBA{0x65, 0x45, 0x22, 0xff}, // yellowish brown
	color.RGBA{0xe2, 0x58, 0x22, 0xff}, // reddish orange
	color.RGBA{0x2b, 0x3d, 0x26, 0xff}, // olive green
}

// KellyLightBG is the same as Kelly22, the first color --
// used for the background is white.
var KellyLightBG = Kelly22

// KellyDarkBG has the same colours as Kelly22 but the first color --
// used for the background is black.
var KellyDarkBG = []color.Color{
	color.RGBA{0x22, 0x22, 0x22, 0xff}, // black
	color.RGBA{0xf2, 0xf3, 0xf4, 0xff}, // white
	color.RGBA{0xf3, 0xc3, 0x00, 0xff}, // yellow
	color.RGBA{0x87, 0x56, 0x92, 0xff}, // purple
	color.RGBA{0xf3, 0x84, 0x00, 0xff}, // orange
	color.RGBA{0xa1, 0xca, 0xf1, 0xff}, // light blue
	color.RGBA{0xbe, 0x00, 0x32, 0xff}, // red
	color.RGBA{0xc2, 0xb2, 0x80, 0xff}, // buff
	color.RGBA{0x84, 0x84, 0x82, 0xff}, // grey
	color.RGBA{0x00, 0x88, 0x56, 0xff}, // green
	color.RGBA{0xe6, 0x8f, 0xac, 0xff}, // purplish pink
	color.RGBA{0x00, 0x67, 0xa5, 0xff}, // blue
	color.RGBA{0xf9, 0x93, 0x79, 0xff}, // yellowish pink
	color.RGBA{0x60, 0x4e, 0x97, 0xff}, // violet
	color.RGBA{0xf6, 0xa6, 0x00, 0xff}, // orange yellow
	color.RGBA{0xb3, 0x44, 0x6c, 0xff}, // purplish red
	color.RGBA{0xdc, 0xd3, 0x00, 0xff}, // greenish yellow
	color.RGBA{0x88, 0x2d, 0x17, 0xff}, // reddish brown
	color.RGBA{0x8d, 0xb6, 0x00, 0xff}, // yellow green
	color.RGBA{0x65, 0x45, 0x22, 0xff}, // yellowish brown
	color.RGBA{0xe2, 0x58, 0x22, 0xff}, // reddish orange
	color.RGBA{0x2b, 0x3d, 0x26, 0xff}, // olive green
}