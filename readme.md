# Random RSS

To run : `go run .`

You can then query an rss-like file from [http://localhost:8080/rss.xml](http://localhost:8080/rss.xml)

You can specify the seed & the size as a query parameter

## Exemple

| URL                                              | Seed | Size |
| ------------------------------------------------ | ---- | ---- |
| `http://localhost:8080/rss.xml`                  | 1    | 20   |
| `http://localhost:8080/rss.xml?size=69`          | 1    | 69   |
| `http://localhost:8080/rss.xml?seed=420`         | 420  | 20   |
| `http://localhost:8080/rss.xml?seed=420&size=69` | 420  | 69   |
