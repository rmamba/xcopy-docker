package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type InputArguments struct {
	Dest          string `default:""`
	Flags         XCopyFlags
	IsDestDir     bool   `default:"false"`
	IsSourceDir   bool   `default:"false"`
	Source        string `default:""`
	UsesWildcards bool   `default:"false"`
}

type XCopyFlags struct {
	A       bool   `default:"false"`
	C       bool   `default:"false"`
	D       bool   `default:"false"`
	E       bool   `default:"false"`
	F       bool   `default:"false"`
	H       bool   `default:"false"`
	I       bool   `default:"false"`
	L       bool   `default:"false"`
	M       bool   `default:"false"`
	N       bool   `default:"false"`
	P       bool   `default:"false"`
	Q       bool   `default:"false"`
	R       bool   `default:"false"`
	S       bool   `default:"false"`
	T       bool   `default:"false"`
	V       bool   `default:"false"`
	W       bool   `default:"false"`
	Y       bool   `default:"false"`
	MinusY  bool   `default:"false"`
	ParamsD string `default:""`
}

func IsDirectory(path string) bool {
	if strings.HasSuffix(path, "/") {
		return true
	}

	fi, err := os.Stat(path)
	if err != nil {
		if os.Getenv("XCOPY_DEBUG") == "true" {
			fmt.Println(err.Error())
		}
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	}
	return false
}

func Choose[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func noEmptyStrings(s string) bool {
	return s != ""
}

// Hello returns a greeting for the named person.
func ParseArguments(args []string) ([]InputArguments, int) {
	arguments := []InputArguments{}
	currentArguments := InputArguments{}

	if len(args) < 2 {
		return arguments, -1
	}

	cleanedArguments := []string{}
	for _, a := range args {
		s := strings.ReplaceAll(a, "\r", "")
		if strings.Contains(s, "\nxcopy") {
			res1 := strings.Split(s, "\nxcopy")
			res2 := Choose(strings.Split(res1[0], " "), noEmptyStrings)
			cleanedArguments = append(cleanedArguments, res2...)
			res3, err := ParseArguments(cleanedArguments)
			if err != 0 {
				return arguments, -5
			}
			arguments = append(arguments, res3...)
			cleanedArguments = Choose(strings.Split(res1[1], " "), noEmptyStrings)
		} else {
			cleanedArguments = append(cleanedArguments, a)
		}
	}

	// Parse arguments
	for i, a := range cleanedArguments {
		if os.Getenv("XCOPY_DEBUG") == "true" {
			fmt.Printf("XCOPY_DEBUG_ARGS[%d]: `%s`\n", i, a)
		}
		if strings.HasPrefix(a, "/") && (len(a) == 2 || len(a) == 3) {
			switch a {
			case "/A":
				currentArguments.Flags.A = true
			case "/C":
				currentArguments.Flags.C = true
			case "/D":
				currentArguments.Flags.D = true
			case "/E":
				currentArguments.Flags.E = true
			case "/F":
				currentArguments.Flags.F = true
			case "/H":
				currentArguments.Flags.H = true
			case "/I":
				currentArguments.Flags.I = true
			case "/L":
				currentArguments.Flags.L = true
			case "/M":
				currentArguments.Flags.M = true
			case "/N":
				currentArguments.Flags.N = true
			case "/P":
				currentArguments.Flags.P = true
			case "/Q":
				currentArguments.Flags.Q = true
			case "/R":
				currentArguments.Flags.R = true
			case "/S":
				currentArguments.Flags.S = true
			case "/T":
				currentArguments.Flags.T = true
			case "/V":
				currentArguments.Flags.V = true
			case "/W":
				currentArguments.Flags.W = true
			case "/Y":
				currentArguments.Flags.Y = true
			case "/-Y":
				currentArguments.Flags.MinusY = true
			}
		} else {
			if strings.HasPrefix(a, "/D") && (len(a) == 8) {
				currentArguments.Flags.D = true
				currentArguments.Flags.ParamsD = a[2:]
			} else {
				if currentArguments.Source == "" {
					currentArguments.Source = strings.ReplaceAll(strings.Trim(a, "\""), "\\", "/")
					currentArguments.UsesWildcards = strings.Contains(currentArguments.Source, "*")
					currentArguments.IsSourceDir = IsDirectory(currentArguments.Source)
				} else {
					currentArguments.Dest = strings.ReplaceAll(strings.Trim(a, "\""), "\\", "/")
					currentArguments.IsDestDir = IsDirectory(currentArguments.Dest)
				}
			}
		}
	}

	if currentArguments.Dest == "" || currentArguments.Source == "" {
		return arguments, -2
	}

	arguments = append(arguments, currentArguments)
	// exitCode := 0
	// arguments.IsSourceDir, exitCode = IsDirectory(arguments.Source)
	// if exitCode != 0 {
	// 	return arguments, exitCode
	// }

	// arguments.IsDestDir, _ = IsDirectory(arguments.Dest)
	return arguments, 0
}

func CopyFile(sourceFileName string, dest string, isDestDir bool) {
	folderStr := ""
	if isDestDir {
		folderStr = "directory "
	}
	if os.Getenv("XCOPY_DEBUG") == "true" {
		fmt.Printf("XCOPY_DEBUG: from `%s` to %s`%s`\n", sourceFileName, folderStr, dest)
	}

	src, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Printf("Error opening source file %s: %v\n", sourceFileName, err)
		os.Exit(1)
	}
	defer src.Close()

	destFileName := dest
	if isDestDir {
		destFileName = filepath.Join(dest, src.Name())
	}
	dst, err := os.Create(destFileName)
	if err != nil {
		fmt.Printf("Error creating destination file %s: %v\n", destFileName, err)
		os.Exit(4)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying data from source to destination")
		os.Exit(4)
	}
}

