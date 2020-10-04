# gomon
Application monitoring tool

## What it should:
**X** - Implemented and covered with tests;

**\-** - Not implemented;

**W** - Implemented, but not tested

| Status | Feature |
|:------:|:--------|
| X | Start processes with given arguments and environment variables |
| X | Stop started processes |
| - | Restart them when they crash |
| W | Relay termination signals |
| - | Read their stdout and stderr |
| X | Compile and work on Linux and macOS |


W ability to stop processes when main processes are SIGKILL'ed;
* comments and documentation in code;
* configurable backoff strategy for restarts;
* README file;
X continuous integration configuration;
X integration tests;
* command (package main) that demonstrates the usage;
* unit tests.

