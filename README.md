# Speedy Test

Speedy test is a full-service implementation of https://github.com/adolfintel/speedtest.

Simply build and run and you'll have a server running at https://localhost:8080
which will serve a speedtest for users.

## Documentation

Please refer to the upstream project documentation at https://github.com/adolfintel/speedtest/blob/master/doc.md
for advanced options you can use in speedtest.js.

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
