# gomon
Application monitoring tool

## Features
**X** - Implemented and covered with tests;

**\-** - Not implemented;

**W** - Implemented, but not tested;

**P** - Partially;

### What it should:

| Status | Feature |
|:------:|:--------|
| X | Start processes with given arguments and environment variables |
| X | Stop started processes |
| - | Restart them when they crash |
| X | Relay termination signals |
| X | Read their stdout and stderr |
| X | Compile and work on Linux and macOS |

### Optional features
| Status | Feature |
|:------:|:--------|
| W | Ability to stop processes when main processes are SIGKILL'ed |
| P | Comments and documentation in code |
| - | Configurable backoff strategy for restarts |
| X | README file |
| X | Continuous integration configuration |
| X | Integration tests | 
| - | Command (package main) that demonstrates the usage |
| - | Unit tests |
