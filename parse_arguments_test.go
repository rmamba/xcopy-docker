package main

import (
	"testing"
)

func CheckBool(t *testing.T, input bool, expected bool) {
	if input != expected {
		t.Errorf("Value: %t; expected: %t", input, expected)
	}
}

func CheckInt(t *testing.T, input int, expected int) {
	if input != expected {
		t.Errorf("Value: %d; expected: %d", input, expected)
	}
}

func CheckStr(t *testing.T, input string, expected string) {
	if input != expected {
		t.Errorf("Value: %s; expected: %s", input, expected)
	}
}

func TestEmptyInput(t *testing.T) {
	_, errCode := ParseArguments([]string{})
	CheckInt(t, errCode, -1)
}

func TestSingleInput(t *testing.T) {
	_, errCode := ParseArguments([]string{"test"})
	CheckInt(t, errCode, -1)
}

func TestFileCopy(t *testing.T) {
	args, errCode := ParseArguments([]string{"go.mod", "go.mod.copied_by_test"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestFileCopyToFolder(t *testing.T) {
	args, errCode := ParseArguments([]string{"go.mod", "/tmp"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "/tmp")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, true)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestFileCopyToFolderRenamed(t *testing.T) {
	args, errCode := ParseArguments([]string{"go.mod", "/tmp/go.mod.copied_by_test"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "/tmp/go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParametersWithQuotes(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParametersWithoutSource(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D"})
	CheckInt(t, errCode, -1)

	CheckStr(t, args.Source, "")
	CheckStr(t, args.Dest, "")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, true)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, true)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParametersWithoutDest(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "\"go.mod\""})
	CheckInt(t, errCode, -1)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, true)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterA(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"go.mod\"", "/A", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, true)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterACD(t *testing.T) {
	args, errCode := ParseArguments([]string{"/C", "\"go.mod\"", "/A", "\"go.mod.copied_by_test\"", "/D"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, true)
	CheckBool(t, args.Flags.C, true)
	CheckBool(t, args.Flags.D, true)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterEFH(t *testing.T) {
	args, errCode := ParseArguments([]string{"/E", "\"go.mod\"", "\"go.mod.copied_by_test\"", "/F", "/H"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, true)
	CheckBool(t, args.Flags.F, true)
	CheckBool(t, args.Flags.H, true)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterILMNP(t *testing.T) {
	args, errCode := ParseArguments([]string{"/P", "/I", "/N", "\"go.mod\"", "\"go.mod.copied_by_test\"", "/L", "/M"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, true)
	CheckBool(t, args.Flags.L, true)
	CheckBool(t, args.Flags.M, true)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, true)
	CheckBool(t, args.Flags.P, true)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterQRST(t *testing.T) {
	args, errCode := ParseArguments([]string{"/T", "/Q", "/S", "/R", "\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, true)
	CheckBool(t, args.Flags.R, true)
	CheckBool(t, args.Flags.S, true)
	CheckBool(t, args.Flags.T, true)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterVWY(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"go.mod\"", "\"go.mod.copied_by_test\"", "/W", "/Y", "/V"})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, true)
	CheckBool(t, args.Flags.W, true)
	CheckBool(t, args.Flags.Y, true)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterMinusY(t *testing.T) {
	args, errCode := ParseArguments([]string{"/-Y", "\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, false)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, true)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "")
}

func TestParameterDate(t *testing.T) {
	args, errCode := ParseArguments([]string{"/D031019", "\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)

	CheckStr(t, args.Source, "go.mod")
	CheckStr(t, args.Dest, "go.mod.copied_by_test")

	CheckBool(t, args.IsSourceDir, false)
	CheckBool(t, args.IsDestDir, false)

	CheckBool(t, args.Flags.A, false)
	CheckBool(t, args.Flags.C, false)
	CheckBool(t, args.Flags.D, true)
	CheckBool(t, args.Flags.E, false)
	CheckBool(t, args.Flags.F, false)
	CheckBool(t, args.Flags.H, false)
	CheckBool(t, args.Flags.I, false)
	CheckBool(t, args.Flags.L, false)
	CheckBool(t, args.Flags.M, false)
	CheckBool(t, args.Flags.MinusY, false)
	CheckBool(t, args.Flags.N, false)
	CheckBool(t, args.Flags.P, false)
	CheckBool(t, args.Flags.Q, false)
	CheckBool(t, args.Flags.R, false)
	CheckBool(t, args.Flags.S, false)
	CheckBool(t, args.Flags.T, false)
	CheckBool(t, args.Flags.V, false)
	CheckBool(t, args.Flags.W, false)
	CheckBool(t, args.Flags.Y, false)
	CheckStr(t, args.Flags.ParamsD, "031019")
}

// // TestHelloEmpty calls greetings.Hello with an empty string,
// // checking for an error.
// func TestHelloEmpty(t *testing.T) {
// 	msg, err := Hello("")
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
// 	}
// }
