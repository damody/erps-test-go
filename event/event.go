package event

 import (
 	"fmt"
 	"math/rand"
	"encoding/json"
	"time"
	"strings"
 	m "erps-test-go/msg"
 	u "erps-test-go/user"
 )

type LoginRes struct {
	Msg string
	Ng int
	Rk int
}

type LoginMsg struct {
	Id string
	Msg string
}

type LogoutRes struct {
	Msg string
}

type LogoutMsg struct {
	Id string
	Msg string
}

type CreateRoomRes struct {
	Msg string
}

type CreateRoomMsg struct {
	Id string
	Msg string
}

type CloseRoomRes struct {
	Msg string
}

type CloseRoomMsg struct {
	Room string
	Msg string
}

type UserNGHeroRes struct {
	Id string
	Hero string
}

type UserNGHeroMsg struct {
	Id string
	Hero string
}

type InviteRes struct {
	Msg string
}

type InviteMsg struct {
	Id string
	Msg string
}

type StartQueueRes struct {
	Msg string
}

type StartQueueMsg struct {
	Id string
	Msg string
}

type PreStartRes struct {
	Msg string
}

type PreStartMsg struct {
	Id string
	Msg string
}

type JoinRes struct {
	Room string
	Msg string
}

type JoinMsg struct {
	Id string
	Room string
	Msg string
}

type StartRes struct {
	Game uint32
	Room string
	Msg string
}

type StartGetMsg struct {
	Id string
	Msg string
}

type StartGameRes struct {
	Game uint32
	Member []Herocell
}

type Herocell struct {
	Id string
	Team uint
	Name string
	Hero string
	Buff map[string]float32
	Tags []string
}

type GameSignalRes struct {
	Game uint32
}

type GameOverData struct {
	Game uint32
	Win []string
	Lose []string
}

type GameInfoData struct {
	Game uint32
	Users []UserInfoData
}

type UserInfoData struct {
	Id string
	Hero string
	Level uint16
	Equ []string
	Damage uint16
	Take_damage uint16
	Heal uint16
	Kill uint16
	Death uint16
	Assist uint16
	Gift UserGift
}

type UserGift struct {
	A uint16
	B uint16
	C uint16
	D uint16
	E uint16
}

func Get_user(id string, users map[string]*u.User) *u.User {
	return users[id]
}

var ServerChan = make(chan m.ServerMsg)

