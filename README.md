# quoteDB

[![Build Status](https://drone.netsplit.uk/api/badges/mike/quoteDB/status.svg)](https://drone.netsplit.uk/mike/quoteDB)

IRC quotemaker with a web frontend for viewing saved quotes.

## Installation

### With Docker Compose

* Clone the repository and enter the directory
* Copy `data/example_config.yaml` to `data/config.yaml` and edit it
* `docker-compose build`
* `docker-compose up`

Changing the config after installation is permitted (edit `data/config.yaml` on
the host machine).

### Manually

* Install the software `go get github.com/n7st/quoteDB/...`
* `cd` to the directory where it's installed (probably `~/go/src/git.netsplit.uk/mike/quoteDB`
* Copy `data/config_example.yaml` to `~/.config/netsplit/quoteDB/config.yaml` and edit it
* Run `quoteDB`

You can pass one command-line argument, the path to a YAML configuration file,
to quoteDB:

`quoteDB ~/config.yaml`

## Usage

* From IRC, `!addquote "first part of first message" "first part of last message"`
will create a quote starting at a line beginning "first part of first message"
and ending at a line beginning "first part of last message".
* From the browser, you can navigate to http://example.site.com:8080/view/10
(where quote ID is "10") or http://example.site.com:8080/channel/%23channelname
(where the channel's name is "#channelname", with the hash escaped).

## Limitations

* See the [project's issues](https://git.netsplit.uk/mike/quoteDB/issues).
* Messages starting with quotation marks currently cannot be used as a start or
end point.

## License

```
MIT License

Copyright (c) 2017-2019 Mike Jones

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
