# Certifiable

This is just a small little utility to help inject root certificates into your Go Application.

It was created for inexpensive retro handheld devices that have Wi-Fi but run into issues connecting to services hosted on
HTTPS due to not having a root certificate store shipped on device.

Import with side effects and you are done!

```go
package main

import (
    _ "github.com/UncleJunVIP/certifiable"
)
```

> [!NOTE]
> These certificates are pulled from Mozilla using the `download_certs.sh` script (it's just a curl).