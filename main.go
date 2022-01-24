package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ggrrrr/bui_api_people/resources"
	"github.com/ggrrrr/bui_lib/api"
	"github.com/ggrrrr/bui_lib/db/mgo"

	"github.com/ggrrrr/bui_lib/token"
)

var (
	root = context.Background()
	err  error
)

func main() {
	flag.Parse()

	// err = mgo.Configure()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	server()
	// log.Printf("end.")
}

func server() {

	err = token.Configure()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = api.Configure()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err := mgo.Configure(root)
	if err != nil {
		// mgo.Configure(ctx)
		log.Fatal(err)
	}

	go func() {
		time.Sleep(5 * time.Second)
		api.Ready()
	}()

	err = api.Create(root, false)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// api.HandleFunc("/userLogin", auth.UserLoginRequest)
	// api.HandleFunc("/tokenVerify", auth.TokenVerifyRequest)
	api.HandleFunc("/people/list", resources.ListRequest)
	api.HandleFunc("/people/person/update", resources.Update)
	api.HandleFunc("/people/person/insert", resources.Insert)
	api.HandleFunc("/people/pin/parse", resources.ParsePin)

	osSignals := make(chan os.Signal, 1)
	go func() {
		err := api.Start()
		defer api.Shutdown()
		if err != nil {
			log.Printf("http error: %+v", err)
			osSignals <- os.Kill
		}
	}()

	signal.Notify(osSignals, os.Interrupt)
	log.Printf("os.signal: %v", <-osSignals)
	api.Shutdown()
	// db.Session.Close()
	log.Printf("end.")
}