func Init(msgtx chan<-m.MqttMsg){	
	update500ms := time.NewTicker(500*time.Millisecond)
	
	TotalUsers := map[string]*u.User{}
	rooms := map[string]*u.RoomRecord{}

	for i:=0; i<10; i++{
		TotalUsers[fmt.Sprintf("%d", i)] = &u.User{
			Id: fmt.Sprintf("%d", i), 
			Hero: "", 
			Cnt: -1,
			IsLogin: false,
			IsInRoom: false,
			IsRoomCreater: false,
			IsChooseNGHero: false,
			IsStartQueue: false,
			IsCanPrestart: false,
			IsPrestart: false,
			IsPlaying: false,
		}
	}	
	
	go func() {
		for {
			select {
				case <- update500ms.C:
					for _, v := range TotalUsers {
						v.Next_action(msgtx, &rooms)
					}
					

				case val := <-ServerChan: 
					if (val.Event == m.Login) {
						//fmt.Printf("In Channel: %s", val.Msg)
						p := &LoginRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Get_login()

					} else if (val.Event == m.Logout) {
						p := &LogoutRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Get_logout()

					} else if (val.Event == m.Create) {
						p := &CreateRoomRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						if (p.Msg == "ok") {
							TotalUsers[val.Id].Get_create()
						} 

					} else if (val.Event == m.Close) {
						p := &CloseRoomRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Get_close()

					} else if (val.Event == m.ChooseNGHero) {
						p := &UserNGHeroRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Get_choose_hero(p.Hero)
					} else if (val.Event == m.Invite) {
						p := &InviteRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Get_invite()

					} else if (val.Event == m.StartQueue) {
						p := &StartQueueRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Get_start_queue()

					} else if (val.Event == m.PreStart) {
						p := &PreStartRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						if (p.Msg == "stop queue") {
							TotalUsers[val.Id].Get_prestart(false, msgtx)
						} else {
							for _, id:= range rooms[val.Id].Ids {
								TotalUsers[id].Get_prestart(true, msgtx)
							}
						}

					} else if (val.Event == m.Join) {
						p := &JoinRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						if (p.Msg == "ok") {
							TotalUsers[val.Id].Get_join(p.Room)
						}
						rooms[p.Room].Ids = append(rooms[p.Room].Ids, val.Id)

					} else if (val.Event == m.StartGame) {
						fmt.Printf("Start Game!\n")
						p := &StartGameRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						data := new(GameOverData)
						data1 := new(GameInfoData)
						rand.Seed(time.Now().UnixNano())
						rng := rand.Intn(2)

						for _, m := range p.Member {
							if (m.Team == uint(rng+1)) {
								data.Win = append(data.Win, m.Id)
							} else {
								data.Lose = append(data.Lose, m.Id)
							}
							TotalUsers[m.Id].Game_over()
							userinfo := new(UserInfoData)
							userinfo.Id = m.Id
							userinfo.Hero = TotalUsers[m.Id].Hero

							userinfo.Level = uint16(rand.Intn(3) + 13)
							userinfo.Damage = uint16(rand.Intn(2001) + 1000)
							userinfo.Equ = append(userinfo.Equ, "bz")
							userinfo.Equ = append(userinfo.Equ, "uti")
							userinfo.Equ = append(userinfo.Equ, "666")
							
							userinfo.Take_damage = uint16(rand.Intn(2001) + 1000)
							userinfo.Heal = uint16(rand.Intn(501) + 500)

							userinfo.Kill = uint16(rand.Intn(5))
							userinfo.Death = uint16(rand.Intn(4))
							userinfo.Assist = uint16(rand.Intn(6))

							userinfo.Gift.A = uint16(rand.Intn(4))
							userinfo.Gift.B = uint16(rand.Intn(4))
							userinfo.Gift.C = uint16(rand.Intn(4))
							userinfo.Gift.D = uint16(rand.Intn(4))
							userinfo.Gift.E = uint16(rand.Intn(4))

							data1.Users = append(data1.Users, *userinfo)
						}
						data.Game = p.Game
						data1.Game = p.Game

						d, _ := json.Marshal(data)
						d1, _ := json.Marshal(data1)

						topic := fmt.Sprintf("game/%d/send/game_over", p.Game)
						msg := string(d)
						msgtx <- m.MqttMsg{Topic: topic, Msg: strings.ToLower(msg)}

						topic = fmt.Sprintf("game/%d/send/game_info", p.Game)
						msg = string(d1)
						msgtx <- m.MqttMsg{Topic: topic, Msg: strings.ToLower(msg)}


					} else if (val.Event == m.Start) {
						
					} else if (val.Event == m.StartGet) {
						p := &StartGetMsg{}
						json.Unmarshal([]byte(val.Msg), &p)
						TotalUsers[val.Id].Start_get()

					} else if (val.Event == m.GameSignal) {
						p := &GameSignalRes{}
						json.Unmarshal([]byte(val.Msg), &p)
						topic := fmt.Sprintf("game/%d/send/start_game", p.Game)
						msg := fmt.Sprintf(`{"game":%d, "action":"init"}`, p.Game)
						msgtx <- m.MqttMsg{Topic: topic, Msg: msg}
					}
					
				
			}
		}
	}()
}

func Login(id string, msg string, sender chan<-m.ServerMsg){
	
}

func logout(id string, msg string, sender chan<-m.ServerMsg){
	
}

func create(id string, msg string, sender chan<-m.ServerMsg){
	
}

func close(id string, msg string, sender chan<-m.ServerMsg){
	
}

func choose_hero(id string, msg string, sender chan<-m.ServerMsg){
	
}

func start_queue(id string, msg string, sender chan<-m.ServerMsg){
	
}

func start_get(id string, msg string, sender chan<-m.ServerMsg){
	
}

func prestart(id string, msg string, sender chan<-m.ServerMsg){
	
}

func join(id string, msg string, sender chan<-m.ServerMsg){
	
}

func start(id string, msg string, sender chan<-m.ServerMsg){
	
}

func start_game(id string, msg string, sender chan<-m.ServerMsg){
	
}

func game_singal(id string, msg string, sender chan<-m.ServerMsg){
	
}