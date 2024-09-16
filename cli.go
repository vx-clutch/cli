// Copyright (C) 2024-2024 vx-clutch
//
// This file is part of vx-clutch/cli.
//
// vx-clutch/cli is free software; you can redistribute it and/or modify it under the terms of the BSD 3-Clause Licence.
package cli

import (
	"fmt"
	"os"
)

const (
	reset  = "\033[0m"
	errC   = "\033[1;91m" // Bold High Intensity Red
	warnC  = "\033[1;95m" // Bold High Intensity Purple
	debugC = "\033[1;97m" // Bold High Intensity White
	mname  = ""
)

func Setup(name string) {
	mname = name
}

func Fatal(m any) {
	fmt.Printf("%s: %sfatal error%s: %s", mname, errC, reset, m)
	fmt.Println("\ncompilation terminated.")
	os.Exit(0)
}

func Fatalf(m any, args ...string) string {
	fmt.Printf("%s: %serror%s: %s\n", mname, errC, reset, fmt.Sprintf(m, ...args))
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, m)
}

func Error(m any) string {
	fmt.Printf("%s: %serror%s: %s\n", mname, errC, reset, m)
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, m)
}

func Warn(warning any) string {
	fmt.Printf("%swarning:%s %s\n", warnC, reset, warning)
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, warning)
}

func Debug(tag string, dbm any) string {
	fmt.Printf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
	return fmt.Sprintf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
}
