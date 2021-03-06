package cmd

type Class int8

const (
	Global Class = iota
	Host
	VM
)

var classes = [...]string{
	"global",
	"host",
	"vm",
}

type handler interface{}

type Command struct {
	Name    string
	Class   Class
	handler handler
}

type Commands []*Command

var commands = byName(Commands{
	&Command{
		Name:    "AddVM",
		Class:   Global,
		handler: addVM,
	},
	&Command{
		Name:    "UpdateVM",
		Class:   Global,
		handler: updateVM,
	},
	&Command{
		Name:    "RunVM",
		Class:   Global,
		handler: runVM,
	},
	&Command{
		Name:    "StopVM",
		Class:   VM,
		handler: stopVM,
	},
	&Command{
		Name:    "RemoveVM",
		Class:   Global,
		handler: removeVM,
	},
	&Command{
		Name:    "PingHost",
		Class:   Host,
		handler: pingHost,
	},
	&Command{
		Name:    "AddHostToCluster",
		Class:   Host,
		handler: addHostToCluster,
	},
	&Command{
		Name:    "RemoveHostFromCluster",
		Class:   Host,
		handler: removeHostFromCluster,
	},
})

func byName(commands Commands) map[string]*Command {
	m := make(map[string]*Command)
	for _, cmd := range commands {
		m[cmd.Name] = cmd
	}
	return m
}
