package cmd

import (
	"errors"
	"traffic-generator/controllers"

	"github.com/spf13/cobra"
)

var url string
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
	Example:               `rungame http://localhost:50000 --gamename "1;random|2;maximo" --players 30 --rungames 30000 --concurrence 10 --timeout 3`,
	Args: func(cmd *cobra.Command, args []string) error {
		if url == "" && len(args) < 1 {
			return errors.New("acepta 1 arg(s)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var argument string
		if url != "" {
			argument = url
		} else {
			argument = args[0]
		}

		controllers.Rungame(argument, gamename, players, rungames, concurrence, timeout)
	},
}

func init() {
	rootCmd.AddCommand(rungameCmd)
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
