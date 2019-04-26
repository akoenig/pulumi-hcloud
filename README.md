# Hetzner Cloud Pulumi Provider

## Usage example (TypeScript)

```typescript
import * as hcloud from "@pulumi/hcloud";

const timestamp = Date.now();

const sshKey = new hcloud.SshKey("my-key", {
  name: "my-key",
  publicKey: "<contents-of-public-ssh-key>"
});

const server = new hcloud.Server(
  "nginx",
  {
    name: "nginx",
    image: "ubuntu-18.04",
    serverType: "cx11",
    sshKeys: [sshKey.name],
    userData: `#!/bin/bash
  apt update
  apt install nginx -y
  echo "Hello. Installation timestamp ${timestamp}" > /var/www/html/index.html
  service nginx start
  `
  },
  { dependsOn: [sshKey] }
);

export { server };
```
