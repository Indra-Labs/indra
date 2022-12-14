package cmds

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/indra-labs/indra/pkg/opts/config"
	"github.com/indra-labs/indra/pkg/path"
)

func TestCommand_Foreach(t *testing.T) {
	cm, _ := Init(GetExampleCommands(), nil)
	log.I.Ln("spewing only droptxindex")
	cm.ForEach(func(cmd *Command, _ int) bool {
		if cmd.Name == "droptxindex" {
			log.I.S(cmd)
		}
		return true
	}, 0, 0, cm)
	log.I.Ln("printing name of all commands found on search")
	cm.ForEach(func(cmd *Command, depth int) bool {
		log.I.F("%s%s #(%d)", path.GetIndent(depth), cmd.Path, depth)
		for i := range cmd.Configs {
			log.I.F("%s%s -%s %v #%v (%d)", path.GetIndent(depth),
				cmd.Configs[i].Path(), i, cmd.Configs[i].String(), cmd.Configs[i].Meta().Aliases(), depth)
		}
		return true
	}, 0, 0, cm)
}

func TestCommand_MarshalText(t *testing.T) {
	o, _ := Init(GetExampleCommands(), nil)
	conf, err := o.MarshalText()
	if log.E.Chk(err) {
		t.FailNow()
	}
	log.I.Ln("\n" + string(conf))
}

func TestCommand_UnmarshalText(t *testing.T) {
	o, _ := Init(GetExampleCommands(), nil)
	var conf []byte
	var err error
	conf, err = o.MarshalText()
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = o.UnmarshalText(conf)
	if err != nil {
		t.FailNow()
	}
}

func TestCommand_GetEnvs(t *testing.T) {
	o, _ := Init(GetExampleCommands(), nil)
	envs := o.GetEnvs()
	var out []string
	err := envs.ForEach(func(env string, opt config.Option) error {
		out = append(out, env)
		return nil
	})
	for i := range out { // verifying ordering groups subcommands
		log.I.Ln(out[i])
	}
	if err != nil {
		t.FailNow()
	}
}

var testSeparator = fmt.Sprintf("%s\n", strings.Repeat("-", 72))

func TestCommand_Help(t *testing.T) {
	ex := GetExampleCommands()
	ex.AddCommand(Help())
	o, _ := Init(ex, nil)
	o.Commands = append(o.Commands)
	args1 := "/random/path/to/server_binary help"
	fmt.Println(args1)
	args1s := strings.Split(args1, " ")
	run, args, err := o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help loglevel"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help help"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help node"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help rpcconnect"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help kopach rpcconnect"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help node rpcconnect"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help nodeoff"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help user"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}
	fmt.Print(testSeparator)
	args1 = "/random/path/to/server_binary help file"
	fmt.Println(args1)
	args1s = strings.Split(args1, " ")
	run, args, err = o.ParseCLIArgs(args1s)
	if log.E.Chk(err) {
		t.FailNow()
	}
	err = run.Entrypoint(o, args)
	if log.E.Chk(err) {
		t.FailNow()
	}

}
func TestCommand_LogToFile(t *testing.T) {
	ex := GetExampleCommands()
	ex.AddCommand(Help())
	ex, _ = Init(ex, nil)
	ex.GetOpt(path.From("pod123 loglevel")).FromString("debug")
	var err error
	// this will create a place we can write the logs
	if err = ex.SaveConfig(); log.E.Chk(err) {
		err = os.RemoveAll(ex.Configs["ConfigFile"].Expanded())
		if log.E.Chk(err) {
		}
		t.FailNow()
	}
	lfp := ex.GetOpt(path.From("pod123 logfilepath"))
	o := ex.GetOpt(path.From("pod123 logtofile"))
	o.FromString("true")
	log.I.F("%s", lfp)
	o.FromString("false")
	var b []byte
	if b, err = os.ReadFile(lfp.Expanded()); log.E.Chk(err) {
		t.FailNow()
	}
	str := string(b)
	log.I.F("'%s'", str)
	if !strings.Contains(str, lfp.String()) {
		t.FailNow()
	}
	if err := os.RemoveAll(ex.Configs["DataDir"].Expanded()); log.E.Chk(err) {
	}
}
