package cmd_serve

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"log"
)

var db *sql.DB

func serve(conf TomlConfig) {
	// Print Header
	fmt.Println("     __ _  ___  _ __ ___  ___      ")
	fmt.Println("    / _` |/ _ \\| '__/ __|/ _ \\   ")
	fmt.Println("   | (_| | (_) | |  \\__ \\  __/   ")
	fmt.Println("    \\__, |\\___/|_|  |___/\\___|  ")
	fmt.Println("     __/ |                         ")
	fmt.Println("    |___/                          ")
	fmt.Println("                                   ")
	fmt.Println("gorse: Go Recommender System Engine")
	fmt.Println("-----------------------------------")

	//// Connect database
	//var err error
	//db, err = sql.Open("mysql", "gorse:password@/gorse")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Start back-end service
	//back()
	//
	//// Start front-end service
	//serveRPC()
}

var CmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start a recommender sever",
	Run: func(cmd *cobra.Command, args []string) {
		// Load configure file
		configFile, _ := cmd.PersistentFlags().GetString("config")
		var conf TomlConfig
		_, err := toml.DecodeFile(configFile, &conf)
		if err != nil {
			// handle error
			log.Println(err)
		}
		ModelDaemon(conf)
	},
}

func init() {
	CmdServer.PersistentFlags().BoolP("init", "i", false, "Initialize empty database")
	CmdServer.PersistentFlags().String("import-users", "", "import users from CSV file")
	CmdServer.PersistentFlags().String("import-items", "", "import items from CSV file")
	CmdServer.PersistentFlags().String("import-ratings", "", "import ratings from CSV file")
	CmdServer.PersistentFlags().StringP("config", "c", "config.toml", "Configure file")
}
