# Pass "Shell is fun" (3 arguments) to the script(prog.sh) as an arguments and print the length of the arguments.

#!/bin/bash
function File {
    # think you are inside the file
    # Change Here
    echo $#
}

# Do not change anything
if [ ! $# -lt 1 ]; then
    File $*
    exit 0
fi

# change here
# here you can pass the arguments
bash prog.sh Shell is fun
