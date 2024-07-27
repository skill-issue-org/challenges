package main

/**
 * Skill Issue Weekly Challenges
 *
 * Title: Ibu Kota Negara (Mangcrack)
 * Points: 8
 * Author: dwisiswant0
**/

import (
	"fmt"

	"skill-issue.org/ibu-kota-negara/pkg/consts"
	"skill-issue.org/ibu-kota-negara/pkg/errors"
	"skill-issue.org/ibu-kota-negara/pkg/flag"
	"skill-issue.org/ibu-kota-negara/pkg/vars"
)

var password string

func main() {
	fmt.Println(consts.Banner)
	fmt.Printf("%s: ", consts.Input)
	fmt.Scan(&password)

	println()

	if f := auth(password); f != "" {
		fmt.Printf("%s: %s", consts.Result, f)

		return
	}

	panic(errors.ErrPassword)
}

func auth(password string) string {
	if password == vars.Password {
		return flag.Get()
	}

	return ""
}
