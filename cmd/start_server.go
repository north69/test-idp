package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"test-idp/internal/config"
	"test-idp/internal/http"
)

func startServerCmd(cnf *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "start-server",
		Short: "start-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			s := echo.New()
			h := http.NewHandler(s, cnf)
			h.InitRoutes()

			errCh := make(chan error)

			go func() {
				if err := s.Start("localhost:" + cnf.App.HttpPort); err != nil {
					errCh <- err
				}
			}()
			select {
			case <-cmd.Context().Done():
				return nil
			case err := <-errCh:
				return err
			}
		},
	}
}
