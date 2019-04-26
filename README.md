# Hetzner Cloud Pulumi Provider

## Usage example (TypeScript)

```typescript
const timestamp = Date.now();

const sshKey = new hcloud.SshKey("fleet-v2", {
  name: "fleet-v2",
  publicKey: "my-ssh-key"
});

const server = new hcloud.Server("nginx", {
  name: "nginx",
  image: "ubuntu-18.04",
  serverType: "cx11",
  userData: `#!/bin/bash
apt update
apt install nginx -y
echo "Hello. Installation timestamp ${timestamp}" > /var/www/html/index.html
service nginx start
`
});

export { server };
```
