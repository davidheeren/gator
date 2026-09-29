# Gator 

Gator is a command-line RSS feed aggregator built in Go as part of the Boot.dev course.

## Dependencies

* Go
* PostgreSQL

## Installation

Clone the repository and build it:

```bash
git clone https://github.com/davidheeren/gator.git
cd gator
go build
```

Run Gator with:

```bash
./gator
```

To install it so you can run `gator` from anywhere, make sure `~/go/bin` is in your `PATH`, then run:

```bash
go install
```

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

## Commands

```
`register`  - make a new user and login  
`login`     - login for an existing user  
`users`     - list all users  
`addfeed`   - add a new rss feed url and follow it  
`follow`    - follow an rss and existing rss feed  
`unfollow`  - unfollow existing rss feed  
`following` - list feeds user is following  
`feeds`     - list all feeds  
`agg`       - update database of posts from rss feeds in the background  
`browse`    - list recent posts of feeds the user is following  
`reset`     - reset everything in the database  

```
