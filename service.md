# Service control

> Linux (systemctl)

```Makefile
SERVICE=

reload:
	systemctl daemon-reload

start:
	systemctl start $(SERVICE)

stop:
	systemctl stop $(SERVICE)

restart:
	systemctl restart $(SERVICE)


enable:
	systemctl enable $(SERVICE)

disable:
	systemctl disable $(SERVICE)

status:
	systemctl status $(SERVICE)

log:
	journalctl -u $(SERVICE) -xe

cat:
	systemctl cat $(SERVICE)

list:
	systemctl list-units
```
