package main

type placeholder [5]string

var a = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}
var l = placeholder{
	"█  ",
	"█  ",
	"█  ",
	"█  ",
	"███",
}

var r = placeholder{
	"██ ",
	"█ █",
	"██ ",
	"█ █",
	"█ █",
}

var m = placeholder{
	"█ █",
	"███",
	"█ █",
	"█ █",
	"█ █",
}

var exclamationPoint = placeholder{
	" █ ",
	" █ ",
	" █ ",
	"   ",
	" █ ",
}

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"██ ",
	"  █",
	" █ ",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	" █ ",
	" █ ",
	" █ ",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"  █",
}

var separator = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var blank = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var period = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	" █ ",
}
var nums = [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine, separator}
var alarm = [...]placeholder{blank, blank, a, l, a, r, m, exclamationPoint, blank, blank}
