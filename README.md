# Plavatar
A stateless microservice that returns **pla**ceholder **avatar**s (=plavatars).

![docs/assets/readme-demo.png](docs/assets/readme-demo.png)

## API Endpoints
* `baseurl:port/laughing/<size>/<name>`
* `baseurl:port/smiley/<size>/<name>`
* `baseurl:port/happy/<size>/<name>`
* `baseurl:port/gradient/<size>/<name>`
* `baseurl:port/marble/<size>/<name>`
* `baseurl:port/solid/<size>/<name>`
* `baseurl:port/pixel/<size>/<name>` (*currently only available as square*)

Without name:
* `baseurl:port/laughing/<size>` and so on

With query params:
* `baseurl:port/laughing/<size>/<name>?format=svg`
* `baseurl:port/laughing/<size>/<name>?format=svg&shape=square`
* `baseurl:port/laughing/<size>/<name>?shape=square` and so on

## Parameters
* `size` the image's size in pixels. has to be min 16, max 1024
* `name` **optional**, the random number generator seed to use. given the same name the same picture will be returned
### Query Params
* `format`**optional**, either png (default) or svg. svg returns the raw svg
* `shape` **optional**, either circle (default) or square.

## **If possible use format=SVG.** 
Not only is format=SVG extremely faster, it also saves you a lot of bandwidth and latency. (A generated SVG is only ~2% the size of a 512px PNG)

## Configuration
You can optionally supply a config file if you are not happy with the preset settings.  
By the default the program looks for a config file at `<running_folder>/config/plavatar.json`. If you want to use an
alternative location you can override this behaviour using the argument `--config <path_to_config>`. If there's neither
a config in the `config/` folder, nor you supply a path with `--config` the default configuration will be used.

### Default configuration file

```json
{
  "dimensions": {
    "min": 128,
    "max": 512
  },
  "webserver": {
    "gzip": true,
    "http": {
      "enabled": true,
      "host": "0.0.0.0",
      "port": 7331
    },
    "https": {
      "enabled": false,
      "host": "0.0.0.0",
      "port": 7332,
      "cert": "testing.crt",
      "key": "testing.key"
    }
  },
  "caching": {
    "enabled": true,
    "ttl": "8h"
  },
  "metrics": {
    "enabled": false,
    "auth": {
      "enabled": true,
      "username": "",
      "password": ""
    }
  }
}
```

## Testing
To run the go tests, use `go test -v ./...` in the root directory of the project.

To generate a self-signed certificate for testing purposes you can
use `openssl req -newkey rsa:4096 -x509 -sha256 -days 3650 -nodes -out testing.crt -keyout testing.key`

For benchmarking, you can use the provided [k6 script](https://github.com/grafana/k6) under `scripts/k6_plavatar_benchmark.js`.
