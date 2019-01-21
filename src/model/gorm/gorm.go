package gormmodel

import (
	"github.com/LyricTian/gin-admin/src/model"
	"github.com/LyricTian/gin-admin/src/model/gorm/common"
	"github.com/LyricTian/gin-admin/src/model/gorm/demo"
	"github.com/LyricTian/gin-admin/src/model/gorm/menu"
	"github.com/LyricTian/gin-admin/src/model/gorm/role"
	"github.com/LyricTian/gin-admin/src/service/gormplus"
	"github.com/facebookgo/inject"
)

// Init 初始化gorm存储
func Init(g *inject.Graph, db *gormplus.DB) {
	g.Provide(&inject.Object{Value: model.ITrans(gormcommon.NewTrans(db)), Name: "ITrans"})
	g.Provide(&inject.Object{Value: model.IDemo(gormdemo.NewModel(db)), Name: "IDemo"})
	g.Provide(&inject.Object{Value: model.IMenu(gormmenu.NewModel(db)), Name: "IMenu"})
	g.Provide(&inject.Object{Value: model.IRole(gormrole.NewModel(db)), Name: "IRole"})
}