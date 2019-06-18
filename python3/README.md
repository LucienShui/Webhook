# Python3 Webhook

## Deploy

```bash
git clone https://github.com/LucienShui/Webhook.git
cd Webhook
cp -r python3 /usr/local/python3_webhook
ln -s /usr/local/python3_webhook/python3_webhook.service /lib/systemd/system
systemctl daemon-reload
```

Then, you should put an shell script `main.sh` into `/usr/local/python3_webhook`.

## Usage

```bash
systemctl start/stop/enable/disable python3_webhook
```

Access `<address>:<port>`, then `/usr/local/python3_webhook/main.sh` would be executed.

`<port>` is `10086` for default.
