gow
===

A Go Web Server (for static, local files).

I wrote this because I needed to be able to serve web sites over
`http://`, instead of `file://`, because things like CSS font
loading won't behave across local fs URIs.  Go figure.

I don't expect anyone else to find this useful, but repos should
not go without README's, so here goes:

`gow` has two options: `--root` to specify where the files are at
(defaults to the current working directory) and `--port` to set
the TCP port to bind / listen on.

If you forget these two options, there is a third option
(surprise) called `--help` that ought to jog your memory.
