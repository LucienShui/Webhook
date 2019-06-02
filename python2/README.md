# Python2 Webhook

## Deploy

```bash
git clone https://github.com/LucienShui/Webhook.git
cd Webhook
cp -r python2 /usr/local/python2_webhook
ln -s /usr/local/python2_webhook/python2_webhook.service /lib/system/system
systemctl daemon-reload
```

Then, you should put an shell script `main.sh` into `/usr/local/python2_webhook`.

## Usage

```bash
systemctl start/stop/enable/disable python2_webhook
```

Access `<address>:<port>`, then `/usr/local/python2_webhook/main.sh` would be executed.

`<port>` is `10086` for default.
