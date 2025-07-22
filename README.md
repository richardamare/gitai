# GitAI

GitAI is a command-line tool that leverages OpenAI's GPT models to assist with Git operations. It can generate commit messages, summarize changes, and potentially more in the future.

## Requirements

To use GitAI, you need to have Go installed.

## Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/richardamare/gitai.git
    cd gitai
    ```
2.  Build the project:
    ```bash
    go build -o gitai
    ```
3.  (Optional) Move the executable to your PATH:
    ```bash
    mv gitai /usr/local/bin/
    ```

## Usage

Currently, GitAI provides a `commit` command to generate commit messages.

### Generating Commit Messages

To generate a commit message for your staged changes:

```bash
gitai commit
```

The tool will analyze your staged changes and suggest a commit message.

## Environment Variables

GitAI requires your OpenAI API key to function. Set the `OPENAI_API_KEY` environment variable:

```bash
export OPENAI_API_KEY="your_openai_api_key_here"
```

It's recommended to add this line to your shell's configuration file (e.g., `.bashrc`, `.zshrc`, or `.profile`) to set the environment variable automatically.

## Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.
