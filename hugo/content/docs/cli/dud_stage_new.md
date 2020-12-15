---
title: dud stage new
---
## dud stage new

Create a stage from the command-line and print it

### Synopsis

Stage creates a stage from the command-line and prints it to STDOUT.

The output of this command can be redirected to a file and modified further as
needed. For example:

dud stage -o data/ python download_data.py > download.yaml

```
dud stage new [flags]
```

### Options

```
  -d, --dep strings       one or more dependent files or directories
  -h, --help              help for new
  -o, --out strings       one or more output files or directories
  -w, --work-dir string   working directory for the stage's command
```

### Options inherited from parent commands

```
      --profile   enable profiling
      --trace     enable tracing
```

### SEE ALSO

* [dud stage]({{< relref "dud_stage.md" >}})	 - Commands for interacting with stages and the index
