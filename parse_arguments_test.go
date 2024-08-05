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
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestFileCopyToFolder(t *testing.T) {
	args, errCode := ParseArguments([]string{"go.mod", "/tmp/"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "/tmp/")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, true)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestFileCopyToFolderRenamed(t *testing.T) {
	args, errCode := ParseArguments([]string{"go.mod", "/tmp/go.mod.copied_by_test"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "/tmp/go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParametersWithQuotes(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParametersWithoutSource(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D"})
	CheckInt(t, errCode, -2)
	CheckInt(t, len(args), 0)
}

func TestParametersWithoutDest(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "\"go.mod\""})
	CheckInt(t, errCode, -2)
	CheckInt(t, len(args), 0)
}

func TestParameterA(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"go.mod\"", "/A", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, true)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterACD(t *testing.T) {
	args, errCode := ParseArguments([]string{"/C", "\"go.mod\"", "/A", "\"go.mod.copied_by_test\"", "/D"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, true)
	CheckBool(t, args[0].Flags.C, true)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterEFH(t *testing.T) {
	args, errCode := ParseArguments([]string{"/E", "\"go.mod\"", "\"go.mod.copied_by_test\"", "/F", "/H"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, true)
	CheckBool(t, args[0].Flags.F, true)
	CheckBool(t, args[0].Flags.H, true)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterILMNP(t *testing.T) {
	args, errCode := ParseArguments([]string{"/P", "/I", "/N", "\"go.mod\"", "\"go.mod.copied_by_test\"", "/L", "/M"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, true)
	CheckBool(t, args[0].Flags.L, true)
	CheckBool(t, args[0].Flags.M, true)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, true)
	CheckBool(t, args[0].Flags.P, true)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterQRST(t *testing.T) {
	args, errCode := ParseArguments([]string{"/T", "/Q", "/S", "/R", "\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, true)
	CheckBool(t, args[0].Flags.R, true)
	CheckBool(t, args[0].Flags.S, true)
	CheckBool(t, args[0].Flags.T, true)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterVWY(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"go.mod\"", "\"go.mod.copied_by_test\"", "/W", "/Y", "/V"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, true)
	CheckBool(t, args[0].Flags.W, true)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterMinusY(t *testing.T) {
	args, errCode := ParseArguments([]string{"/-Y", "\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, false)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, true)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestParameterDate(t *testing.T) {
	args, errCode := ParseArguments([]string{"/D031019", "\"go.mod\"", "\"go.mod.copied_by_test\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "go.mod")
	CheckStr(t, args[0].Dest, "go.mod.copied_by_test")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, false)
	CheckStr(t, args[0].Flags.ParamsD, "031019")
}

func TestRealMissingSource(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D", "\"/tmp/unknown.json\"", "\"/tmp/renamed.json\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "/tmp/unknown.json")
	CheckStr(t, args[0].Dest, "/tmp/renamed.json")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestCopyFoldersSuccess(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D", "\"/tmp/\"", "\"/tmp2/\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "/tmp/")
	CheckStr(t, args[0].Dest, "/tmp2/")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, true)
	CheckBool(t, args[0].IsDestDir, true)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestCopyFoldersWithoutQuotesSuccess(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D", "/tmp/", "/tmp2/"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "/tmp/")
	CheckStr(t, args[0].Dest, "/tmp2/")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, true)
	CheckBool(t, args[0].IsDestDir, true)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestCopySuccess(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D", "\"main.go\"", "\"tested_main.go\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "main.go")
	CheckStr(t, args[0].Dest, "tested_main.go")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestCopyToNewLocationSuccess(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D", "\"main.go\"", "\"tests/main.go\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "main.go")
	CheckStr(t, args[0].Dest, "tests/main.go")

	CheckBool(t, args[0].UsesWildcards, false)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, false)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestCopyWithWildcardSuccess(t *testing.T) {
	args, errCode := ParseArguments([]string{"/Y", "/D", "\"main*.go\"", "\"tests/\""})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 1)

	CheckStr(t, args[0].Source, "main*.go")
	CheckStr(t, args[0].Dest, "tests/")

	CheckBool(t, args[0].UsesWildcards, true)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, true)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, false)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")
}

func TestMultilineCommands(t *testing.T) {
	args, errCode := ParseArguments([]string{"\"/tmp/appsettings*.json\"", "/tmp/Dest/\" /S /Y /D\r\nxcopy /tmp2/appsettings*.json /tmp/Dest2/\"", "/S", "/Y", "/D"})
	CheckInt(t, errCode, 0)
	CheckInt(t, len(args), 2)

	// IDX 0

	CheckStr(t, args[0].Source, "/tmp/appsettings*.json")
	CheckStr(t, args[0].Dest, "/tmp/Dest/")

	CheckBool(t, args[0].UsesWildcards, true)
	CheckBool(t, args[0].IsSourceDir, false)
	CheckBool(t, args[0].IsDestDir, true)

	CheckBool(t, args[0].Flags.A, false)
	CheckBool(t, args[0].Flags.C, false)
	CheckBool(t, args[0].Flags.D, true)
	CheckBool(t, args[0].Flags.E, false)
	CheckBool(t, args[0].Flags.F, false)
	CheckBool(t, args[0].Flags.H, false)
	CheckBool(t, args[0].Flags.I, false)
	CheckBool(t, args[0].Flags.L, false)
	CheckBool(t, args[0].Flags.M, false)
	CheckBool(t, args[0].Flags.MinusY, false)
	CheckBool(t, args[0].Flags.N, false)
	CheckBool(t, args[0].Flags.P, false)
	CheckBool(t, args[0].Flags.Q, false)
	CheckBool(t, args[0].Flags.R, false)
	CheckBool(t, args[0].Flags.S, true)
	CheckBool(t, args[0].Flags.T, false)
	CheckBool(t, args[0].Flags.V, false)
	CheckBool(t, args[0].Flags.W, false)
	CheckBool(t, args[0].Flags.Y, true)
	CheckStr(t, args[0].Flags.ParamsD, "")

	// IDX 1

	CheckStr(t, args[1].Source, "/tmp2/appsettings*.json")
	CheckStr(t, args[1].Dest, "/tmp/Dest2/")

	CheckBool(t, args[1].UsesWildcards, true)
	CheckBool(t, args[1].IsSourceDir, false)
	CheckBool(t, args[1].IsDestDir, true)

	CheckBool(t, args[1].Flags.A, false)
	CheckBool(t, args[1].Flags.C, false)
	CheckBool(t, args[1].Flags.D, true)
	CheckBool(t, args[1].Flags.E, false)
	CheckBool(t, args[1].Flags.F, false)
	CheckBool(t, args[1].Flags.H, false)
	CheckBool(t, args[1].Flags.I, false)
	CheckBool(t, args[1].Flags.L, false)
	CheckBool(t, args[1].Flags.M, false)
	CheckBool(t, args[1].Flags.MinusY, false)
	CheckBool(t, args[1].Flags.N, false)
	CheckBool(t, args[1].Flags.P, false)
	CheckBool(t, args[1].Flags.Q, false)
	CheckBool(t, args[1].Flags.R, false)
	CheckBool(t, args[1].Flags.S, true)
	CheckBool(t, args[1].Flags.T, false)
	CheckBool(t, args[1].Flags.V, false)
	CheckBool(t, args[1].Flags.W, false)
	CheckBool(t, args[1].Flags.Y, true)
	CheckStr(t, args[1].Flags.ParamsD, "")
}
