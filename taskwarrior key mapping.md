### [文档链接](https://taskwarrior.org/docs/commands/) :taskwarrior:

## Commands

Here are the commands alphabetically. Version-specific features are labelled with the version in which they were first available.

[add](https://taskwarrior.org/docs/commands/add.html)
Add a new task

annotate
Add an annotation to a task

[append](https://taskwarrior.org/docs/commands/append.html)
Append words to a task description

[calc](https://taskwarrior.org/docs/commands/calc.html) 2.4.0
Expression calculator

config
Modify configuration settings

context
Manage contexts

[count](https://taskwarrior.org/docs/commands/count.html)
Count the tasks matching a filter

delete
Mark a task as deleted

denotate
Remove an annotation from a task

[done](https://taskwarrior.org/docs/commands/done.html)
Complete a task

duplicate
Clone an existing task

edit
Launch your text editor to modify a task

execute
Execute an external command

[export](https://taskwarrior.org/docs/commands/export.html)
Export tasks in JSON format

help
Show high-level help, a cheat-sheet

import
Import tasks in JSON form

[log](https://taskwarrior.org/docs/commands/log.html)
Record an already-completed task

logo
Show the Taskwarrior logo

[modify](https://taskwarrior.org/docs/commands/modify.html)
Modify one or more tasks

[prepend](https://taskwarrior.org/docs/commands/prepend.html)
Prepend words to a task description

purge 2.6.0
Completely removes tasks, rather than change status to `deleted`

start
Start working on a task, make active

stop
Stop working on a task, no longer active

[synchronize](https://taskwarrior.org/docs/commands/synchronize.html)
Syncs tasks with Taskserver

undo
Revert last change

version
Version details and copyright

### Customizable Reports

[Customizable reports](https://taskwarrior.org/docs/report.html) can be modified to suit your needs. They all work in same manner, and mostly differ by the [columns](https://taskwarrior.org/docs/commands/columns.html) shown, the [filter](https://taskwarrior.org/docs/filter.html) applied, and the sorting performed.

Because all the customizable reports work in the same way, [the `list` report](https://taskwarrior.org/docs/commands/list.html) is the only report discussed.

active
Started tasks

all
Pending, completed and deleted tasks

blocked
Tasks that are blocked by other tasks

blocking
Tasks that block other tasks

completed
Tasks that have been completed

[list](https://taskwarrior.org/docs/commands/list.html)
Pending tasks

long
Pending tasks, long form

ls
Pending tasks, short form

minimal
Pending tasks, minimal form

newest
Most recent pending tasks

next
Most urgent tasks

oldest
Oldest pending tasks

overdue
Overdue tasks

ready
Pending, unblocked, scheduled tasks

recurring
Pending recurring tasks

unblocked
Tasks that are not blocked

waiting
Hidden, waiting tasks

### Fixed Reports

The fixed reports have minimal customization, typically only color.

[burndown.daily](https://taskwarrior.org/docs/commands/burndown.html)
Burndown chart, by day

[burndown.monthly](https://taskwarrior.org/docs/commands/burndown.html)
Burndown chart, by month

[burndown.weekly](https://taskwarrior.org/docs/commands/burndown.html)
Burndown chart, by week

calendar
Calendar and holidays

colors
Demonstrates all supported colors

[columns](https://taskwarrior.org/docs/commands/columns.html)
List of report columns and supported formats

commands 2.5.0
List of commands, with their behaviors

diagnostics
Show diagnostics, for troubleshooting

ghistory.annual
History graph, by year

ghistory.monthly
History graph, by month

ghistory.weekly 2.6.0
History graph, by week

ghistory.daily 2.6.0
History graph, by day

history.annual
History report, by year

history.monthly
History report, by month

history.weekly 2.6.0
History report, by week

history.daily 2.6.0
History report, by day

ids
Filtered list of task IDs

[information](https://taskwarrior.org/docs/commands/info.html)
All attributes shown

projects
Filtered list of projects, with task counts

reports
List of available reports

show
Filtered list of configuration settings

stats
Filtered statistics

summary
Filtered project summary

tags
Filtered list of tags

timesheet
Weekly timesheet report

udas
Details of all defined UDAs

uuids
Filtered list of UUIDs

### Helper Commands

Helper commands all begin with an underscore and they exist to provide support for third party add-ons, such as shell completion scripts.

Helper commands generate no extraneous output, and are therefore easy to parse. Helper commands are not intended for regular use, but there is no reason for that to stop you.

_aliases
List of active aliases

_columns
List of supported columns

_commands
List of supported commands

_config
List of confguration setting names

_context
List of defined context names

[_get](https://taskwarrior.org/docs/commands/_get.html)
DOM accessor

_ids
Filtered list of task IDs

_projects
Filtered list of project names

_show
List of `name=value` configuration settings

_tags
Filtered list of tags in use

_udas
List of configured UDA names

[_unique](https://taskwarrior.org/docs/commands/_unique.html) 2.5.0
List of unique values for the specified attribute

_urgency
Filtered list of task urgencies

_uuids
Filtered list of pending UUIDs

_version
Task version (and optional git commit)

_zshattributes
Zsh formatted task attribute list

_zshcommands
Zsh formatted command list

_zshids
Zsh formatted ID list

_zshuuids
Zsh formatted UUID list

