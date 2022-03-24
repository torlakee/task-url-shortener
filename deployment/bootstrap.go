package deployment

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/couchbase/gocb/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartUp() *Application {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	appConfig := initConfig()

	storage := prepareStorage(appConfig)

	app := &Application{
		storage,
		"URL-Shortener",
		appConfig,
	}

	return app
}

func initConfig() *AppConf {

	var conf *AppConf
	var err error

	if _, err = toml.DecodeFile("app.cnf", &conf); err != nil {
		fmt.Println(err)
	}

	return conf
}

func prepareStorage(config *AppConf) *gocb.Bucket {

	db := config.Internal.Database

	cluster, err := gocb.Connect(
		db.Host,
		gocb.ClusterOptions{
			Username: db.User,
			Password: db.Password,
		})

	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket(db.Storage)

	// gocb.SetLogger(gocb.VerboseStdioLogger())

	err = bucket.WaitUntilReady(15*time.Second, nil)

	if err != nil {
		panic(err)
	}

	return bucket
}

func CleanUp() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)

	go func() {
		s := <-sigs
		log.Printf("[SIG] %s", s)
		os.Exit(1)
	}()
}

func Reloadable() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGHUP)
	go func() {
		for {
			<-s
		}
	}()
}
