package cmd

import (
	"errors"
	"strings"
	"traffic-generator/controllers"

	"github.com/spf13/cobra"
)

var ip string
var urlNipIo string
var isHttps bool
var gamename string
var players int
var rungames int
var concurrence int
var timeout int

// rungameCmd represents the rungame command
var rungameCmd = &cobra.Command{
	Use:   "rungame",
	Short: "Genera tráfico con nip.io",
	Long: `Genera tráfico con nip.io especificando los juegos a jugar la cantidad de jugadores
de cada juego, el número de veces para ejecutar los juegos, las solicitudes simultánea
a la API para ejecutar los juegos, y el timeout, si el tiempo restante es mayor que
este valor, el comando se detendrá.`,
	DisableFlagsInUseLine: true,
	Example:               `rungame -g "juego1 juego2" -p 2 -r 2 -c 2 -t 10`,
	Args: func(cmd *cobra.Command, args []string) error {
		if ip == "" && len(args) < 1 {
			return errors.New("acepta 1 arg(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var argument string
		if ip != "" {
			argument = ip
		} else {
			argument = args[0]
		}
		var url strings.Builder
		if isHttps {
			url.WriteString("https://")
		} else {
			url.WriteString("http://")
		}
		url.WriteString(argument)
		url.WriteString(".nip.io")
		urlNipIo = url.String()
		controllers.Rungame(urlNipIo, gamename, players, rungames, concurrence, timeout)
	},
}

func init() {
	rootCmd.AddCommand(rungameCmd)
	rungameCmd.Flags().BoolVarP(&isHttps, "https", "s", false, "Usa https")
	rungameCmd.Flags().StringVarP(&gamename, "gamename", "g", "", "Nombres de los juegos")
	rungameCmd.Flags().IntVarP(&players, "players", "p", 0, "Cantidad de jugadores")
	rungameCmd.Flags().IntVarP(&rungames, "rungames", "r", 0, "Cantidad de veces a ejecutar")
	rungameCmd.Flags().IntVarP(&concurrence, "concurrence", "c", 0, "Cantidad de solicitudes simultáneas")
	rungameCmd.Flags().IntVarP(&timeout, "timeout", "t", 0, "Tiempo de espera")
	rungameCmd.MarkFlagRequired("gamename")
	rungameCmd.MarkFlagRequired("players")
	rungameCmd.MarkFlagRequired("rungames")
	rungameCmd.MarkFlagRequired("concurrence")
	rungameCmd.MarkFlagRequired("timeout")
}
