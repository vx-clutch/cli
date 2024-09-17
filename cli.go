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
)

var (
	mname = "call Setup()"
)

// Gives project name to program, MUST be called first
func Setup(name string) {
	mname = name
}

// Fatal call exits with zero, takes a fatal error message
func Fatal(m string) {
	fmt.Printf("%s: %sfatal error%s: %s", mname, errC, reset, m)
	fmt.Println("\ncompilation terminated.")
	os.Exit(0)
}

// Fatal call exits with zero, Takes a error message and a formatted string
func Fatalf(m string, args ...interface{}) {
	fmt.Printf("%s: %serror%s: %s\n", mname, errC, reset, fmt.Sprintf(m, args...))
	fmt.Println("compilation terminated.")
	os.Exit(0)
}

// Error call does not exit, Takes a error message
func Error(m string) string {
	fmt.Printf("%s: %serror%s: %s\n", mname, errC, reset, m)
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, m)
}

// Error call does not exit, Takes a error message and a formatted string
func Errorf(m string, args ...interface{}) string {
	fmt.Printf("%s: %serror%s: %s\n", mname, errC, reset, fmt.Sprintf(m, args...))
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, fmt.Sprintf(m, args...))
}

// Warn call does not exit, Takes a warning message
func Warn(warning string) string {
	fmt.Printf("%swarning:%s %s\n", warnC, reset, warning)
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, warning)
}

// Warn call does not exit, Takes a warning message and a formatted string
func Warnf(warning string, args ...interface{}) string {
	fmt.Printf("%swarning:%s %s\n", warnC, reset, fmt.Sprintf(warning, args...))
	return fmt.Sprintf("%s: %serror%s: %s\n", mname, errC, reset, warning)
}

// Debug call does not exit. can be passed a tag and a debug message
func Debug(tag string, dbm any) string {
	fmt.Printf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
	return fmt.Sprintf("%sdebug%s: %s\n%s\n", debugC, reset, tag, dbm)
}
