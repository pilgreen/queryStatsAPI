# queryStatsAPI
A simple program to build the correct signature to make queries to the STATS API.

## Environment variables

In order to use this program, you need to set environment variables for your key and secret. 

To do so in a single command you can use:

```
env STATS_KEY=xxx STATS_SECRET=xxx queryStatsApi -endpoint {YOUR ENDPOINT}
```

I use Makefiles to build more than one part of a site at a time with the following pattern:

```
stats: export STATS_KEY=xxx
stats: export STATS_SECRET=xxx
stats: 
  queryStatsApi -endpoint xxx > data/xxx.json
  ...
```
