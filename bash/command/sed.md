# Sed command

sed: a command to deal will txt file.

    sed [Flag] file

## Flag

    -e script: use a script to deal with file.

> ```bash
> $ sed -e '2,$d' -e 's/bash/blueshell/'
> # delete from 2 to the last line, and change 'bash' to 'blueshell'
> # if there is only one script, -e can be omit.
> ```

    -f file: use a script file to deal with file.

    -n, --quiet, --silence: only show the line changed.

    -i: modify the file(It is dangerous for the file will be changed.)

## Script

### object to be operated

- line: use number/^/\$ to specify the line. and '2,5' is a score.
- content: use /content(regular expression)/ to search.

```bash
# line
$ sed '2,5 script'
# content
$ sed '/content/script'
```

### operation

a,i,d,p,c will operate a hold line

s will operate the matched part

    a: add something to the next line.

> ```bash
> $ sed '4 a newline' file
> # add a new line at the 5th line
> ```

    i: insert something to the previous line.

> ```bash
> $ sed '4 i newline' file
> # insert new line at the 4th line
> ```

    d: delete something

> ```bash
> $ sed '2,5 d'
> ```

    p: print

    c: change

> same to i and a

    s: substitute

> ```bash
> $ sed 's/content/newline/num'
> # the num used to specify which part will be operated. If set g, all the parts that match
> # the content will be operated.
> ```
