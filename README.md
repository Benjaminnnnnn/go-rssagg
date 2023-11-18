An RSS aggregator web server, allowing clients to:

- Add RSS feeds to be collected, the server will scrape all RSS feeds periodically
- Follow and unfollow RSS feeds
- Fetch all of the latest posts from the RSS feeds they follow

Run server:

```
go build && ./rssagg
```

Check out API docs after server started:

```
http://localhost:8080/docs/index.html
```
