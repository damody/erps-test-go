package msg

type MqttMsg struct {
	Topic string
	Msg string
}

type UserEvent int 
const (
	Login	= iota
	Logout
	Create
	Close
	ChooseNGHero
	Invite
	StartQueue
	PreStart
	Join
	StartGame
	Start
	StartGet
	GameSignal
)

type ServerMsg struct {
	Event UserEvent
	Id string
	Msg string
}

// use serde_derive::{Serialize, Deserialize};

// #[derive(Serialize, Deserialize, Clone, Debug)]
// pub struct MqttMsg {
//     pub topic: String,
//     pub msg: String,
// }
