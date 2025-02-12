package main

import (
	pr "Go_grpc_Server-streaming/protoc"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pr.UnimplementedNewsModelServer
}

func main() {

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("cant run the server", err.Error())
	}

	srv := grpc.NewServer()
	pr.RegisterNewsModelServer(srv, &server{})

	if err := srv.Serve(listen); err != nil {
		log.Fatal("cant serve the listener", err.Error())
	}

}

func (s *server) TodaysNews(req *pr.Request, stream pr.NewsModel_TodaysNewsServer) error {

	sl := []string{
		"🌍🔥🚨 Global temperatures hit record highs!",
		"📉💰📊 Stock market takes a sharp dip!",
		"🚗⚡🔋 New EV battery charges in 5 minutes!",
		"🏆⚽🇦🇷 Messi wins another Ballon d'Or!",
		"🛰️🌕🚀 NASA plans new moon mission!",
		"🍔🚫⚠️ Fast food chain recalls products!",
		"🎬🍿💥 Blockbuster movie breaks records!",
	}

	for _, value := range sl {
		response := pr.Reciever{Headlines: value}

		if err := stream.Send(&response); err != nil {
			log.Println("something went wrone while send the data through the stream", err.Error())
			return err

		}
		time.Sleep(time.Second * 2)

	}
	return nil

}
