# Movies-NAS-DB

A simple script to save the movie data from a NAS (with SMB) in a database.

## Requirement

* Movie file names should follow this naming scheme: `title (year) quality.extension` eg. `Pulp Fiction (1994) 1080p.mkv`
* Create a table like this : 
```sql
CREATE TABLE IF NOT EXISTS `movies` (
  `id` varchar(36) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `year` year(4) DEFAULT NULL,
  `quality` varchar(255) DEFAULT NULL,
  `size` bigint(20) DEFAULT NULL,
  `date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)
```

## Installation

* Create a `.env` from the `.env.example`

## Build

* Default : `go build`
* Linux : `GOOS=linux go build`
* Windows : `GOOS=windows go build`

## Usage

Execute `./movies-nas-db`

You can now use the script via a task scheduler to save your movies.