package main

import (
	"fmt"
	"time"
)

type IModule interface {
	Init() bool
	MainLoop()
	Reload()
	Final() bool
}

const (
	ModuleId_PlayerModle  = 1
)

type ModuleMgr struct {
	Mgr map[int32]IModule
}

func NewModuleMgr() *ModuleMgr {
	return &ModuleMgr{
		Mgr: make(map[int32]IModule),
	}
}

func (mgr *ModuleMgr) Init() bool {
	mgr.Mgr[ModuleId_PlayerModle] = NewPlayerModule()
	return true
}

func (mgr *ModuleMgr) MainLoop() {
	for modId, mod := range mgr.Mgr {
		fmt.Println("modId ", modId)
		mod.MainLoop()
	}
}

func (mgr *ModuleMgr) Reload() {

}

func (mgr *ModuleMgr) Final() bool {
	return true
}


type PlayerModule struct {
	Derived IModule
}

func (mgr *PlayerModule) Init() bool {
	return true
}

func (mgr *PlayerModule) MainLoop() {
	fmt.Println("PlayerModule is running")
}

func (mgr *PlayerModule) Reload() {

}

func (mgr *PlayerModule) Final() bool {
	return true
}

func NewPlayerModule() *PlayerModule {
	module := &PlayerModule{}
	module.Derived = module
	return module
}

func main() {
	modMgr := NewModuleMgr()
	modMgr.Init()

	time.Sleep(10 * time.Second)

	modMgr.MainLoop()

	time.Sleep(10 * time.Second)

	modMgr.Final()

	fmt.Println("final")
}
