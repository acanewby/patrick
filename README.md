# patrick

## Introduction

`patrick` is a software engineering support utility.

Its purpose is to facilitate the conversion of source code containing hard-coded string literals into a form compatible with most resource management approaches.

It can be used to prepare projects for multilingual support and/or centralization of string resources.

Specifically, `patrick`:

* parses source code
* identifies candidate string literals
* generates a substitute resource token
* outputs converted source code
* outputs resource files

_Note: Currently, only [Go](https://go.dev) source code parsing is explicitly supported.  However, other closely-related languages (e.g. Java/C#) may also work. Extension to support other languages is feasible, and may be supported in the future._

## Who was Patrick?

This tool is named in honor of [J. Patrick Desbrow](https://github.com/PatrickDesbrow), one of my closest friends, who introduced me to the world of professional-grade software product engineering.

One of Patrick's many marquee achievements was the systematic globalisation and localisation of [TSER](https://www.cnet.com/tech/tech-industry/oracle-buys-financial-software-maker/) (Treasury Services Evaluation and Reporting)
from hard-coded US English to Japanese (and, indeed, any written language).  He achieved this by writing a tool that parsed approximately 2M lines of C++ source code, identified and tokenized string literals, and generated resource files suitable for inclusion
in the Borland C++ build system used by TSER (and for which he also ported the [OWL](https://en.wikipedia.org/wiki/Object_Windows_Library) library from 16- to 32-bit as a party trick).  Quite the achievement in 1995.

Sadly, Patrick passed away suddenly in September, 2018. I still miss him.

## Sample output

__Input:__ in/mypackage/myfile.go

This original source file contains two hardcoded string literals.

```go
func (cfg *Service) CreateOrUpdateAnalysisConfig(conf AnalysisConfig) {

	runtime.LogDebugf(cfg.ctx, "Saving analysis config: %+v", conf)
	fullKey := cfg.analysisConfigFullKey(conf.Key)
	runtime.LogDebugf(cfg.ctx, "Config key: %s", fullKey)

	...
```

__Output:__ out/mypackage/mypackage.resource

The literals have been identified, extracted, and associated with resource tokens.

```go
Resource_00010001 = "Saving analysis config: %+v"
Resource_00010002 = "Config key: %s"
```

__Output:__ out/mypackage/myfile.go

The new output source file contains resource lookup function calls, which reference the appropriate resource token.

```go
func (cfg *Service) CreateOrUpdateAnalysisConfig(conf AnalysisConfig) {

    runtime.LogDebugf(cfg.ctx, util.ResourceLookup(mypackage.Resource_00010001), conf)
    fullKey := cfg.analysisConfigFullKey(conf.Key)
    runtime.LogDebugf(cfg.ctx, util.ResourceLookup(mypackage.Resource_00010002), fullKey)

    ...
```

## Useful features

* define a list of excludable files
* template strings to customize output resource file format
* start point and format parameters for sequential resource token numbering
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
./patrick list --inputDir=/Users/Anewby/Dropbox/scratch/patrick/input --excludeFiles=/Users/Anewby/Dropbox/scratch/patrick/exclude.list
================================================================================
Input directory               : /Users/Anewby/Dropbox/scratch/patrick/input
Output directory              : 
Overwrite output              : false
Exclude files                 : /Users/Anewby/Dropbox/scratch/patrick/exclude.list
Log level                     : info
Package identifier            : package
Resource file prefix          : 
Resource file delimiter       :  = 
Resource file suffix          : 
Resource index start          : 10000
Resource index zero pad       : 8
Resource token prefix         : Resource_
Resource function template    : 
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

```shell
./patrick convert --help 

Processes the identified list of files and:

- identifies and tokenizes string literals
- outputs converted source files
- outputs associated resource files

Usage:
  patrick convert [flags]

Flags:
  -h, --help                              help for convert
      --outputDir string                  output directory
      --overwriteOutput                   replace contents of outputDir
      --resourceFileDelimiter string      text to be written between resource identifier and value in resource file (default " = ")
      --resourceFilePrefix string         text to be written before resource identifier in resource file
      --resourceFileSuffix string         text to be written after value in resource file
      --resourceFunctionTemplate string   resource function to be substituted into source code in place of each identified string literal. Must include %%RESOURCE_TOKEN%%, which will be the placeholder for the generated resource token
      --resourceIndexStart uint           starting value for sequentially-numbered resource tokens (default 10000)
      --resourceIndexZeroPad uint8        width of zero-padded resource token index number (default 8)
      --resourceTokenPrefix string        prefix string to be prepended to resource tokens (default "Resource_")

Global Flags:
      --excludeFiles string   file containing base filenames to exclude - one per line
      --inputDir string       input directory
      --logLevel string       log level (debug,info,warn,error,fatal) (default "info")
```

#### Run with defaults

The built-in defaults produce a resource file format, usable in a variety of contexts:

```shell
./patrick convert --inputDir=/Users/acanewby/Dropbox/scratch/patrick/input \
--outputDir=/Users/acanewby/Dropbox/scratch/patrick/output --overwriteOutput=true \
--excludeFiles=/Users/acanewby/Dropbox/scratch/patrick/exclude.list \
--resourceFunctionTemplate="PKG.utilSvc.Resource(PKG.%%RESOURCE_TOKEN%%)"
====================================================================================================
Configuration
----------------------------------------------------------------------------------------------------
General
----------------------------------------------------------------------------------------------------
Input directory               : /Users/acanewby/Dropbox/scratch/patrick/input
Output directory              : /Users/acanewby/Dropbox/scratch/patrick/output
Overwrite output              : true
Exclude files                 : /Users/acanewby/Dropbox/scratch/patrick/exclude.list
Log level                     : info
Resource file prefix          : 
Resource file delimiter       :  = 
Resource file suffix          : 
Resource index start          : 10000
Resource index zero pad       : 8
Resource token prefix         : Resource_
Resource function template    : PKG.utilSvc.Resource(PKG.%%RESOURCE_TOKEN%%)
----------------------------------------------------------------------------------------------------
Source code-specific
----------------------------------------------------------------------------------------------------
Package identifier            : package
String delimiter              : "
Single line comment delimiter : //
Block comment begin delimiter : /*
Block comment end delimiter   : */
Import keyword                : import
Import block delimiters       : import ( ... )
Const keyword                 : const
Const block delimiters        : const ( ... )
====================================================================================================
====================================================================================================
Files to process
----------------------------------------------------------------------------------------------------
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/config/config.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/config/defaults.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/docker/docker.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/util/smoke.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/util/util.go
====================================================================================================
```

```shell
cat /Users/acanewby/Dropbox/scratch/patrick/output/config/config.resource 
Resource_00010001 = "Saving analysis config: %+v"
Resource_00010002 = "Config key: %s"
<snip/>
Resource_00010061 = "Removing analysis subkey from config : %s"
Resource_00010062 = "Analysis map (after prune) : %v"
Resource_00010063 = "Resetting config"
```

#### Advanced use cases

However, it is also possible to generate a wide variety of resource file formats:

```shell
./patrick convert --inputDir=/Users/acanewby/Dropbox/scratch/patrick/input \
--outputDir=/Users/acanewby/Dropbox/scratch/patrick/output --overwriteOutput=true \
--excludeFiles=/Users/acanewby/Dropbox/scratch/patrick/exclude.list \
--resourceFunctionTemplate="PKG.utilSvc.Resource(PKG.%%RESOURCE_TOKEN%%)" \
--resourceFilePrefix="{" --resourceFileDelimiter=", " --resourceFileSuffix="}," \
--resourceTokenPrefix=Res#  --resourceIndexStart 1000 --resourceIndexZeroPad 6
====================================================================================================
Configuration
----------------------------------------------------------------------------------------------------
General
----------------------------------------------------------------------------------------------------
Input directory               : /Users/acanewby/Dropbox/scratch/patrick/input
Output directory              : /Users/acanewby/Dropbox/scratch/patrick/output
Overwrite output              : true
Exclude files                 : /Users/acanewby/Dropbox/scratch/patrick/exclude.list
Log level                     : info
Resource file prefix          : {
Resource file delimiter       : , 
Resource file suffix          : },
Resource index start          : 1000
Resource index zero pad       : 6
Resource token prefix         : Res#
Resource function template    : PKG.utilSvc.Resource(PKG.%%RESOURCE_TOKEN%%)
----------------------------------------------------------------------------------------------------
Source code-specific
----------------------------------------------------------------------------------------------------
Package identifier            : package
String delimiter              : "
Single line comment delimiter : //
Block comment begin delimiter : /*
Block comment end delimiter   : */
Import keyword                : import
Import block delimiters       : import ( ... )
Const keyword                 : const
Const block delimiters        : const ( ... )
====================================================================================================
====================================================================================================
Files to process
----------------------------------------------------------------------------------------------------
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/config/config.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/config/defaults.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/docker/docker.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/util/smoke.go
Processing file: /Users/acanewby/Dropbox/scratch/patrick/input/util/util.go
====================================================================================================

```

```shell
cat /Users/acanewby/Dropbox/scratch/patrick/output/config/config.resource        \
{Res#001001, "Saving analysis config: %+v"},
{Res#001002, "Config key: %s"},
{Res#001003, "%s.%s"},
<snip/>
{Res#001061, "Removing analysis subkey from config : %s"},
{Res#001062, "Analysis map (after prune) : %v"},
{Res#001063, "Resetting config"},
```

