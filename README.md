# Goony

&ndash; Simple Load Testing


## Premise

Load testing a server should be a straightforward process.
Goony uses a text-file that contains all routes/endpoints you want to request on a given host.


## Install

The easiest way at the moment: `go get http://github.com/christophgockel/goony`


## Usage

> `goony [-t|--threads n] [-h|--host http://target-host] file`

Goony has to be called with at least the route file as its argument.
Additionally to that there are two flags: one to configure the number of threads (goroutines), and one to specify the target host.

- `goony file.log`
  - Uses `file.log` to be used as the routes to be requested.
  - Uses the default host `http://localhost` and number of threads (10).

- `goony --host http://hostname.dev:8080 file.log`
  - Uses `file.log` to be used as the routes to be requested.
  - Uses the host `http://hostname.dev:8080` and the default number of threads (10).

- `goony -t 100 -h http://hostname.dev:8080 file.log`
  - Uses `file.log` to be used as the routes to be requested.
  - Uses the host `http://hostname.dev:8080`.
  - Uses 100 threads to execute the requests.
    - If the file has less routes in total, excessive threads will do no work.


### File Format

The expected format of the log file to be used is as follows:

```
/
/an/endpoint
/another/endpoint
/yet/another/endpoint?with=query&strings=possible
```

The routes in this file will be appended to the hostname specified as a command-line argument.


### Output Format

Goony's output is a list of comma separated values.
The idea is that it can be piped to another process for further processing.

```
<Status>,<Start Date>,<Start Time>,<URL>,<HTTP Status>,<Request duration in nanoseconds>,<End Date>,<End Time>,<Status Message>
```

A successful request looks like this:

```
S,2015-11-17,09:10:42.655774898,http://localhost:8080/some/endpoint,200,2353789705,2015-11-17,11:10:45.009564603,
```

In case of a connection error, a line looks like this:

```
F,2015-11-17,09:00:34.782749629,http://localhost:8081/some/endpoint,0,4845,2015-11-17,09:00:34.782754474,Get http://http:localhost:8081/some/endpoint: dial tcp: dial tcp [::1]:8081: getsockopt: connection refused
```


## Contribute

I prefer having my development projects outside of my Go workspace.
The file `link_project_in_workspace.sh` has been added for that.
It will create a symlink in your `$GOPATH` to the directory you cloned this repository into.
This way all import statements will work as expected.

### Running the Tests

Goony uses [gom](https://github.com/mattn/gom) for its dependencies.
Make sure to run `gom -test install` before running the tests.

Execute `gom exec ginkgo -r` to run the test suite.


## Goony?

The name &ldquo;goony&rdquo; was chosen because at the time of writing it, it felt like a relatively _goony_ problem we tried to solve.

All we wanted was to replay a given access-log file from a webserver.

JMeter and httperf seemed to be the tool of choice but weren't capable of just replaying it with a certain amount of threads.
We ran into the issue that while you can specify the number of threads in JMeter, it will actually only replay this amount of requests.
So when we said _run the log file with 100 threads_ JMeter ran only the first 100 lines of the log file.
Adding another iteration in JMeter just replayed the same first 100 lines again.

I'm open to suggestions, and in case anyone knows how to replay a given access-log file with ~450k lines in JMeter, please [open an issue](https://github.com/christophgockel/goony/issues) or get in touch with me. Thank you!
