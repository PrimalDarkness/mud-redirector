mud-redirector - notify your users that your MUD has moved
==============

mud-redirector is a dual notification and proxy system for letting your
players know your new MUD address (as long as you still have access to the old).

Usage
-----

```text
Usage: mud-redirector --listen=STRING --mud-name=STRING <command>

A simple notification and/or proxy service for redirecting MUD connections.

Flags:
      --help               Show context-sensitive help.
  -v, --verbose=INT        Tweak the verbosity of the logs.
  -l, --listen=STRING      The address for the service to listen on (0.0.0.0:5000)
  -n, --mud-name=STRING    The name of your MUD.

Commands:
  redirect --listen=STRING --mud-name=STRING <redirect-to>
    Proxy connecting user to the MUD (warning, IP information will be hidden).

  banner --listen=STRING --mud-name=STRING <redirect-to>
    Display a banner to the connecting user and disconnect immediately.

Run "mud-redirector <command> --help" for more information on a command.
```
