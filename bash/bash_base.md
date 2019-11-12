# Base

[**Bash tutorial**](https://ryanstutorials.net/bash-scripting-tutorial/bash-input.php)

## Variable

### Arguments

    $0 - The name of the Bash script.
    $1 - $9 - The first 9 arguments to the Bash script. (As mentioned above.)
    $# - How many arguments were passed to the Bash script.
    $@ - All the arguments supplied to the Bash script.
    $? - The exit status of the most recently run process.
    $$ - The process ID of the current script.
    $USER - The username of the user running the script.
    $HOSTNAME - The hostname of the machine the script is running on.
    $SECONDS - The number of seconds since the script was started.
    $RANDOM - Returns a different random number each time is it referred to.
    $LINENO - Returns the current line number in the Bash script.

### Normal variable

There should be no space in a variable, but we can use quote to make it.  
It we want put a variable into another variable, we should use double quotes.

    VARIABLE=Hello
    VARIABLE_WITH_SPACE='Hello World!!!'
    VARIABLE_WITH_VARIABLE_IN_IT="$VARIABLE_WITH_SPACE --wangdi"

### Command substitution

    myCommand=$(echo 'Hello World!!!')

myCommand will store what the command output onto the desk.  
The newlines in the output will be removed, so all the result will be put into one line.

### Export variable

    export VARIABLE

Export exports variable to another script called by the current script.

## Input
