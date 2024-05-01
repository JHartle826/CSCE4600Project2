Custom Shell

This is a simple custom shell implemented in Go. It provides a command-line interface where users can execute built-in commands and external commands.

Functionality

The shell supports the following built-in commands:

cd: Change directory.

env: Print environment variables.

exit: Exit the shell.

echo: Print arguments to the standard output.

pwd: Print the current working directory.

mkdir: Create directories.

rmdir: Remove directories.

touch: Create files.

External commands are executed as subprocesses.

How to Execute/Test

To execute the shell:

Clone this repository to your local machine.

Ensure you have Go installed.

Build the project:

go build main.go

Run the executable:

./main

Once the shell is running, you can type commands and press Enter to execute them. You can test built-in commands like cd, env, echo, etc., as well as external commands available in your system.

To exit the shell, simply type exit and press Enter.
