# XCOPY for building docker containers

This is xcopy golang program that acts in a same way as original MS-DOS xcopy command
so that it can be used in docker containers with outh any source modifications.

# Use

XCOPY source [destination] [/switches]
  source       Specifies the directory and/or name of file(s) to copy.
  destination  Specifies the location and/or name of new file(s).
  /A           Copies only files with the archive attribute set and doesn't change the attribute.
  /C           Continues copying even if errors occur.
  /D[:M/D/Y]   Copies only files which have been changed on or after the specified date. When no date is specified, only files which are newer than existing destination files will be copied.
  /E           Copies any subdirectories, even if empty.
  /F           Display full source and destination name.
  /H           Copies hidden and system files as well as unprotected files.
  /I           If destination does not exist and copying more than one file, assume destination is a directory.
  /L           List files without copying them. (simulates copying)
  /M           Copies only files with the archive attribute set and turns off the archive attribute of the source files after copying them.
  /N           Suppresses prompting to confirm you want to overwrite an existing destination file and skips these files.
  /P           Prompts for confirmation before creating each destination file.
  /Q           Quiet mode, don't show copied filenames.
  /R           Overwrite read-only files as well as unprotected files.
  /S           Copies directories and subdirectories except empty ones.
  /T           Creates directory tree without copying files. Empty directories will not be copied. To copy them add switch /E.
  /V           Verifies each new file.
  /W           Waits for a keypress before beginning.
  /Y           Suppresses prompting to confirm you want to overwrite an existing destination file and overwrites these files.
  /-Y          Causes prompting to confirm you want to overwrite an existing destination file.

# ToDO
  - [] Implement /A switch
  - [] Implement /C switch
  - [] Implement /D switch
  - [] Implement /D[:M/D/Y] switch
  - [] Implement /E switch
  - [] Implement /F switch
  - [] Implement /H switch
  - [] Implement /I switch
  - [] Implement /L switch
  - [] Implement /M switch
  - [] Implement /N switch
  - [] Implement /P switch
  - [] Implement /Q switch
  - [] Implement /R switch
  - [] Implement /S switch
  - [] Implement /T switch
  - [] Implement /V switch
  - [] Implement /W switch
  - [] Implement /Y switch
  - [] Implement /-Y switch
  - [] Add testing
