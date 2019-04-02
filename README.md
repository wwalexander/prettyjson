prettyjson
==========

Building
--------

    go build

Usage
-----

    prettyjson [-i indent] file

prettyjson pretty-prints JSON input from file. If file is absent, prettyjson
reads from stdin. prettyjson will indent using a tab character by default,
otherwise the string specified by -i.
