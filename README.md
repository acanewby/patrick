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

## Command summary



## Build

```shell
go build -o ./dist/patrick -ldflags="-X 'github.com/acanewby/patrick/cmd/patrick.version=0.0.2'" main.go  
```
