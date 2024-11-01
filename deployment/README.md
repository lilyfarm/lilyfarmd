# Deployment

The deployment module facilitates the deployment of the lilyfarm service on Linux
machines using `systemd`.

To get started, have all your API keys ready and run

```
sudo python3 deploy.py
```

You might need `sudo` with the above.

The CLI will guide you through deploying your service on Linux machines. It
will create a `lilyfarmd.service` to be run by `systemd`.
