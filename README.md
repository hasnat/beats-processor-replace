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

Start a Beat with the plugin.

```
filebeat -e --plugin ./processor-replace-linux-amd64.so
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

## References & Thanks
Big thanks to [Andrew Kroh](https://github.com/andrewkroh) for example plugins implementation

https://github.com/andrewkroh/beats-processor-fingerprint

https://github.com/s12v/awsbeats

https://github.com/elastic/beats/issues/6760

