# Speedy Test

Speedy test is a full-server implementation of https://github.com/adolfintel/speedtest.

Simply build and run and you'll have a server running at https://localhost:8080
which will serve a speedtest for users.

## Speed Test Documentation

Please refer to the upstream project documentation at https://github.com/adolfintel/speedtest/blob/master/doc.md
for advanced options you can use in speedtest.js.

## Server Documentation

### Building

`go build -o bin/speedtest speedtest.go`

### Installing

`go install github.com/lfkeitel/speedtest`

### Usage

`speedtest [options]`

Options:

- `-addr`: Address and port used for HTTP server
- `-db`: Type of database to use for telemetry. One of "none" (default), "log", "csv", "psql".
- `-dbaddr`: Database address (Default: localhost)
- `-dbfile`: Filename for CSV database file (Default: telemetry.txt)
- `-dbpass`: Database user password
- `-dbuser`: Database username (Default: speedtest)
- `-dbname`: Database name (Default: speedtest)
- `-dbport`: Database port


## Telemetry

The client will, by default, send telemetry data back to the server. This data contains the results of
the speed, ping, and jitter tests. Extra data such as a timestamp, the remote address, and a session
ID are added to the event before saving to a data store.

By default, telemetry storage is disabled even though the client will send it.

### Using PostgreSQL

Make a new database using the schema in the misc folder as a template. The sql file assumes the database
is named "speedtest". Please edit accordingly.

`speedtest -db psql -dbport 5432 -dbuser postgres -dbpass example` (Connect to localhost:5432 and use the database speedtest)

### Other database types

* `none`: Disables telemetry storage.
* `console`: Prints telemetry data to stdout.
* `csv`: Saves telemetry data in CSV format to a file.
* `json`: Saves telemetry data in JSON format to a file. One JSON object per line.

### Session ID

When a client loads the speed test page, they are given a unique session ID which allows for data to be grouped
by a single test session. For example, if the client runs three speed tests in the same browsing session, those
results will have the same session ID so queries can be done against a single ID to get an overall average on
a data point for each session. The session cookie will last until the client closes their browser.

## Contributing

Contributions are welcome! Before submitting PRs, please make sure your changes pass the lint tests
and make sure to regenerate the minified file.

```shell
npm install
npm test
npm run uglify
```

## License

Server code:

Copyright (C) 2018 Lee Keitel

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License along with this program. If not, see https://www.gnu.org/licenses/lgpl.

Speedtest code:

Copyright (C) 2016-2018 Federico Dossena

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Lesser General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License along with this program. If not, see https://www.gnu.org/licenses/lgpl.
