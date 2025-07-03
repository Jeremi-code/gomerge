# GoMerge

GoMerge is a private CLI tool to simplify my daily development workflow.  
It helps automate common git operations, starts Docker containers, and manages migration processes with a single command.

## Features

- Simplifies git branch switching, pulling, and merging
- Remembers and quickly opens your most recent folder
- Starts Docker containers for your project
- Runs database migrations automatically

### Common Flags

- `-recent`, `-r`   Open recent folder
- `-folder`, `-f`   Folder to open
- `-branch`, `-b`   Branch to use
- `-help`, `-h`    Show help message

### Examples

```sh
gomerge -recent
gomerge -folder myfolder
gomerge -branch feature-xyz
gomerge myfolder
```

---

**This tool is for private/internal use only.**