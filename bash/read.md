# Read command

read line: read from terminal, and store in variable 'line'

```bash
$ read line
hello
$ echo $line
hellp
```

## flag

-t: time limit.

```bash
$ read -t 30
# it will stop after 30 second
```

-p: some tips before input.

```bash
$ read -p 'Input some words:' words
Input some words:
$
```

-d: keep accept input until the key word.

```bash
$ read -d ';' line
a
b
; $ echo $line
a b
```

-a: split input line into a string list.

```bash
$ read -a array
a b c
$ echo ${array[1]}
b
```

-s: silence. words will not print while input.

```bash
$ read -s line
# we type in 'aaaaa', but it will show nothing.

$ echo $line
aaaaa
```

-e: deal with up/down/lift/right and so on.  
it can use up and down to get previous command.

-n: read n chars, rather than a line.

-r: not use '/' as escape character.

-u: use file descripter.
