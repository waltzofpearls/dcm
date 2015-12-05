package main

import (
	"fmt"
	"os"
)

type Dcm struct {
	Config *Config
	Args   []string
}

func NewDcm(c *Config, args []string) *Dcm {
	return &Dcm{c, args}
}

func (d *Dcm) Command() int {
	if len(d.Args) < 1 {
		d.Usage()
		return 1
	}

	switch d.Args[0] {
	case "help", "h":
		return d.Usage()
	case "setup":
		return d.Setup()
	case "run", "r":
		return d.Run()
	case "build", "b":
		return d.Build()
	case "update", "u":
		return d.Update()
	case "shell", "sh":
		return d.Shell()
	case "goto", "gt":
		return d.Goto()
	case "purge", "rm":
		return d.Purge()
	case "unload", "ul":
		return d.Unload()
	default:
		d.Usage()
		return 127
	}

	return 0
}

func (d *Dcm) Setup() int {
	if _, err := os.Stat(d.Config.Srv); os.IsNotExist(err) {
		os.MkdirAll(d.Config.Srv, 0777)
	}

	for service, configs := range d.Config.Config {
		service, _ := service.(string)
		configs, ok := configs.(map[interface{}]interface{})
		if !ok {
			panic("Error reading git repository config for service: " + service)
		}
		repo, _ := getMapVal(configs, "labels", "com.dcm.repository").(string)
		dir := d.Config.Srv + "/" + service
		if err := cmd("git", "clone", repo, dir).Run(); err != nil {
			panic("Error cloning git repository for service: " + service)
		}
	}

	return 0
}

func (d *Dcm) Run(args ...string) int {
	return 0
}

func (d *Dcm) Build() int {
	return 0
}

func (d *Dcm) Update() int {
	return 0
}

func (d *Dcm) Shell() int {
	return 0
}

func (d *Dcm) Goto() int {
	return 0
}

func (d *Dcm) Purge() int {
	return 0
}

func (d *Dcm) Unload() int {
	return 0
}

func (d *Dcm) Usage() int {
	fmt.Println("")
	fmt.Println("Docker Compose Manager")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  dcm help                Show this message")
	fmt.Println("                          shorthand ver.: `dcm h`")
	fmt.Println("  dcm setup               Check out all the repositories for API, UI and services")
	fmt.Println("  dcm run [<args>]        Run docker-compose commands. If <args> is not given, by")
	fmt.Println("                          default DCM will run `up` command.")
	fmt.Println("                          <args>: up, build, start, stop, restart")
	fmt.Println("                          shorthand ver.: `dcm r [<args>]`")
	fmt.Println("  dcm run build           Run build command that (re)create all the service images")
	fmt.Println("                          shorthand ver.: `dcm build` or `dcm b`")
	fmt.Println("  dcm shell <service>     Log into a given service container")
	fmt.Println("                          <service>: api, ui, postgres, mongo, redis, nginx, php")
	fmt.Println("                          shorthand ver.: `dcm sh <service>`")
	fmt.Println("  dcm purge [<type>]      Remove either all the containers or all the images or")
	fmt.Println("                          everything. If <type> is not given, by default DCM will")
	fmt.Println("                          purge everything")
	fmt.Println("                          <type>: images|img, containers|con, all")
	fmt.Println("                          shorthand ver.: `dcm rm [<type>]`")
	fmt.Println("  dcm branch <service>    Display the current branch for the given service name")
	fmt.Println("                          <service>: api, ui, postgres, mongo, redis, nginx, php")
	fmt.Println("                          shorthand ver.: `dcm br <service>`")
	fmt.Println("  dcm goto [<service>]    Go to the service's folder. If <service> is not given,")
	fmt.Println("                          by default DCM will go to $DCM_DIR")
	fmt.Println("                          <service>: api, ui, postgres, mongo, redis, nginx, php")
	fmt.Println("                          shorthand ver.: `dcm gt [<service>]`")
	fmt.Println("  dcm update [<service>]  Update DCM and dependent services (PostgrSQL, MongoDB,")
	fmt.Println("                          Redis, Nginx and Base PHP). If <service> is not given,")
	fmt.Println("                          by default DCM will update everything except api and ui.")
	fmt.Println("                          <service>: postgres, mongo, redis, nginx, php")
	fmt.Println("                          shorthand ver.: `dcm u`")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  Initial setup")
	fmt.Println("    dcm setup")
	fmt.Println("    dcm run")
	fmt.Println("")
	fmt.Println("  Rebuild API or UI after switching branch")
	fmt.Println("    dcm build")
	fmt.Println("    dcm run")
	fmt.Println("")
	fmt.Println("  Log into different service containers")
	fmt.Println("    dcm shell api")
	fmt.Println("    dcm shell ui")
	fmt.Println("    ...")
	fmt.Println("")

	return 0
}
