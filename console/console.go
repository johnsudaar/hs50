package console

import (
	"strings"

	"github.com/johnsudaar/hs50/config"
	errgo "gopkg.in/errgo.v1"
)

// Contains the console globals command

// PackCommand Generate Console Commands (Wrap it with the correct parameters)
func PackCommand(command string, params ...string) []byte {
	// TODO: Optimise this function, the query should be bypassable and the function could run 2 times faster
	query := command
	for i := 0; i < len(params); i++ {
		query += params[i]
	}

	rawQuery := make([]byte, len(query)+2)
	rawQuery[0] = config.START_TRANSMISSION
	rawQuery[len(rawQuery)-1] = config.END_TRANSMISSION
	for i := 0; i < len(query); i++ {
		rawQuery[i+1] = query[i]
	}
	return rawQuery
}

// UnPackCommand take a correctly formed command and unpack the command and the params
func UnPackCommand(query []byte) (string, []string, error) {
	if query[0] != config.START_TRANSMISSION {
		return "", nil, errgo.New("NO START_TRANSMISSION packet found")
	}

	if query[len(query)-1] != config.END_TRANSMISSION {
		return "", nil, errgo.New("NO END_TRANSMISSION packet found")
	}

	rawCommand := string(query[1 : len(query)-1])

	parts := strings.Split(rawCommand, ":")
	command := parts[0]

	if len(command) != 4 {
		return "", nil, errgo.New("Invalid command received")
	}

	params := make([]string, len(parts)-1)

	for i := 1; i < len(parts); i++ {
		if len(parts[i]) != 2 {
			return "", nil, errgo.New("Invalid sub command")
		}
		params[i-1] = ":" + parts[i]
	}

	return command, params, nil
}