func main() {
	args, errCode := ParseArguments(os.Args[1:])
	if errCode != 0 {
		fmt.Println("Go implementation of MS-DOS xcopy command v0.0.7")
		fmt.Println("xcopy source [destination]") //[/switches]
		fmt.Println("  source       Specifies the directory and/or name of file(s) to copy.")
		fmt.Println("  destination  Specifies the location and/or name of new file(s).")
		// fmt.Println("  /A           Copies only files with the archive attribute set and doesn't change the attribute.")
		// fmt.Println("  /C           Continues copying even if errors occur.")
		// fmt.Println("  /D[:M/D/Y]   Copies only files which have been changed on or after the specified date. When no date is specified, only files which are newer than existing destination files will be copied.")
		// fmt.Println("  /E           Copies any subdirectories, even if empty.")
		// fmt.Println("  /F           Display full source and destination name.")
		// fmt.Println("  /H           Copies hidden and system files as well as unprotected files.")
		// fmt.Println("  /I           If destination does not exist and copying more than one file, assume destination is a directory.")
		// fmt.Println("  /L           List files without copying them. (simulates copying)")
		// fmt.Println("  /M           Copies only files with the archive attribute set and turns off the archive attribute of the source files after copying them.")
		// fmt.Println("  /N           Suppresses prompting to confirm you want to overwrite an existing destination file and skips these files.")
		// fmt.Println("  /P           Prompts for confirmation before creating each destination file.")
		// fmt.Println("  /Q           Quiet mode, don't show copied filenames.")
		// fmt.Println("  /R           Overwrite read-only files as well as unprotected files.")
		// fmt.Println("  /S           Copies directories and subdirectories except empty ones.")
		// fmt.Println("  /T           Creates directory tree without copying files. Empty directories will not be copied. To copy them add switch /E.")
		// fmt.Println("  /V           Verifies each new file.")
		// fmt.Println("  /W           Waits for a keypress before beginning.")
		// fmt.Println("  /Y           Suppresses prompting to confirm you want to overwrite an existing destination file and overwrites these files.")
		// fmt.Println("  /-Y          Causes prompting to confirm you want to overwrite an existing destination file.")
		if os.Getenv("XCOPY_DEBUG") == "true" {
			fmt.Println("DEBUG:")
			for i, a := range args {
				fmt.Printf("      SOURCE[%d]: `%s`\n", i, a.Source)
				fmt.Printf("        DEST[%d]: `%s`\n", i, a.Dest)
				fmt.Printf("  SOURCE_DIR[%d]: `%t`\n", i, a.IsSourceDir)
				fmt.Printf("    DEST_DIR[%d]: `%t`\n", i, a.IsDestDir)
				fmt.Printf("           *[%d]: `%t`\n", i, a.UsesWildcards)
				fmt.Printf("          /A[%d]: `%t`\n", i, a.Flags.A)
				fmt.Printf("          /C[%d]: `%t`\n", i, a.Flags.C)
				fmt.Printf("          /D[%d]: `%t`\n", i, a.Flags.D)
				fmt.Printf("           D[%d]: `%s`\n", i, a.Flags.ParamsD)
				fmt.Printf("          /E[%d]: `%t`\n", i, a.Flags.E)
				fmt.Printf("          /F[%d]: `%t`\n", i, a.Flags.F)
				fmt.Printf("          /H[%d]: `%t`\n", i, a.Flags.H)
				fmt.Printf("          /I[%d]: `%t`\n", i, a.Flags.I)
				fmt.Printf("          /L[%d]: `%t`\n", i, a.Flags.L)
				fmt.Printf("          /M[%d]: `%t`\n", i, a.Flags.M)
				fmt.Printf("          /N[%d]: `%t`\n", i, a.Flags.N)
				fmt.Printf("          /P[%d]: `%t`\n", i, a.Flags.P)
				fmt.Printf("          /Q[%d]: `%t`\n", i, a.Flags.Q)
				fmt.Printf("          /R[%d]: `%t`\n", i, a.Flags.R)
				fmt.Printf("          /S[%d]: `%t`\n", i, a.Flags.S)
				fmt.Printf("          /T[%d]: `%t`\n", i, a.Flags.T)
				fmt.Printf("          /V[%d]: `%t`\n", i, a.Flags.V)
				fmt.Printf("          /W[%d]: `%t`\n", i, a.Flags.W)
				fmt.Printf("          /Y[%d]: `%t`\n", i, a.Flags.Y)
				fmt.Printf("         /-Y[%d]: `%t`\n", i, a.Flags.MinusY)
			}
		}
		os.Exit(errCode)
	}

	for _, arg := range args {
		if arg.IsSourceDir {
			os.Exit(13)
			// files, err = os.ReadDir(source)
			// if err != nil {
			// 	fmt.Println("Error reading source directory")
			// 	os.Exit(4)
			// }

			// // Destination also needs to be a folder
			// fi, err = os.Stat(dest)
			// if err != nil {
			// 	fmt.Println(err)
			// 	os.Exit(4)
			// }
			// switch mode := fi.Mode(); {
			// case mode.IsDir():
			// 	isDestDir = true
			// default:
			// 	fmt.Println("Error reading destination directory")
			// 	os.Exit(4)
			// }
		} else if arg.UsesWildcards {
			files, err := filepath.Glob(arg.Source)
			if err != nil {
				fmt.Printf("Error opening source files %s: %v\n", arg.Source, err)
				os.Exit(1)
			}
			for _, f := range files {
				CopyFile(f, arg.Dest, arg.IsDestDir)
			}
		} else {
			CopyFile(arg.Source, arg.Dest, arg.IsDestDir)
		}
	}

	os.Exit(0)
}
