# logger
Golang logger for squarescale

## purposed

This is a (voluntary) very simple logger lib that basically rely upon two ENV variables. `$LOG_LEVEL` and `$SERVICE_NAME`.
The choice of using ENV has been done because of many reasons :

 * We mainly use it for services deployed into container orchestrator. For that case, ENV variables are the best way to configure services.
 * It is, by far, the choice that leads to the most simple implementation.
 * By the simplicity we reduce the bugs
 * By reducing the bugs, serenity to the developper we bring
 * Throught serenity, balance to the force we bring :godmode:

There is no external dependencies. It is basically a simple wrapper around the standard `logger` [package](https://golang.org/pkg/log/) which define and configure five logger. `Debug`, `Info`, `Warn`, `Error` and `Critical`

## how it works

first you must import the package logger:

```
import "github.com/squarescale/logger"

```

Then you may access the five defined [loggers](https://golang.org/pkg/log/#Logger): `Debug`, `Info`, `Warn`, `Error` and `Critical`.

You may change your log level using `$LEVEL`. Five case are considered: empty, `debug`, `info`, `warn` and `error`. When empty, the default level is `info` and `Critical` is alway enabled.

If non empty, `$SERVICE_NAME` will be used in the logs prefix.

May the logs be with you.
