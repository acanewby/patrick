# patrick

## Introduction

`patrick` is a software engineering support utility.

Its purpose is to facilitate the conversion of source code containing hard-coded string literals into a form compatible with most resource management approaches.

Specifically, `patrick`:

* parses source code
* identifies candidate string literals
* generates a substitute resource token
* outputs converted source code
* outputs resource files

### Who is (was) Patrick?

This tool is named in honor of [J. Patrick Desbrow](https://github.com/PatrickDesbrow), one of my closest friends, who introduced me to the world of professional-grade software product engineering.

One of Patrick's many marquee achievements was the systematic globalisation and localisation of [TSER](https://www.cnet.com/tech/tech-industry/oracle-buys-financial-software-maker/) (Treasury Services Evaluation and Reporting)
from hard-coded US English to Japanese (and, indeed, any written language).  He achieved this by writing a tool that parsed approximately 2M lines of C++ source code, identified and tokenized string literals, and generated resource files suitable for inclusion
in the Borland C++ build system used by TSER.  Quite the achievement in 1995.

Sadly, Patrick passed away suddenly in September, 2018.

### Useful features

* define a list of excludable files
* specify resource token prefixes for absolute `"abc = xyz"` and templated `"abc: %s"` literals
* specify the form of a substitute resource lookup function e.g. `util.Resource(<token goes here>)`

## Command examples

### General

* Flags for `inputDir` and `outputDir` should be obvious.

* File format for the `excludesFile` is like this (e.g. one unqualified filename per line):

```shell
resource.go
resources.go
types.go
constants.go
.DS_Store
```

`patrick` will ignore any file matching any of the above filenames in the `excludeFiles` file as it traverses the `inputDir` directory tree.

* A log file (`patrick.log`) will be output to the current working directory, with detail level driven by `logLevel`go.

### List

Allows you to get a listing of the files that would be processed.  Useful to test your input directory / exclusions configuration.

```shell
./patrick list --help                                                                                                                                                                                                                          feature/implement-convert ● ? ↑1

Lists all files that will be targeted for processing, taking into account inclusion and exclusion criteria.

Usage:
  patrick list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --excludeFiles string   file containing base filenames to exclude - one per line
      --inputDir string       input directory
      --logLevel string       log level (debug,info,warn,error,fatal) (default "info")
```

```shell
./patrick list --inputDir=/Users/Anewby/Dropbox/scratch/patrick/input --excludeFiles=/Users/Anewby/Dropbox/scratch/patrick/exclude.list                                                                                                        feature/implement-convert ● ? ↑1
================================================================================
Input directory               : /Users/Anewby/Dropbox/scratch/patrick/input
Output directory              : 
Exclude files                 : /Users/Anewby/Dropbox/scratch/patrick/exclude.list
Log level                     : info
Package identifier            : package
Resource file delimiter       : |
String delimiter              : "
Single line comment delimiter : //
Block comment begin delimiter : /*
Block comment end delimiter   : */
Import keyword                : import
Import block delimiters       : import ( ... )
Const keyword                 : const
Const block delimiters        : const ( ... )
================================================================================
Files to process
--------------------------------------------------------------------------------
Processing file: /Users/Anewby/Dropbox/scratch/patrick/input/config/config.go
Processing file: /Users/Anewby/Dropbox/scratch/patrick/input/config/defaults.go
Processing file: /Users/Anewby/Dropbox/scratch/patrick/input/docker/docker.go
Processing file: /Users/Anewby/Dropbox/scratch/patrick/input/util/smoke.go
Processing file: /Users/Anewby/Dropbox/scratch/patrick/input/util/util.go
--------------------------------------------------------------------------------
```

### Convert

Processes files according to `inputDir` and `excludesFile`, as described above.

It identifies string literals and substitutes them in the output file with tokens recorded in the corresponding package-specific resource file.




## Build

```shell
go build -o ./dist/patrick -ldflags="-X 'github.com/acanewby/patrick/cmd/patrick.version=0.0.2'" main.go  
```
