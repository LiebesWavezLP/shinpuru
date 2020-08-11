package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/zekroTJA/shinpuru/internal/inits"

	"github.com/zekroTJA/shinpuru/internal/core/config"
	"github.com/zekroTJA/shinpuru/internal/core/middleware"
	"github.com/zekroTJA/shinpuru/internal/util"
)

var (
	flagConfigLocation = flag.String("c", "config.yml", "The location of the main config file")
	flagDocker         = flag.Bool("docker", false, "wether shinpuru is running in a docker container or not")
)

//////////////////////////////////////////////////////////////////////
//
//   SHINPURU
//   --------
//   This is the main initialization for shinpuru which initializes
//   all instances like the database middleware, the twitch notify
//   listener service, life cycle timer, storage middleware,
//   permission middleware, command handler and - finally -
//   initializes the discord session event loop.
//   shinpuru is configured via a configuration file which location
//   can be passed via the '-c' parameter.
//   When shinpuru is run in a Docker container, the '-docker' flag
//   should be passed to fix configuration values like the location
//   of the sqlite3 database (when the sqlite3 driver is used) or
//   the web server exposure port.
//
//////////////////////////////////////////////////////////////////////

func main() {
	// Parse command line flags
	flag.Parse()

	// Initial log output
	util.Log.Infof("シンプル (shinpuru) v.%s (commit %s)", util.AppVersion, util.AppCommit)
	util.Log.Info("© zekro Development (Ringo Hoffmann)")
	util.Log.Info("Covered by MIT Licence")
	util.Log.Info("Starting up...")

	// Initialize discordgo session
	session, err := discordgo.New()
	if err != nil {
		util.Log.Fatal(err)
	}

	// Initialize config
	conf := inits.InitConfig(*flagConfigLocation, new(config.YAMLConfigParser))

	// Set static config values when docker flag is passed
	if *flagDocker {
		if conf.Database.Sqlite == nil {
			conf.Database.Sqlite = new(config.DatabaseFile)
		}
		conf.Database.Sqlite.DBFile = "/etc/db/db.sqlite3"
		conf.WebServer.Addr = ":8080"
	}

	// Setting log level from config
	util.SetLogLevel(conf.Logging.LogLevel)

	// Initialize database middleware and shutdown routine
	database := inits.InitDatabase(conf.Database)
	defer func() {
		util.Log.Info("Shutting down database connection...")
		database.Close()
	}()

	// Initialize twitch notify handler and shutdown routine
	tnw, tnl := inits.InitTwitchNotifyer(session, conf, database)
	defer func() {
		util.Log.Info("Tearing down twitch notify listener...")
		tnl.TearDown()
	}()

	// Initialize life cycle timer
	lct := inits.InitLTCTimer()

	// Initialize storage middleware
	st := inits.InitStorage(conf)

	// Initialize permissions command handler middleware
	pmw := middleware.NewPermissionMiddleware(database, conf)
	// Initialize ghost ping ignore command handler middleware
	gpim := middleware.NewGhostPingIgnoreMiddleware()

	// Initialize discord bot session and shutdown routine
	inits.InitDiscordBotSession(session, conf, database, lct, pmw, gpim)
	defer func() {
		util.Log.Info("Shutting down bot session...")
		session.Close()
	}()

	// Initialize command handler
	cmdHandler := inits.InitCommandHandler(
		session, conf, database, st, tnw, lct, pmw, gpim)

	// Initialize web server
	inits.InitWebServer(session, database, st, cmdHandler, lct, conf, pmw)

	// Block main go routine until one of the following
	// specified exit syscalls occure.
	util.Log.Info("Started event loop. Stop with CTRL-C...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
