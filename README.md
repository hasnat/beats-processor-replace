# beats-processor-replace plugin

[![Build Status](http://img.shields.io/travis/hasnat/beats-processor-replace.svg?style=flat-square)][travis]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: http://travis-ci.org/hasnat/beats-processor-replace
[godocs]: http://godoc.org/github.com/hasnat/beats-processor-replace
[releases]: https://github.com/hasnat/beats-processor-replace/releases

beats-processor-replace is a processor plugin for Elastic Beats that can replace
info in event.

## Installation and Usage

Build the plugin. Go plugins are only supported on Linux at the current time. They must be
compiled with the same Go version as the Beat it will be used with. Likewise this plugin
must be compiled against the same Beat codebase version as the Beat it will be used
with.

```
go build -buildmode=plugin
```
Run example
```bash
cd example
docker-compose up
```

Start a Beat with the plugin.

```
filebeat --plugin ./processor-replace-linux.so
```

If using docker, you can copy across pre-built plugin and add it to your entrypoint. Check Dockerfile

```
COPY --from=hasnat/beats-processor-replace /usr/local/plugins/processor-replace-linux.so /usr/local/plugins/processor-replace-linux.so
CMD ["/bin/sh", "-c", "'/usr/local/bin/docker-entrypoint -e --plugin /usr/local/plugins/processor-replace-linux.so'"]
```

Add the processor to your configuration file.

```
processors:
- replace:
    field: message
    target: replaced_message
    find: "\t"
    replace: ","
    regex: false
```

## Configuration Options

- **`field`**: Field to do replacement on.
- **`target`**: Where to write replaced value result, if not provided will replace value of `field`.
- **`find`**: Find, can be regex.
- **`replace`**: Replace.
- **`regex`**: Define if find expression is regex, default is false. As substitution by regex is slower.

## Example Output

```json
{
  "@timestamp": "2017-10-07T03:09:50.201Z",
  "@metadata": {
    "beat": "filebeat",
    "type": "doc",
    "version": "7.0.0-alpha1"
  },
  "source": "/some/log/file/messages",
  "offset": 68379,
  "message": "Message has tabs [a	b	c	d].",
  "beat": {
    "name": "host.example.com",
    "hostname": "host.example.com",
    "version": "7.0.0-alpha1"
  },
  "replaced_message": "Message has tabs [a,b,c,d]"
}
```

## Common errors
If getting errors like, it might be image is built don different architecture than its being run on
instead of copy so file build it on same architecture

e.g. set build context to https://github.com/hasnat/beats-processor-replace.git
```bash
fatal error: runtime: no plugin module data

goroutine 1 [running]:
runtime.throw(0x16dbd59, 0x1e)
...
main.main()
	/go/src/github.com/elastic/beats/filebeat/main.go:18 +0x2f fp=0xc42022df80 sp=0xc42022df58 pc=0x146b56f
runtime.main()
	/usr/local/go/src/runtime/proc.go:195 +0x226 fp=0xc42022dfe0 sp=0xc42022df80 pc=0xae2df6
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2337 +0x1 fp=0xc42022dfe8 sp=0xc42022dfe0 pc=0xb12551
goroutine 36 [syscall]:
os/signal.signal_recv(0x17111b0)
	/usr/local/go/src/runtime/sigqueue.go:131 +0xa6
os/signal.loop()
	/usr/local/go/src/os/signal/signal_unix.go:22 +0x22
created by os/signal.init.0
	/usr/local/go/src/os/signal/signal_unix.go:28 +0x41

```

## References & Thanks
Big thanks to [Andrew Kroh](https://github.com/andrewkroh) for example plugins implementation

https://github.com/andrewkroh/beats-processor-fingerprint

https://github.com/s12v/awsbeats

https://github.com/elastic/beats/issues/6760

