# Blog-aggregator

Blog aggregator 

Requirements to run:
Golang and Postgres installed

To install you need to run this command on the root: go install .

To use you need:
1. A .json file at home directory called: .gatorconfig.json
2. A database using Postgres on terminal called: gator
3. A local connection URL on "db_url" from .gatorconfig.json
   
template (URL): protocol://username:password@host:port/database

template (Postgres connection URL): postgres://postgres:postgres@localhost:5432/gator

Commands:
register: register and set the current user on .gatoconfig.json (needs username)
login: Set other user to current user (needs username)
users: list all registered users (doesn't need arguments)
reset: delete all registered users records (doesn't need arguments)

addfeed: Add a feed to be aggregate (needs feed name and url)
feeds: show all added feeds (doesn't need arguments)
follow: follows a feed (needs url) 
unfollow: unfollow a feed (needs url)
following: shows all following feeds (doesn't need arguments)

agg: fetch feeds added with addfeed in a time interval and transform them into a post (needs time interval)
Remember not to set the time interval too low, 1m is the recommended time

browse: show a specific number of posts, 2 is the default quantity (quantity is optional)
