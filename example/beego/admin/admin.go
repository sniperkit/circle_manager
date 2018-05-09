package admin

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"github.com/jungju/circle_manager/example/beegoapp/admin/bindatafs"
	"github.com/jungju/qor_admin_auth"
	"github.com/rs/cors"
)

func SetAdmin(m *qor_admin_auth.QorAdminManager, db *gorm.DB) error {
	bd := bindatafs.AssetFS.NameSpace("admin")
	m.QorPage.SetAssetFS(bd)
	if m.Auth != nil {
		m.Auth.Render.SetAssetFS(bd)
	}

	setViews(m.QorPage)

	mux := m.DefaultServeMux()
	handler := cors.AllowAll().Handler(mux)
	beego.Handler(fmt.Sprintf("/%s/*", m.Config.AuthURL), handler)
	beego.Handler(fmt.Sprintf("/%s/*", m.Config.AdminURL), handler)

	return nil
}
