# Awk

    awk '{script}' file

do something on each line

## BEGIN and END

    awk 'BEGIN {script1} {script2} END {script3}'

Invoke script1, then use script2 on each line. Finally, use script3.

## Regular expression

    awk '/regular expression/ {script}'

Use script on the lines match regular expression.

## Bool

    awk 'bool {script}'

```bash
awk '$1=="root" {print $0}'
```

## Constants

    $0: the whole line
    $1~$n: the n'th segment
    FS: segment separator, default ' '
    RS: line separator, default '\n'
    NF: segments num this line
    NR: lines num
    OFS: output segment separator
    ORS: output line separator
    ...

## Flag

-F: set a charactor as FS

-v var=value: set a value

...
