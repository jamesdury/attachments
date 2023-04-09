# Email Attachments

## Intro

Inspired by
[hey.com/features/all-files](https://www.hey.com/features/all-files/),
"attachments" reads emails from your local [notmuch](https://notmuchmail.org)
database, with a query to only display emails with attachments. 

The plan is to provide quick access to the files, with only a small amount of
filters e.g. "by Date", "by Sender".

## Status

Still in development
![WIP](https://user-images.githubusercontent.com/97810962/229850607-7e36ea67-fe25-4844-924f-c421159b6a0d.gif)

## Requirements

- [notmuchmail.org](https://notmuchmail.org) on local file system (path currently hard coded to $HOME/Mail)
- The emails stored locally on the machine
  - _I use [isynrc](https://isync.sourceforge.io)_
