# safe-renamer
Golang program to rename files to shell-safe characters using ASCII transliteration

## Purpose
To give files safe names, that consist of nothing more than:
* a-z
* A-Z
* 0-9
* underscore
* hyphen
* dot

The hyphen is not allowed at the beginning of the filename.

Any character not allowed is replaced by an underscore.

I use this utility to remove spaces and Unicode characters from filenames because those characters often mess up other
programs and make files more difficult to process.

## Status
Ready to use.

## Installation
This program is written in Google Go language. Make sure that Go is installed and the GOPATH is set up as described in
[How to Write Go Code](https://golang.org/doc/code.html).

The install this program and its dependencies by running:

    go get github.com/BluntSporks/safe-renamer

## Usage
Usage:

    safe-renamer [-commit] FILENAME

Options:

    -commit  Actually do the file renaming, otherwise it just checks that renaming can be done.
