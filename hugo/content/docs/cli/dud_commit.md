---
title: dud commit
---
## dud commit

Save artifacts to the cache and record their checksums

### Synopsis

Commit saves artifacts to the cache and record their checksums.

For each stage file passed in, commit saves all output artifacts in the cache
and records their checksums in a stage lock file. If no stage files are passed
in, commit will act on all stages in the index. By default, commit will act
recursively on all upstream stages (i.e. dependencies).

```
dud commit [flags] [stage_file]...
```

### Options

```
  -c, --copy   On checkout, copy the file instead of linking.
  -h, --help   help for commit
```

### Options inherited from parent commands

```
      --profile   enable profiling
      --trace     enable tracing
```

### SEE ALSO

* [dud]({{< relref "dud.md" >}})	 - 
