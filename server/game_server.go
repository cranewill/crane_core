package server

type GameServer struct {
	Uid    string
	Role   int
	Status int
}

func NewGameServer(uid string, role int) *GameServer {
	return &GameServer{
		Uid:  uid,
		Role: role,
	}
}

func (gs *GameServer) Start() {

}

func (gs *GameServer) OnStart() {

}

func (gs *GameServer) Stop() {

}

func (gs *GameServer) OnStop() {

}
