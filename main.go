package main

import (
	"flag"
	"strings"

	hipache "github.com/catalyst-zero/hipache-config-go"
)

func main() {
	redisHost := flag.String("h", "127.0.0.1", "Set the redis host which holds hipache route configs. i.e. 127.0.0.1")
	redisPort := flag.String("p", "6379", "Set the redis port which holds hipache route configs. i.e. 6739")
	add := flag.Bool("add", false, "Set the command to add mode")
	rm := flag.Bool("rm", false, "Set the command to rm mode")
	frontend := flag.String("f", "", "Set the frontend to add backend for. i.e. http://ruqqq.sg")
	backend := flag.String("b", "", "Set the backend to add to the frontend for. i.e. 192.168.1.1:10000,192.168.1.2:10000")
	flag.Parse()

	cmd := ""

	if *add {
		cmd = "add"
	}

	if *rm {
		cmd = "rm"
	}

	if cmd == "" {
		println("No command specified")
		return
	}

	if *frontend == "" {
		println("frontend not specified")
		return
	}

	if cmd == "add" && *backend == "" {
		println("backend not specified")
		return
	}

	client, err := hipache.DialHipacheConfig(*redisHost + ":" + *redisPort)
	if err != nil {
		panic(err)
	}

	var binding hipache.Binding
	binding, err = client.BindingGet(*frontend)
	if err != nil {
		if _, ok := err.(*hipache.BindingNotFoundError); ok {
			client.BindingCreate(*frontend)
			binding = hipache.Binding{DomainName: *frontend, Hosts: []string{}}
		} else {
			panic(err)
		}
	}

	if *backend != "" {
		backends := strings.Split(*backend, ",")

		for _, backend := range backends {
			found := false

			for _, val := range binding.Hosts {
				if backend == val {
					found = true
				}
			}

			if cmd == "add" && !found {
				if err := client.BindingAddHost(*frontend, backend); err != nil {
					panic(err)
				}
			}

			if cmd == "rm" && found {
				if err := client.BindingRemoveHost(*frontend, backend); err != nil {
					panic(err)
				}
			}
		}
	} else if cmd == "rm" {
		if err := client.BindingDelete(*frontend); err != nil {
			panic(err)
		}
	}

	if binding, err := client.BindingGet(*frontend); err == nil {
		if len(binding.Hosts) == 0 {
			client.BindingDelete(*frontend)
		}
	}

	println("Success.")
}
