// #![allow(warnings)]
// use log::{info, warn, error, trace};
package main

import (
	"fmt"	
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
	"regexp"
	m "erps-test-go/msg"
	e "erps-test-go/event"
)


// use std::env;
// use std::io::Write;
// use failure::Error;
// use std::net::TcpStream;
// use std::str;
// use clap::{App, Arg};
// use uuid::Uuid;
// use rumqtt::{MqttClient, MqttOptions, QoS};

// use std::thread;
// use std::time::Duration;
// use log::Level;
// use serde_json::{self, Value};
// use regex::Regex;
// use ::futures::Future;
// use mysql;

// use crossbeam_channel::{bounded, tick, Sender, Receiver, select};

// mod msg;
// mod event;
// mod user;
// use crate::msg::*;
// use crate::event::*;



var MqttChan = make(chan m.MqttMsg)

// fn generate_client_id() -> String {
//     format!("ERPS_TEST_{}", Uuid::new_v4())
// }
func generate_client_id() string {
	unix32bits := uint32(time.Now().UTC().Unix())
	uuid := fmt.Sprintf("%x",unix32bits)
	return uuid
}

func main() {


//     let server_addr = matches.value_of("SERVER").unwrap_or("59.126.81.58").to_owned();
//     let server_port = matches.value_of("PORT").unwrap_or("1883").to_owned();
//     let client_id = matches
//         .value_of("CLIENT_ID")
//         .map(|x| x.to_owned())
//         .unwrap_or_else(generate_client_id);
//     let mut mqtt_options = MqttOptions::new(client_id.as_str(), server_addr.as_str(), server_port.parse::<u16>().unwrap());
//     mqtt_options = mqtt_options.set_keep_alive(100);
//     mqtt_options = mqtt_options.set_request_channel_capacity(10000);
//     mqtt_options = mqtt_options.set_notification_channel_capacity(100000);
//     let (mut mqtt_client, notifications) = MqttClient::start(mqtt_options.clone()).unwrap();




	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	e.Init(MqttChan)
	//go func() {
	opts := MQTT.NewClientOptions().AddBroker("59.126.81.58:1883")
	opts.SetClientID(generate_client_id())
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {

		panic(token.Error())
	}

	c.Subscribe("member/+/res/login", 0, nil);
	c.Subscribe("member/+/res/logout", 0, nil);
	c.Subscribe("member/+/res/choose_hero", 0, nil);
	c.Subscribe("room/+/res/create", 0, nil);
	c.Subscribe("room/+/res/close", 0, nil);
	c.Subscribe("room/+/res/start_queue", 0, nil);
	c.Subscribe("room/+/res/cancel_queue", 0, nil); 
	c.Subscribe("room/+/res/invite", 0, nil);
	c.Subscribe("room/+/res/join", 0, nil);
	c.Subscribe("room/+/res/accept_join", 0, nil);
	c.Subscribe("room/+/res/kick", 0, nil);
	c.Subscribe("room/+/res/leave", 0, nil);
	c.Subscribe("room/+/res/prestart", 0, nil);
	c.Subscribe("room/+/res/start", 0, nil);
	c.Subscribe("room/+/res/start_get", 0, nil);
	c.Subscribe("game/+/res/game_signal", 0, nil);
	c.Subscribe("game/+/res/game_over", 0, nil);
	c.Subscribe("game/+/res/start_game", 0, nil);
	c.Subscribe("game/+/res/choose", 0, nil);
	c.Subscribe("game/+/res/exit", 0, nil);

//     mqtt_client.subscribe("member/+/res/login", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("member/+/res/logout", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("member/+/res/choose_hero", QoS::AtLeastOnce).unwrap();

//     mqtt_client.subscribe("room/+/res/create", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/close", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/start_queue", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/cancel_queue", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/invite", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/join", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/accept_join", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/kick", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/leave", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/prestart", QoS::AtLeastOnce).unwrap();
//     //mqtt_client.subscribe("room/+/res/prestart_get", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/start", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("room/+/res/start_get", QoS::AtLeastOnce).unwrap();

//     mqtt_client.subscribe("game/+/res/game_singal", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("game/+/res/game_over", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("game/+/res/start_game", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("game/+/res/choose", QoS::AtLeastOnce).unwrap();
//     mqtt_client.subscribe("game/+/res/exit", QoS::AtLeastOnce).unwrap();



	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	

//     let (tx, rx):(Sender<MqttMsg>, Receiver<MqttMsg>) = bounded(10000);
//     let mut sender: Sender<UserEvent> = event::init(tx.clone());
//     thread::sleep_ms(100);
//     for _ in 0..16 {
//         let server_addr = server_addr.clone();
//         let server_port = server_port.clone();
//         let rx = rx.clone();
//         thread::spawn(move || {
//             let mut pkid = 100;
//             let mut mqtt_options = MqttOptions::new(generate_client_id(), server_addr, server_port.parse::<u16>().unwrap());
//             mqtt_options = mqtt_options.set_keep_alive(100);
//             mqtt_options = mqtt_options.set_request_channel_capacity(10000);
//             mqtt_options = mqtt_options.set_notification_channel_capacity(10000);
//             let (mut mqtt_client, notifications) = MqttClient::start(mqtt_options.clone()).unwrap();
//             let update = tick(Duration::from_millis(1000));
//             loop {
//                 select! {
//                     recv(update) -> _ => {
//                         let size = rx.len();
//                         if size > 0 {
//                             println!("publish rx len: {}", rx.len());
//                         }
//                     },
//                     recv(rx) -> d => {
//                         let handle = || -> Result<(), Error> {
//                             if let Ok(d) = d {
//                                 match mqtt_client.publish(d.topic, QoS::AtLeastOnce, false, d.msg) {
//                                     Ok(_) => {},
//                                     Err(x) => {
//                                         println!("publish failed!!!!");
//                                     }
//                                 }
//                             }
//                             Ok(())
//                         };
//                         if let Err(msg) = handle() {
//                             println!("{:?}", msg);
//                         }
//                     }
//                 }
//             }
//         });
//     }


	// Mqtt Client send message
	
	for {
		select {
		case val := <- MqttChan:
			//fmt.Printf("Channel Publish!\n")
			token := c.Publish(val.Topic, 0, false, val.Msg)
			token.Wait()
		}
	}
	//}()
	
}
// fn main() -> std::result::Result<(), Error> {
//     // configure logging
//     env::set_var("RUST_LOG", env::var_os("RUST_LOG").unwrap_or_else(|| "info".into()));
//     env_logger::init();

