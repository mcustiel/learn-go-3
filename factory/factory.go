package factory

//func dummy() {

//}

// "configLoader"
//"fileGameLoader"
//"game"
import "github.com/mcustiel/game/render"

import "github.com/mcustiel/game/logic"

//"simpleCollisionHandler"
//"simplePhysics"

//func GetGameLoader() {
//	return fileGameLoader.create(GetConfig()["gameConfigPath"])
//}

func CreateGameLogic() *logic.Game {
	return logic.CreateGameState()
}

//func GetCollisionHandler() {
//	return simpleCollisionHandler.create()
//}

func CreateRenderer() render.Renderer {
	return render.Create()
}

//func GetConfig() {
//	config := make(map[string]string)
//	config["gameConfigPath"] = "config"
//	return config
//}
