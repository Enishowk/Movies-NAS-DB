# Movies-NAS-DB

A simple script to save the movie data from a NAS (with SMB) in a database.

## Requirement

* Movie file names should follow this naming scheme: `title (year) quality.extension` eg. `Pulp Fiction (1994) 1080p.mkv`
* Create a table like this : 
```sql
CREATE TABLE IF NOT EXISTS `movies` (
  `id` varchar(36) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `year` year DEFAULT NULL,
  `quality` varchar(255) DEFAULT NULL,
  `size` bigint DEFAULT NULL,
  `downloadedAt` datetime DEFAULT NULL,
  `createdAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)
```

## Installation

* Create a `.env` from the `.env.example`

## Build

* Default : `go build`
* Linux : `GOOS=linux go build -ldflags '-X main.env_path=/YOUR_PATH/.env'
* Windows : `GOOS=windows go build -ldflags '-X main.env_path=C:\YOUR_PATH\\.env'

## Usage

Execute `./movies-nas-db`

You can now use the script via a task scheduler to save your movies.