//     let matches = App::new("erps")
//         .author("damody <t1238142000@gmail.com>")
//         .arg(
//             Arg::with_name("SERVER")
//                 .short("S")
//                 .long("server")
//                 .takes_value(true)
//                 .help("MQTT server address (127.0.0.1)"),
//         ).arg(
//             Arg::with_name("PORT")
//                 .short("P")
//                 .long("port")
//                 .takes_value(true)
//                 .help("MQTT server port (1883)"),
//         ).arg(
//             Arg::with_name("USER_NAME")
//                 .short("u")
//                 .long("username")
//                 .takes_value(true)
//                 .help("Login user name"),
//         ).arg(
//             Arg::with_name("PASSWORD")
//                 .short("p")
//                 .long("password")
//                 .takes_value(true)
//                 .help("Password"),
//         ).arg(
//             Arg::with_name("CLIENT_ID")
//                 .short("i")
//                 .long("client-identifier")
//                 .takes_value(true)
//                 .help("Client identifier"),
//         ).get_matches();

 

var relogin, _ = regexp.Compile("(\\w+)/(\\w+)/res/login")
var relogout, _ = regexp.Compile("(\\w+)/(\\w+)/res/logout")
var recreate, _ = regexp.Compile("(\\w+)/(\\w+)/res/create")
var reclose, _ = regexp.Compile("(\\w+)/(\\w+)/res/close")
var restart_queue, _ = regexp.Compile("(\\w+)/(\\w+)/res/start_queue")
var recancel_queue, _ = regexp.Compile("(\\w+)/(\\w+)/res/cancel_queue")
var represtart, _ = regexp.Compile("(\\w+)/(\\w+)/res/prestart")
var reinvite, _ = regexp.Compile("(\\w+)/(\\w+)/res/invite")
var rejoin, _ = regexp.Compile("(\\w+)/(\\w+)/res/join")
var rechoose_hero, _ = regexp.Compile("(\\w+)/(\\w+)/res/choose_hero")
var restart_game, _ = regexp.Compile("(\\w+)/(\\w+)/res/start_game")
var restart_get, _ = regexp.Compile("(\\w+)/(\\w+)/res/start_get")
var restart, _ = regexp.Compile("(\\w+)/(\\w+)/res/start")
var regame_signal, _ = regexp.Compile("(\\w+)/(\\w+)/res/game_signal")

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {	

	if (relogin.MatchString(msg.Topic())) {		
		substr := relogin.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Login, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg

	} else if (relogout.MatchString(msg.Topic())) {
		
		substr := relogout.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Logout, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (recreate.MatchString(msg.Topic())) {
		//fmt.Printf("Create\n");
		substr := recreate.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Create, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (reclose.MatchString(msg.Topic())) {
		//fmt.Printf("Close\n");
		substr := reclose.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Close, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (restart_queue.MatchString(msg.Topic())) {
		//fmt.Printf("Start Queue\n");
		substr := restart_queue.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.StartQueue, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (represtart.MatchString(msg.Topic())) {
		//fmt.Printf("Prestart\n");
		substr := represtart.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.PreStart, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (reinvite.MatchString(msg.Topic())) {
		//fmt.Printf("Invite\n");
		substr := reinvite.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Invite, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (rejoin.MatchString(msg.Topic())) {
		//fmt.Printf("Join\n");
		substr := rejoin.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Join, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (rechoose_hero.MatchString(msg.Topic())) {
		substr := rechoose_hero.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.ChooseNGHero, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (restart_game.MatchString(msg.Topic())) {
		//fmt.Printf("Start Game\n");
		substr := restart_game.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.StartGame, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (restart_get.MatchString(msg.Topic())) {
		//fmt.Printf("Start Get\n");
		substr := restart_get.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.StartGet, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (restart.MatchString(msg.Topic())) {
		//fmt.Printf("Start\n");
		substr := restart.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.Start, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	} else if (regame_signal.MatchString(msg.Topic())) {
		//fmt.Printf("Game Signal\n");
		substr := regame_signal.FindStringSubmatch(msg.Topic())[2]
		Msg:= m.ServerMsg{Event: m.GameSignal, Id: substr, Msg: string(msg.Payload())}
		e.ServerChan <- Msg
	}
}

//     let relogin = Regex::new(r"\w+/(\w+)/res/login").unwrap();
//     let relogout = Regex::new(r"\w+/(\w+)/res/logout").unwrap();
//     let recreate = Regex::new(r"\w+/(\w+)/res/create").unwrap();
//     let reclose = Regex::new(r"\w+/(\w+)/res/close").unwrap();
//     let restart_queue = Regex::new(r"\w+/(\w+)/res/start_queue").unwrap();
//     let recancel_queue = Regex::new(r"\w+/(\w+)/res/cancel_queue").unwrap();
//     let represtart = Regex::new(r"\w+/(\w+)/res/prestart").unwrap();
//     //let represtart_get = Regex::new(r"\w+/(\w+)/res/prestart_get").unwrap();
//     let reinvite = Regex::new(r"\w+/(\w+)/res/invite").unwrap();
//     let rejoin = Regex::new(r"\w+/(\w+)/res/join").unwrap();
//     let rechoosehero = Regex::new(r"\w+/(\w+)/res/choose_hero").unwrap();
//     let restart_game = Regex::new(r"\w+/(\w+)/res/start_game").unwrap();
//     let restart_get = Regex::new(r"\w+/(\w+)/res/start_get").unwrap();
//     let restart = Regex::new(r"\w+/(\w+)/res/start").unwrap();
//     let regame_singal = Regex::new(r"\w+/(\w+)/res/game_singal").unwrap();

//     let redead = Regex::new(r"\w+/(\w+)/res/dead").unwrap();

//     loop {
//         use rumqtt::Notification::Publish;
//         select! {
//             recv(notifications) -> notification => {
//                 let handle = || -> Result<(), Error> {
//                     if let Ok(x) = notification {
//                         if let Publish(x) = x {
//                             //println!("{:?}", x);
//                             let payload = x.payload;
//                             let msg = match str::from_utf8(&payload[..]) {
//                                 Ok(msg) => msg,
//                                 Err(err) => {
//                                     return Err(failure::err_msg(format!("Failed to decode publish message {:?}", err)));
//                                 }
//                             };
//                             let topic_name = x.topic_name.as_str();
//                             let vo : serde_json::Result<Value> = serde_json::from_str(msg);
//                             if let Ok(v) = vo {
//                                 if relogin.is_match(topic_name) {
//                                     let cap = relogin.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     info!("login: userid: {} json: {:?}", userid, v);
//                                     event::login(userid, v, sender.clone())?;
//                                 }
//                                 else if relogout.is_match(topic_name) {
//                                     let cap = relogout.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("logout: userid: {} json: {:?}", userid, v);
//                                     event::logout(userid, v, sender.clone())?;
//                                 }
//                                 else if recreate.is_match(topic_name) {
//                                     let cap = recreate.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("create: userid: {} json: {:?}", userid, v);
//                                     event::create(userid, v, sender.clone())?;
//                                 }
//                                 else if reclose.is_match(topic_name) {
//                                     let cap = reclose.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("close: userid: {} json: {:?}", userid, v);
//                                     event::close(userid, v, sender.clone())?;
//                                 }
//                                 else if rejoin.is_match(topic_name) {
//                                     let cap = rejoin.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("join: userid: {} json: {:?}", userid, v);
//                                     event::join(userid, v, sender.clone())?;
//                                 }
//                                 else if restart_queue.is_match(topic_name) {
//                                     let cap = restart_queue.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("start_queue: userid: {} json: {:?}", userid, v);
//                                     event::start_queue(userid, v, sender.clone())?;
//                                 }
//                                 else if rechoosehero.is_match(topic_name) {
//                                     let cap = rechoosehero.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("choose hero: userid: {} json: {:?}", userid, v);
//                                     event::choose_hero(userid, v, sender.clone())?;
//                                 }
//                                 else if represtart.is_match(topic_name) {
//                                     let cap = represtart.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("prestart hero: userid: {} json: {:?}", userid, v);
//                                     event::prestart(userid, v, sender.clone())?;
//                                 }
//                                 else if restart_get.is_match(topic_name) {
//                                     let cap = restart_get.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("prestart hero: userid: {} json: {:?}", userid, v);
//                                     event::start_get(userid, v, sender.clone())?;
//                                 }
//                                 else if restart_game.is_match(topic_name) {
//                                     let cap = restart_game.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("start_game hero: userid: {} json: {:?}", userid, v);
//                                     event::start_game(userid, v, sender.clone())?;
//                                 }
//                                 else if restart.is_match(topic_name) {
//                                     let cap = restart.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("start hero: userid: {} json: {:?}", userid, v);
//                                     event::start(userid, v, sender.clone())?;
//                                 }
//                                 else if regame_singal.is_match(topic_name) {
//                                     let cap = regame_singal.captures(topic_name).unwrap();
//                                     let userid = cap[1].to_string();
//                                     //info!("game_singal: userid: {} json: {:?}", userid, v);
//                                     event::game_singal(userid, v, sender.clone())?;
//                                 }
//                                 else if redead.is_match(topic_name) {
//                                     thread::sleep(Duration::from_millis(5000));
//                                 }
//                                 else {
//                                     warn!("Topic Error {}", topic_name);
//                                 }
//                             } else {
//                                 warn!("Json Parser error");
//                             };
//                         }
//                     }
//                     Ok(())
//                 };
//                 if let Err(msg) = handle() {
//                     println!("{:?}", msg);
//                 }
//             }
//         }
//     }
//     Ok(())
// }






