package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func badge(cx *cli.Context) (err error) {
	fmt.Println("Starting up badge")
	return
}

func beforeStart(cx *cli.Context) (err error) {

	gocrud, err = NewGocrud(
		WithCustomLogger(
			cx.String("log.level"),
			cx.String("log.encoding"),
			cx.String("log.output"),
		),
		WithHTTPTransport(
			cx.String("http.host"),
			cx.String("http.port"),
			cx.StringSlice("http.monitor"),
		),
		WithCrudControlPlane(),
	)
	if err != nil {
		return cli.Exit(
			fmt.Sprintf(
				"-- \nfailed to initialize Gocrud. \n--\n Caused By:\n%s\n--",
				errorstack(err.Error()),
			), 9,
		)
	}
	return
}

func actionStart(cx *cli.Context) (err error) {

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Startup Flags")

	return gocrud.Open(cx.Context)
}

func errorstack(errorstr string) string {
	parts := strings.Split(errorstr, ": ")

	var buff bytes.Buffer

	for ix, p := range parts {
		buff.WriteRune('\n')

		for i := 0; i <= ix; i++ {
			buff.WriteRune(' ')
		}

		buff.WriteString("> ")
		buff.WriteString(p)
		if ix > 3 {
			break
		}
	}

	for i := 4; i < len(parts); i++ {
		buff.WriteString(parts[i])
		buff.WriteString(": ")
	}

	return buff.String()
}
