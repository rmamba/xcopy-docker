package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

func main() {
	var fi fs.FileInfo
	// var files []fs.DirEntry
	var err error

	// useD := false
	// useY := false

	source := ""
	dest := ""
	isSourceDir := false
	// isDestDir := false

	args := os.Args[1:] // get arguments from command line
	if len(args) < 2 {
		fmt.Println("Go implementation of MS-DOS xcopy coppand v0.0.1")
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
		os.Exit(1)
	}

	// Parse arguments
	for _, a := range args {
		if strings.HasPrefix(a, "/") {
			// switch a {
			// case "/Y":
			// 	useY = true
			// case "/D":
			// 	useD = true
			// 	// default:
			// 	// 	// freebsd, openbsd,
			// 	// 	// plan9, windows...
			// 	// 	fmt.Printf("%s.\n", os)
			// }
		} else {
			if source == "" {
				source = a
			} else {
				dest = a
			}
		}
	}

	fi, err = os.Stat(source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		isSourceDir = true
	}

	if isSourceDir {
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
	} else {
		src, err := os.Open(source)
		if err != nil {
			fmt.Printf("Error opening source file %s: %v\n", args[0], err)
			os.Exit(1)
		}
		defer src.Close()

		dst, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Error creating destination file %s: %v\n", args[1], err)
			os.Exit(4)
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		if err != nil {
			fmt.Println("Error copying data from source to destination")
			os.Exit(4)
		}
	}

	// src, err := os.Open(args[0])
	// // if err != nil {
	// // 	fmt.Printf("Error opening source directory %s: %v\n", args[0], err)
	// // 	os.Exit(1)
	// // }
	// // defer src.Close()

	// // dst, err := os.Create(args[1])
	// // if err != nil {
	// // 	fmt.Printf("Error creating destination directory %s: %v\n", args[1], err)
	// // 	os.Exit(1)
	// // }
	// // defer dst.Close()

	// for _, file := range files {
	// 	// skip directories and links
	// 	if !file.IsDir() && !file.Mode().IsRegular() {
	// 		continue
	// 	}

	// 	srcFile, err := os.Open(filepath.Join(src, file.Name()))
	// 	if err != nil {
	// 		fmt.Printf("Error opening source file %s: %v\n", file.Name(), err)
	// 		continue
	// 	}
	// 	defer srcFile.Close()

	// 	dstFile, err := os.Create(filepath.Join(dst, file.Name()))
	// 	if err != nil {
	// 		fmt.Printf("Error creating destination file %s: %v\n", file.Name(), err)
	// 		continue
	// 	}
	// 	defer dstFile.Close()

	// 	_, err = io.Copy(dstFile, srcFile)
	// 	if err != nil {
	// 		fmt.Println("Error copying data from source to destination")
	// 		return
	// 	}
	// }

	// fmt.Printf("Successfully copied %d files\n", len(files))
	os.Exit(0)
}
