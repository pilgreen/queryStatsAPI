# queryStatsAPI
A simple program to build the correct signature to make queries to the STATS API.

## Flags

**key** -- your Stats API key

**secret** -- your Stats API secret

**endpoint** -- the API endpoint you would like to retrieve

The endpoint flag will take the first positional argument as well. You can also pass in the key and secret using environment variables. This is handy when multiple calls need to be made in a Makefile or another similar situation.

```
stats: export STATS_KEY=xxx
stats: export STATS_SECRET=xxx
stats: 
  queryStatsApi -endpoint xxx > data/xxx.json
  ...
```
