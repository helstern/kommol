# Not-found handling

Decided by the user, 2026-10-25.

## How it works (v0.3.0)

- 404 pages live in the bucket.
- For a missing object kommol serves, with status 404:
  1. the bucket's own not-found page (its website setting `NotFoundPage`; read once per bucket and cached), else
  2. the page given on the command line, `-not-found-page <path in the bucket>`.
- With neither, kommol answers an empty 404 (`Content-Length: 0`; v0.4.0) and the proxy in front
  serves the error page.
- Any other read error stays 500.

## Todo

- [x] haproxy on the gateway (control.helstern.org, weblb role's `haproxy.cfg` template) serves the error page when
  kommol answers a plain 404. Deployed 2026-11-11 (image `gateway-control-helstern-org-20261111`); tested by
  control.helstern.org `srv/weblb/test/haproxy.bats`.
- [ ] Gateway on the empty 404 (user decision 2026-11-11). Done: v0.4.0 tagged locally; control.helstern.org
  `e39b754` (pin 0.4.0, rule on `res.hdr(content-length) -m str 0`, bats 3/3). Open: publish the v0.4.0 release
  (user), gateway image build, VM replace (needs a yes), check prod.
