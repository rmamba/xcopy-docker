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

// func IsDirectory(path string) (bool, int) {
// 	fi, err := os.Stat(path)
// 	if err != nil {
// 		if strings.HasSuffix(path, "/") {
// 			return true, 0
// 		}
// 		if os.Getenv("XCOPY_DEBUG") == "true" {
// 			fmt.Println(err.Error())
// 		}
// 		return false, -3
// 	}
// 	switch mode := fi.Mode(); {
// 	case mode.IsDir():
// 		return true, 0
// 	}
// 	return false, 0
// }

// Hello returns a greeting for the named person.
func ParseArguments(args []string) (InputArguments, int) {
	arguments := InputArguments{}

	if len(args) < 2 {
		return arguments, -1
	}

	// Parse arguments
	for _, a := range args {
		if strings.HasPrefix(a, "/") && (len(a) == 2 || len(a) == 3) {
			switch a {
			case "/A":
				arguments.Flags.A = true
			case "/C":
				arguments.Flags.C = true
			case "/D":
				arguments.Flags.D = true
			case "/E":
				arguments.Flags.E = true
			case "/F":
				arguments.Flags.F = true
			case "/H":
				arguments.Flags.H = true
			case "/I":
				arguments.Flags.I = true
			case "/L":
				arguments.Flags.L = true
			case "/M":
				arguments.Flags.M = true
			case "/N":
				arguments.Flags.N = true
			case "/P":
				arguments.Flags.P = true
			case "/Q":
				arguments.Flags.Q = true
			case "/R":
				arguments.Flags.R = true
			case "/S":
				arguments.Flags.S = true
			case "/T":
				arguments.Flags.T = true
			case "/V":
				arguments.Flags.V = true
			case "/W":
				arguments.Flags.W = true
			case "/Y":
				arguments.Flags.Y = true
			case "/-Y":
				arguments.Flags.MinusY = true
			}
		} else {
			if strings.HasPrefix(a, "/D") && (len(a) == 8) {
				arguments.Flags.D = true
				arguments.Flags.ParamsD = a[2:]
			} else {
				if arguments.Source == "" {
					arguments.Source = strings.ReplaceAll(strings.Trim(a, "\""), "\\", "/")
					arguments.UsesWildcards = strings.Contains(arguments.Source, "*")
					arguments.IsSourceDir = strings.HasSuffix(arguments.Source, "/")
				} else {
					arguments.Dest = strings.ReplaceAll(strings.Trim(a, "\""), "\\", "/")
					arguments.IsDestDir = strings.HasSuffix(arguments.Dest, "/")
				}
			}
		}
	}

	if arguments.Dest == "" || arguments.Source == "" {
		return arguments, -2
	}

	// exitCode := 0
	// arguments.IsSourceDir, exitCode = IsDirectory(arguments.Source)
	// if exitCode != 0 {
	// 	return arguments, exitCode
	// }

	// arguments.IsDestDir, _ = IsDirectory(arguments.Dest)
	return arguments, 0
}

func CopyFile(sourceFileName string, dest string, isDestDir bool) {
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
		fmt.Println("Go implementation of MS-DOS xcopy command v0.0.5")
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
		os.Exit(errCode)
	}

	if args.IsSourceDir {
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
	} else if args.UsesWildcards {
		files, err := filepath.Glob(args.Source)
		if err != nil {
			fmt.Printf("Error opening source files %s: %v\n", args.Source, err)
			os.Exit(1)
		}
		for _, f := range files {
			CopyFile(f, args.Dest, args.IsDestDir)
		}
	} else {
		CopyFile(args.Source, args.Dest, args.IsDestDir)
	}

	os.Exit(0)
}
