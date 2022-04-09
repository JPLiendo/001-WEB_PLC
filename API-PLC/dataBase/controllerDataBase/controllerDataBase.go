package controllerDatabase

import "gorm.io/gorm"

// Type states.
type states struct {
	Running        bool
	AlertStop      bool
	FuncionalStop  bool
	StopByFoward   bool
	StopByBackward bool
	UnderVelocity  bool
}

//Type BlowerStates.
type BlowerStates struct {
	gorm.Model
	States states
}

//Type FillerStates.
type FillerStates struct {
	gorm.Model
	States states
}

//Type LabellerStates.
type LabellerStates struct {
	gorm.Model
	States states
}

//Type PackerStates.
type PackerStates struct {
	gorm.Model
	States states
}

//Type PalletizerStates.
type PalletizerStates struct {
	gorm.Model
	States states
}

//Interface.
type MachineState interface {
	createRegister(*gorm.DB)
}

//Polymorphism method CreateRegister for type BlowerStates.
func (f *BlowerStates) createRegister(db *gorm.DB) *gorm.DB {
	newRegister := db.Create(f)
	return newRegister

}

//Polymorphism method CreateRegister for type FillerStates.
func (f *FillerStates) createRegister(db *gorm.DB) *gorm.DB {
	newRegister := db.Create(f)
	return newRegister

}

//Polymorphism method CreateRegister for type LabellerStates.
func (f *LabellerStates) createRegister(db *gorm.DB) *gorm.DB {
	newRegister := db.Create(f)
	return newRegister

}

//Polymorphism method CreateRegister for type PackerStates.
func (f *PackerStates) createRegister(db *gorm.DB) *gorm.DB {
	newRegister := db.Create(f)
	return newRegister

}

//Polymorphism method CreateRegister for type PalletizerStates.
func (f *PalletizerStates) createRegister(db *gorm.DB) *gorm.DB {
	newRegister := db.Create(f)
	return newRegister

}

// CreateMachineRegister is the exported function to be used to write a new machine's register.
func CreateMachineRegister(machineState MachineState, db *gorm.DB) {
	machineState.createRegister(db)
}
