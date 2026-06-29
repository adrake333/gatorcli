Thank you for using GatorCLI! A blog aggregator built using Go and Postgres in your CLI.



To run GatorCLI you will need the following installed:
~Postgres
~Go



To install GatorCLI:
~run: "go install github.com/adrake333/gatorcli@latest"



Setting up Config:
1) The Config file should live in your home directory (~/.gatorconfig.json on Linux)
2) What it should contain:
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
3) Replace the information above to include your own Postgres username, password, host, port, and database name.



Using GatorCLI:
~Commands:
- "gatorcli login {name}" - Logs in as a specific user
- "gatorcli register {name}" - Registers a new user
- "gatorcli reset" - Resets the instance
- "gatorcli users" - Displays registered users
- "gatorcli agg {duration}" - Begins aggregating feeds for the current user every interval of duration
- "gatorcli addfeed {feed_name} {url} - Adds a feed to GatorCLI
- "gatorcli feeds" - Prints feeds
- "gatorcli follow {feed_id}" - Follows a feed for the current user
- "gatorcli following" - Displays feeds followed by current user
- "gatorcli unfollow {feed_id}" - Unfollows a feed for the current user
- "gatorcli browse {optional amount} - Displays specified amount of feed Titles and URLs for current user (default 2). Recently published first.
