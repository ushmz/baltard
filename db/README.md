# ratri database


## Requirements

- docker 
- [sql-migrate](https://github.com/rubenv/sql-migrate)

## Structure
```
mysql/
  ├── backup
  │  └── # Backup data. Created by "backup.sh" in this directory.
  ├── data
  │  └──  # Database raw data.
  │       # You have to create this directory before build docker image.
  ├── init.d/
  │  └── # Table structure SQL files that exected in initialization.
  ├── migrations/
  │  └── # Migration files handled with "sql-migrate"
  ├── backup.sh # Backup script for whole table data.
  └── my.cnf    # Mysql global configuration file.
```
