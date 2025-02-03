# `pubbo`

## when you absolutely positively want a file to be available

There are lots of ways of making content available online. Most of those ways are deceptively heavy. Run a web server behind a reverse-proxy behind a load-balancer, etc etc ... all to make a single file available.

What if you went the other direction? What if you removed all the dependencies, attack surfaces, and general cruft traditionally included in hosting static content, and went as minimal as possible?

You'd get `pubbo`. Or something similar.

`pubbo` hosts a single file on a Unix socket. You read the socket, you get the file. That's it.

When would you want to use `pubbo`? Almost never. Your static content will be available over a CDN, or hosted locally on a Web server behind nginx, or in a cluster. But how do you access and maintain the CDN? How do you standup your Web server? How do you configure nginx? How do you connect to the cluster? That basic access information needs to live somewhere, and that basic access information does you no good if it's locked behind the very systems that it tells you how to access.

`pubbo` doesn't need any of that stuff. All `pubbo` needs is a path to content an a Unix socket.

Unix sockets are all you need, if you happen to be logged onto the Unix socket's host system. If not, you'll need remote access. Unix sockets play well with nginx, which supports Unix sockets natively. And nginx plays well with pretty much everything else. So any access pattern you want to support can be supported with pubbo.

And it's a fun name.
