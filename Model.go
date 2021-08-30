package main


type Info struct {
	ProgramName string
	Pid         int
	CmdLine     string
	User        string
}

type Points struct {
	Info      Info
	MetricMap map[string]float64
}

type Point struct {
	Info       Info
	MetricName string
	Value      float64
}