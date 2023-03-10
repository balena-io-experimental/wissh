# TODO

* Allow to copy the check results to the clipboard. Especially important for the
  details portion.
* Differences in balenaOS versions may cause some checks to fail (e.g., because
  of differences in the supported command-line options for certain commands).
  Would be nice to test a variety of OS versions.

## Checks we may want to add

* `nc -w 5 -G 1 cloudlink.balena-cloud.com 443 && echo "Reachable." || echo "Not
  reachable."` Or an equivalent to this that works (I am not sure this does).
* IPv6 vs IPv4 variations of our pings to backend. Like `ping -4
  api.balena-cloud.com`.
* Any relevant DNS checks?
* Is it worth pinging some non-balena URLs (say, `google.com`)?
* NetworkManager configs
    * For example, I've seen bad attempts at configuring a static IP in support.
* `openssl s_client -connect api.balena-cloud.com:443`, to [check for Deep
  Packet Inspection
  firewalls](https://docs.balena.io/learn/more/masterclasses/device-debugging/#641-deep-packet-inspection).
    * Actually a tricky thing to check. IIUC, we'd need to compare the results
      with running the same command on a computer that is known to not be behind
      a DPI firewall.
