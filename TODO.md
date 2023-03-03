# TODO

## Checks to consider

* IPv6 vs IPv4 variations of our pings to backend. Like `ping -4
  api.balena-cloud.com`.
* Any relevant DNS checks?
* Is it worth pinging some non-balena URLs (say, `google.com`)?
* NetworkManager configs
    * For example, I've seen bad attempts at configuring a static IP in support.
* `openssl s_client -connect api.balena-cloud.com:443`, to [check for Deep
  Packed Inspection
  firewalls](https://docs.balena.io/learn/more/masterclasses/device-debugging/#641-deep-packet-inspection).
