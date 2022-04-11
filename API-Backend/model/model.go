package model

type MachineStates struct {
	Id             int `gorm:"primaryKey"`
	Running        bool
	AlertStop      bool
	FuncionalStop  bool
	StopByFoward   bool
	StopByBackward bool
	UnderSpeed     bool
}

var Registers []MachineStates
