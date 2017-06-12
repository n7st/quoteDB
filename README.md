# quoteDB

IRC quotemaker with a web frontend for viewing saved quotes. This project is
currently WIP, so consider it unstable and make sure you back up your quote
database regularly.

## Installation

* Install the software `go get github.com/n7st/quoteDB/...`
* Run `quote-ircbot` for the IRC end
* Run `quote-webui` for the web end

## Usage

* From IRC, `!addquote "first part of first message" "first part of last message"`
will create a quote starting at a line beginning "first part of first message"
and ending at a line beginning "first part of last message".
* From the browser, you can navigate to http://example.site.com:8080/view/{id} or
http://example.site.com:8080/channel/%23channelname.

## License

```
MIT License

Copyright (c) 2017 Mike Jones

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