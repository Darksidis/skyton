# skyton

About
-------

This script based on the "Rclone" program works with cloud storages.


How it works
-------

Using the command line, the program monitors updates in cloud storages with a certain frequency, and if they appear, it writes about it. More specifically, it tracks deletion, renaming, moving, changing. Planned as a microservice.

Technologies used
-------

Rclone, command line, postgres


Navigation
-------

-- main.py - the main file.

-- sqlihter.py - file that works with databases
