package main

import (
	"bufio"
	"fmt"
	"github.com/Symantec/Dominator/lib/objectcache"
	"github.com/Symantec/Dominator/proto/sub"
	"net/rpc"
	"os"
)

func fetchSubcommand(client *rpc.Client, args []string) {
	if err := fetch(client, args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching\t%s\n", err)
		os.Exit(2)
	}
	os.Exit(0)
}

func fetch(client *rpc.Client, hashesFilename string) error {
	hashesFile, err := os.Open(hashesFilename)
	if err != nil {
		return err
	}
	defer hashesFile.Close()
	scanner := bufio.NewScanner(hashesFile)
	var request sub.FetchRequest
	var reply sub.FetchResponse
	request.ServerAddress = fmt.Sprintf("%s:%d",
		*objectServerHostname, *objectServerPortNum)
	for scanner.Scan() {
		hashval, err := objectcache.FilenameToHash(scanner.Text())
		if err != nil {
			return err
		}
		request.Hashes = append(request.Hashes, hashval)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return client.Call("Subd.Fetch", request, &reply)
}