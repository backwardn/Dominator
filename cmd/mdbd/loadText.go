package main

import (
	"bufio"
	"github.com/Symantec/Dominator/lib/mdb"
	"io"
	"log"
	"strings"
)

func loadText(reader io.Reader, logger *log.Logger) *mdb.Mdb {
	scanner := bufio.NewScanner(reader)
	var newMdb mdb.Mdb
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			var machine mdb.Machine
			machine.Hostname = fields[0]
			if len(fields) > 1 {
				machine.RequiredImage = fields[1]
				if len(fields) > 2 {
					machine.PlannedImage = fields[2]
				}
			}
			newMdb.Machines = append(newMdb.Machines, machine)
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Println(err)
		return nil
	}
	return &newMdb
}