Project 2: Custom Shell Implementation in Go 

This custom shell is built in Go and provides a command-line interface similar to traditional Unix shells like Bash. It allows users to execute both built-in commands (handled internally by the shell) and external commands (executed as subprocesses)

There are five main components of the shell. They include:

Initialization:
The main function initializes the shell by creating a channel for graceful exit.

The Main Loop:
The runLoop function continuously reads input from the user, prints the prompt, and handles user input until an exit signal is received.

Input Handling:
User input is processed by the handleInput function, which identifies the command name and arguments.

Command Execution:
If the command is a built-in command, it is handled internally. Otherwise, it is executed as an external command using the executeCommand function.

Output Display:
The output of the command (if any) is displayed to the user.
	
 
 There are multiple unique commands. Including:

cd: Change directory.

env: Print environment variables.

exit: Exit the shell.

echo: Print arguments to the standard output.

pwd: Print the current working directory.

mkdir: Create directories.

rmdir: Remove directories.

touch: Create files.

To build the program, execute the command “go build main.go” followed by “./main”. Once there you will be prompted to enter commands. To exit the shell, type exit and press Enter.
