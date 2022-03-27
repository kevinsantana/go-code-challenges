package cmd

import (
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/api/http"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg app.Config
		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
		app.Init(cfg)
		http.Run(cfg)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	viper.BindEnv("PORT")
	viper.BindEnv("HOST")
	viper.BindEnv("ENV")
}
