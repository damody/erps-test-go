package user

import (
	//"fmt"
	"math/rand"
	"time"
	m "erps-test-go/msg"
)

type RoomRecord struct {
	Id string
	Ids []string
}

const TEAM_SIZE = 1

type User struct {
	Id string
	Hero string
	Room string
	Cnt int
	IsLogin bool
	IsInRoom bool
	IsRoomCreater bool
	IsChooseNGHero bool
	IsStartQueue bool
	IsCanPrestart bool
	IsPrestart bool
	IsPlaying bool
}

func (u *User) Next_action(tx chan<-m.MqttMsg, rooms *map[string]*RoomRecord  ) {
	rand.Seed(time.Now().UnixNano())
	rng := rand.Intn(10)
	if (rng > 4) {
		return 
	}
	if (u.Cnt >= 0 && u.Cnt < 10 || u.IsPlaying) {
		u.Cnt += 1
		return
	}
	if (!u.IsLogin) {
		u.Login(tx)
	} else if (!u.IsChooseNGHero) {
		u.Choose_hero(tx, "freyja")
	} else if (!u.IsInRoom) {
		rng = rand.Intn(10)
		if (rng) < 5 {
			u.Create(tx)
			id:= u.Id
			(*rooms)[id] = &RoomRecord {
				Id: id,
				Ids: []string{id},
			}
		} else if len(*rooms) > 0 {
			rng = rand.Intn(len(*rooms))
			cnt:=0;
			for _, r := range *rooms{
				if (cnt == rng) {
					if (len(r.Ids) < TEAM_SIZE){
						u.Join(tx, r)
					}
					cnt++
				}
			}
		}
	} else if (!u.IsStartQueue) {
		u.Start_queue(tx)
	} else if (u.IsCanPrestart) {
		
		u.Prestart(tx)
	}
	u.Cnt = 0;
}

func (u *User) Back_action(tx chan<-m.MqttMsg) {
	if (u.IsLogin) {
		u.Logout(tx)
	} else if (u.IsInRoom && u.IsRoomCreater) {
		u.Close(tx)
	} else if (!u.IsStartQueue) {
		u.Start_queue(tx)
	} else if (u.IsCanPrestart) {
		u.Prestart(tx)
	}
}

func (u *User) Login(tx chan<-m.MqttMsg) {
	if (!u.IsLogin) {
		topic := "member/" + u.Id + "/send/login"
		msg := `{"id":"` + u.Id + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Join(tx chan<-m.MqttMsg, room *RoomRecord) {
	if (!u.IsInRoom) {
		topic := "room/" + u.Id + "/send/join"
		msg := `{"room:"` + room.Id + `", "join":"` + u.Id + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Get_join(room string) {
	u.IsInRoom = true
	u.Room = room
	u.Cnt = -1
}

func (u *User) Get_login() {
	u.IsLogin = true;
	u.Cnt = -1
}

func (u *User) Logout(tx chan<-m.MqttMsg) {
	if (u.IsLogin) {
		topic := "member/" + u.Id + "/send/logout"
		msg := `{"id":"` + u.Id + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Game_over() {
	u.IsStartQueue = false
	u.IsPrestart = false
	u.IsPlaying = false
	u.IsInRoom = false
	u.IsRoomCreater = false
	u.IsCanPrestart = false
}

func (u *User) Get_logout() {
	u.Cnt = -1
	u.IsLogin = false
	u.IsChooseNGHero = false
	u.IsStartQueue = false
	u.IsPrestart = false
	u.IsPlaying = false
	u.IsInRoom = false
	u.IsRoomCreater = false
}

func (u *User) Choose_hero(tx chan<-m.MqttMsg, hero string) {
	u.Hero = hero
	if (!u.IsChooseNGHero) {
		topic := "member/" + u.Id +"/send/choose_hero"
		msg := `{"id":"`+ u.Id + `", "hero":"` + u.Hero + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Get_choose_hero(hero string) {
	u.Hero = hero
	u.IsChooseNGHero = true
	u.Cnt = -1
}

func (u User) Choose_random_hero(tx chan<-m.MqttMsg) {

}

func (u *User) Create(tx chan<-m.MqttMsg) {
	if (!u.IsInRoom) {
		topic := "room/" + u.Id + "/send/create"
		msg := `{"id":"` + u.Id + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
		u.Room = u.Id
	}
}

func (u *User) Get_create() {
	u.IsInRoom = true
	u.Room = u.Id
	u.IsRoomCreater = true
	u.Cnt = -1
}

func (u *User) Close(tx chan<-m.MqttMsg) {
	if (u. IsInRoom) {
		topic := "room/" + u.Id + "/send/close"
		msg := `{"id":"` + u.Id + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Get_close() {
	u.IsInRoom = false
	u.Room = ""
	u.Cnt = -1
}

func (u *User) Start_queue(tx chan<-m.MqttMsg) {
	if (u.IsInRoom && !u.IsRoomCreater) {
		u.IsStartQueue = true
	}
	if (!u.IsStartQueue) {
		topic := "room/" + u.Room + "/send/start_queue"
		msg := `{"id":"`+ u.Id + `", "action":"start queue"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Get_start_queue() {
	u.IsStartQueue = true
	u.Cnt = -1
}

func (u *User) Start_get() {
	u.IsPrestart = true
	u.IsPlaying = true
}

func (u *User) Prestart(tx chan<-m.MqttMsg) {
	u.Cnt = -1
	if (u.IsCanPrestart && !u.IsPrestart) {
		topic := "room/" + u.Id + "/send/prestart"
		msg := `{"room":"`+ u.Room + `", "id":"` + u.Id + `", "accept":true}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}
}

func (u *User) Get_prestart(res bool, tx chan<-m.MqttMsg) {
	u.Cnt = -1

	if (u.IsPrestart) {
		return
	}
	if (!res) {
		u.IsPrestart = false
	} else {
		u.IsCanPrestart = res
		topic := "room/" + u.Id + "/send/prestart_get"
		msg := `{"room":"`+ u.Room + `", "id":"` + u.Id + `"}`
		tx <- m.MqttMsg{Topic: topic, Msg: msg}
	}

}

func (u User) Invite(tx chan<-m.MqttMsg) {

}

func (u User) Get_invite(){

}

func (u *User) afk() {
	u.IsLogin = false
	u.IsChooseNGHero = false
	u.IsStartQueue = false
	u.IsPrestart = false
	u.IsPlaying = false
}
