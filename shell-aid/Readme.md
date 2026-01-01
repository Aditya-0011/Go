# Shell-Aid

An AI-powered shell command generator that converts natural language requests into executable shell commands for any operating system and shell environment.

## Purpose

Shell-Aid is a command-line tool that leverages Google's Gemini AI to help users generate shell commands from natural language descriptions. It automatically detects your operating system and shell environment, then generates appropriate commands with detailed explanations and side-effect warnings. Whether you're a beginner learning shell commands or an expert working across multiple platforms, Shell-Aid makes command-line operations more accessible and safer.

## Features

- 🤖 AI-powered command generation using Google Gemini 2.5 Flash
- 🖥️ Cross-platform support (Windows, Linux, macOS)
- 🐚 Multi-shell support (PowerShell, Bash, Zsh, etc.)
- 📝 Step-by-step workflow explanations
- ⚠️ Side-effect warnings for potentially destructive operations
- 🔗 Automatic command chaining for multi-step tasks
- 💬 Interactive console interface

## Configuration

### Prerequisites

- Go 1.25.3 or higher
- Google Gemini API key ([Get one here](https://ai.google.dev/))

### API Key Setup

Before running Shell-Aid, you need to configure your Google Gemini API key:

1. Open `utils/ai/command.go`
2. Locate the `key` constant at the top of the file
3. Replace the empty string with your API key:

```go
const key string = "your-api-key-here"
```

**Note:** For production use, consider using environment variables instead of hardcoding the API key.

## Setup

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Aditya-0011/Go.git
cd Go/shell-aid
```

2. Install dependencies:

```bash
go mod download
```

3. Build the application:

```bash
go build -o shell-aid
```

4. Run the application:

```bash
./shell-aid
```

On Windows:

```powershell
go build -o shell-aid.exe
.\shell-aid.exe
```

## Example

### Basic Usage

```
$ ./shell-aid
Enter your request (press Enter on an empty line to finish), enter 'exit' to quit:
find all .txt files modified in the last 7 days

Processing Request

Workflow:
This command searches for files with the .txt extension that have been modified
within the last 7 days. It uses the find command to recursively search from the
current directory. The -type f flag ensures only files are matched, -name "*.txt"
filters for text files, and -mtime -7 specifies files modified in the last 7 days.

SideEffects:
This is a read-only operation with no side effects. It will not modify, delete,
or create any files. However, searching large directory structures may take time
and consume system resources.

Command:
find . -type f -name "*.txt" -mtime -7
```

### Multi-Step Example

```
Enter your request (press Enter on an empty line to finish), enter 'exit' to quit:
create a directory called backup, copy all .go files to it, and count how many files were copied

Processing Request

Workflow:
This command performs three operations: First, it creates a new directory named
"backup". Second, it copies all files with the .go extension from the current
directory to the backup directory. Third, it counts and displays the number of
.go files that were copied.

SideEffects:
This command will create a new directory and copy files. If the backup directory
already exists, mkdir may fail unless the -p flag is used. The copy operation will
overwrite existing files in the destination without warning.

Command:
mkdir -p backup && cp *.go backup/ && ls backup/*.go | wc -l
```

### Interactive Session

```
$ ./shell-aid
Enter your request (press Enter on an empty line to finish), enter 'exit' to quit:
show disk usage

Processing Request
[... command output ...]

Enter your request (press Enter on an empty line to finish), enter 'exit' to quit:
exit
Exiting.
```

## How It Works

1. **System Detection**: Automatically detects your OS and shell environment
2. **Natural Language Input**: Accept multi-line natural language requests
3. **AI Processing**: Sends your request to Google Gemini with system context
4. **Command Generation**: Receives structured response with workflow, side effects, and command
5. **Display**: Shows formatted output with color-coded sections for easy reading

## Project Structure

```
shell-aid/
├── main.go              # Main application entry point
├── go.mod               # Go module dependencies
├── lib/
│   └── objects.go       # Data structures and types
└── utils/
    ├── ai/
    │   └── command.go   # AI command generation logic
    ├── console/
    │   └── display.go   # Console output formatting
    └── shell/
        ├── posix.go     # POSIX shell detection
        └── windows.go   # Windows shell detection
```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## Security Note

⚠️ **Always review generated commands before executing them**, especially those with potential side effects. Shell-Aid provides explanations and warnings, but you should understand what a command does before running it on your system.
