package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GenerateCode(cmd *cobra.Command) {
	fmt.Print(`package docker

import (
	"fmt"
	"os/exec"
)
	` + "\n")

	generateCode(cmd, nil)
	os.Exit(0)
}

func generateCode(cmd *cobra.Command, parents []string) {
	var names []string
	if parents == nil {
		names = []string{cmd.Name()}
	} else {
		names = append(parents, cmd.Name())
	}

	fmt.Print(generateStructDefinition(cmd, names))
	for _, cmd := range cmd.Commands() {
		generateCode(cmd, names)
	}
}

func generateStructDefinition(cmd *cobra.Command, names []string) string {
	cmdName := generateCmdName(names)

	result := "type " + cmdName + "Option struct {\n"

	type fieldValue struct{ flag, name, typ string }

	fields := ""
	fieldValues := []fieldValue{}
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fields += "	/*\n	" + flag.Usage + "\n	*/\n"
		name := generateFieldName(flag)
		typ := generateTypeName(flag)
		fields += "	" + name +
			" " + typ + "\n\n"
		fieldValues = append(fieldValues, fieldValue{flag.Name, name, typ})
	})
	result += fields + "}\n\n"

	if fields == "" {
		result = ""
	}

	arg := ""
	if fields != "" {
		arg = "opt " + cmdName + "Option, "
	}
	arg += "args []string"

	result += "/*\n" +
		cmdName + "Cmd is wrapper of '" + strings.Join(names, " ") + "'\n" +
		"------------------------------\n" +
		cmd.Use + "\n" +
		cmd.Short + "\n" +
		"------------------------------\n" +
		"*/\n"
	result += "func " + cmdName + "Cmd" + "(" + arg + ") *exec.Cmd {\n"
	result += "	cargs := []string{"
	for i := range names {
		if i == 0 {
			continue
		}

		result += "\"" + names[i] + "\""

		if i != len(names)-1 {
			result += ", "
		}
	}
	result += "}\n"

	for _, v := range fieldValues {
		result += "	if opt." + v.name + " != nil {\n"
		switch v.typ {
		case "[]string":
			result += "	for _, str := range opt." + v.name + " {\n"
			result += "		cargs = append(cargs, \"--" + v.flag + "\")\n"
			result += "		cargs = append(cargs, str)\n"
			result += "	}\n"
		case "map[string]string":
			result += "	for key, val := range opt." + v.name + " {\n"
			result += "		cargs = append(cargs, \"--" + v.flag + "\")\n"
			result += "		cargs = append(cargs, key + \"=\" + val)\n"
			result += "	}\n"
		default:
			result += "		cargs = append(cargs, \"--" + v.flag + "=\" + fmt.Sprint(*opt." + v.name + "))\n"
		}
		result += "	}\n"
	}

	result += "	cargs = append(cargs, args...)\n"
	result += "	return exec.Command(\"docker\", cargs...)\n}\n\n"

	return result
}

func generateFieldName(flag *pflag.Flag) string {
	name := flag.Name
	var s []string
	if !strings.Contains(name, "-") {
		s = []string{name}
	} else {
		s = strings.Split(name, "-")
	}

	return toCamel(s)
}

func generateTypeName(flag *pflag.Flag) string {
	typ := flag.Value.Type()
	switch typ {
	case "list":
		return "[]string"
	case "map":
		return "map[string]string"
	default:
		if !isBasicType(typ) {
			return "*string"
		}

		return "*" + typ
	}
}

func isBasicType(typ string) bool {
	types := []string{
		"bool",
		"string",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"byte",
		"rune",
		"float32", "float64",
		"complex64", "complex128",
	}

	for _, v := range types {
		if v == typ {
			return true
		}
	}

	return false
}

func generateCmdName(names []string) string {
	s := []string{}

	for _, v := range names {
		if strings.Contains(v, "-") {
			s = append(s, strings.Split(v, "-")...)
			continue
		}

		s = append(s, v)
	}

	return toCamel(s)
}

func toCamel(s []string) string {
	res := ""
	for _, v := range s {
		upper := strings.ToUpper(string(v[0]))
		if len(v) < 2 {
			res += upper
		}
		res += upper + v[1:]
	}
	return res
}
