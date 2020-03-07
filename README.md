# kommol [![Build Status](https://travis-ci.org/helstern/kommol.svg?branch=master)](https://travis-ci.org/helstern/kommol)
reverse proxy for gcp storage buckets

## Installation

Download the tar archive for your operating system from the latest release in github: https://github.com/helstern/kommol/releases/latest

Extract the contents into a folder of your choosing

## Usage

You will need some exported service credentials that will allow the proxy to read items from your buckets. Read the [getting started with authentication for Google Cloud Platform guide](https://cloud.google.com/docs/authentication/getting-started) to set up and download service account credentials

The permissions required by the proxy are:

- storage.buckets.get
- storage.objects.get

```
    kommol -bind <ip:port> -gcp.credentials <path to service credentials> [--log-level <info|warn|debug>]
```
