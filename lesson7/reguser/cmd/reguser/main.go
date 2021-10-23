package main

import (
	"context"
	"goback1/lesson7/reguser/api/handler"
	"goback1/lesson7/reguser/api/routeroapi"
	"goback1/lesson7/reguser/api/server"
	"goback1/lesson7/reguser/app/repos/user"
	"goback1/lesson7/reguser/app/starter"
	"goback1/lesson7/reguser/db/sql/pgstore"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	ust, err := pgstore.NewUsers("postgres://postgres:1110@127.0.0.1/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer ust.Close()

	a := starter.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewHandlers(us)

	rh := routeroapi.NewRouterOpenAPI(h)

	srv := server.NewServer(":8000", rh)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
