package main

import (
	"context"
	"goback1/lesson6/reguser/api/handler"
	"goback1/lesson6/reguser/api/routergin"
	"goback1/lesson6/reguser/api/server"
	"goback1/lesson6/reguser/app/repos/user"
	"goback1/lesson6/reguser/app/starter"
	"goback1/lesson6/reguser/db/fstore/userfstore"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	ust, err := userfstore.NewUserFileStore("./db")
	if err != nil {
		log.Fatal(err)
	}
	defer ust.Close()

	a := starter.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewHandlers(us)

	rh := routergin.NewRouterGin(h)

	srv := server.NewServer(":8000", rh)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
