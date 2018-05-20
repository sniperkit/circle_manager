package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSacnLineForRouter(t *testing.T) {
	key, value := getMetaKeyAndValueOfRouter("// @APIVersion 1.2.100")
	assert.Equal(t, "APIVersion", key)
	assert.Equal(t, "1.2.100", value)

	// scanLineForRouter(flagRead, cs, &currentWhere, "// @Title TestApp!!!")
	// scanLineForRouter(flagRead, cs, &currentWhere, "// @Description TestApp입니다")
	// scanLineForRouter(flagRead, cs, &currentWhere, "// @Contact test@adminmail.com")
	// scanLineForRouter(flagRead, cs, &currentWhere, "// @TermsOfServiceUrl http://circle-test.com")
	// scanLineForRouter(flagRead, cs, &currentWhere, "// @License PRIVATE")
	// scanLineForRouter(flagRead, cs, &currentWhere, "// @SecurityDefinition xxxxxxx")

	// assert.Equal(t, "1.2.100", cs.AppVersion)
	// assert.True(t, flagRead.RouterReadedAppVersion)
	// assert.Equal(t, "TestApp!!!", cs.AppTitle)
	// assert.True(t, flagRead.RouterReadedAppTitle)
	// assert.Equal(t, "TestApp입니다", cs.AppDescription)
	// assert.True(t, flagRead.RouterReadedAppDescription)
	// assert.Equal(t, "test@adminmail.com", cs.AppContact)
	// assert.True(t, flagRead.RouterReadedAppContact)
	// assert.Equal(t, "http://circle-test.com", cs.AppTermsOfServiceUrl)
	// assert.True(t, flagRead.RouterReadedAppTermsOfServiceUrl)
	// assert.Equal(t, "PRIVATE", cs.AppLicense)
	// assert.True(t, flagRead.RouterReadedAppLicense)
	// assert.Equal(t, "xxxxxxx", cs.AppSecurityDefinition)
	// assert.True(t, flagRead.RouterReadedAppSecurityDefinition)

	// currentWhere = "system"
	// scanLineForRouter(flagRead, cs, &currentWhere, "     &modules.Test1Controller{},")
	// scanLineForRouter(flagRead, cs, &currentWhere, "          &modules.Test2Controller{},")

	// assert.Equal(t, len(cs.Units), 2)

	// // assert.Equal(t, cs.Units[0].Name, "Test1")
	// // assert.True(t, cs.Units[0].IsSystem)
	// // assert.False(t, cs.Units[0].IsManual)

	// // assert.Equal(t, cs.Units[1].Name, "Test2")
	// // assert.True(t, cs.Units[1].IsSystem)
	// // assert.False(t, cs.Units[1].IsManual)

	// // currentWhere = "manual"
	// // scanLineForRouter(flagRead, cs, &currentWhere, "     &modules.Test3Controller{},")
	// // scanLineForRouter(flagRead, cs, &currentWhere, "          &modules.Test4Controller{},")

	// // assert.Equal(t, cs.Units[2].Name, "Test3")
	// // assert.False(t, cs.Units[2].IsSystem)
	// // assert.True(t, cs.Units[2].IsManual)

	// // assert.Equal(t, cs.Units[3].Name, "Test4")
	// // assert.False(t, cs.Units[3].IsSystem)
	// // assert.True(t, cs.Units[3].IsManual)
}

// func TestSacnLineForAdmin(t *testing.T) {
// 	cs := &modules.CircleSet{}

// 	currentWhere := "system"
// 	scanLineForAdmin(flagRead, cs, &currentWhere, `"addResourceAndMenu(&models.GithubCommit{}, "Commits", "이벤트관리", anyoneAllow, -1)"`)

// 	assert.Equal(t, len(cs.Units), 1)
// 	assert.Equal(t, cs.Units[0].Name, "GithubCommit")
// 	assert.Equal(t, cs.Units[0].MenuName, "Commits")
// 	assert.Equal(t, cs.Units[0].MenuGroup, "이벤트관리")
// 	assert.True(t, cs.Units[0].IsSystem)
// 	assert.False(t, cs.Units[0].IsManual)
// }

// func TestSacnSourceForModel(t *testing.T) {
// 	cs, err := scanSource("models", ".example/models")
// 	assert.Nil(t, err)
// 	assert.NotZero(t, len(cs.Units))
// }

// func TestSacnSourceForController(t *testing.T) {
// 	_, err := scanSource("controllers", ".example/controllers")
// 	assert.Nil(t, err)
// }

// func TestSacnSourceForRequests(t *testing.T) {
// 	cs, err := scanSource("requests", ".example/requests")
// 	assert.Nil(t, err)
// 	assert.NotZero(t, len(cs.Units))
// 	for _, unit := range cs.Units {
// 		assert.NotEmpty(t, unit.Name)
// 		assert.True(t, unit.EnableRequestSource)
// 		assert.False(t, unit.EnableControllerSource)
// 		assert.False(t, unit.EnableAdminSource)
// 		assert.False(t, unit.EnableModelSource)
// 		assert.False(t, unit.EnableResponseSource)
// 	}
// }

// func TestSacnSourceForResponses(t *testing.T) {
// 	cs, err := scanSource("responses", ".example/responses")
// 	assert.Nil(t, err)
// 	assert.NotZero(t, len(cs.Units))
// 	for _, unit := range cs.Units {
// 		assert.NotEmpty(t, unit.Name)
// 		assert.False(t, unit.EnableRequestSource)
// 		assert.False(t, unit.EnableControllerSource)
// 		assert.False(t, unit.EnableAdminSource)
// 		assert.False(t, unit.EnableModelSource)
// 		assert.True(t, unit.EnableResponseSource)
// 	}
// }

// func TestSourceMerge(t *testing.T) {
// 	envs = &Envs{
// 		AppDir: ".example",
// 	}
// 	cm := &CircleManager{}

// 	_, err := cm.ImportCircle()
// 	assert.Nil(t, err)

// 	//spew.Dump(cs)
// }

// func TestSacnLineForModel(t *testing.T) {
// 	flagRead := &FlagRead{}
// 	cu := &modules.CircleUnit{
// 		Name: "Test1",
// 	}

// 	currentWhere := ""
// 	scanLineForModel(flagRead, cu, &currentWhere, `type Test1 struct {`)
// 	assert.Equal(t, currentWhere, "in_model")

// 	scanLineForModel(flagRead, cu, &currentWhere, `}`)
// 	assert.Equal(t, currentWhere, "end_model")

// 	currentWhere = "in_model"
// 	scanLineForModel(flagRead, cu, &currentWhere, "ID      uint       `description:\"등록일\"` ")
// 	assert.Equal(t, len(cu.Properties), 1)
// 	assert.Equal(t, cu.Properties[0].Name, "ID")
// 	assert.Equal(t, cu.Properties[0].Description, "등록일")
// 	assert.Equal(t, cu.Properties[0].Type, "uint")
// 	assert.True(t, cu.Properties[0].IsSystem)

// 	scanLineForModel(flagRead, cu, &currentWhere, "Prop1      string       `description:\"가나다라마바사 !~!! 1  !!\"` ")
// 	assert.Equal(t, len(cu.Properties), 2)
// 	assert.Equal(t, cu.Properties[1].Name, "Prop1")
// 	assert.Equal(t, cu.Properties[1].Description, "가나다라마바사 !~!! 1  !!")
// 	assert.Equal(t, cu.Properties[1].Type, "string")
// 	assert.False(t, cu.Properties[1].IsSystem)
// }
