/*
Copyright © 2022 Jorge Romero jorge@daduic.com
*/
package cmd

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
// Executes the PocketBase web application
// FIXME: Make this a proxy command that yields control to the PocketBase bin, maybe like app.RootCmd.SetArgs(os.Args[2:])
// FIXME: Alternatively, make shelf's daemon a custom console command https://pocketbase.io/docs/use-as-framework/. For example:
// Registering custom console commands, eg.:
// app.RootCmd.AddCommand(&cobra.Command{
//     Use: "hello",
//     Run: func(command *cobra.Command, args []string) {
//         print("Hello world!")
//     },
// })

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "PocketBase CLI",
	Run: func(cmd *cobra.Command, args []string) {
		app := pocketbase.New()

		// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// 	// add new "GET /api/hello" route to the app router (echo)
		// 	e.Router.AddRoute(echo.Route{
		// 		Method: http.MethodGet,
		// 		Path:   "/api/hello",
		// 		Handler: func(c echo.Context) error {
		// 			return c.String(200, "Hello world!")
		// 		},
		// 		Middlewares: []echo.MiddlewareFunc{
		// 			apis.RequireAdminOrUserAuth(),
		// 		},
		// 	})

		// 	return nil
		// })

		if err := app.Start(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
