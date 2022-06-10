PREFIX := /usr/local/bin/hwsmysqlclear
SERVICE := /etc/systemd/system/hws-rc-local.service
RCLOCALDIR := /etc/hws-rc.local.d
RCLOCAL := /etc/hws-rc.local

install:
	cp -f hws-rc-local.service $(SERVICE)
	systemctl enable hws-rc-local
	systemctl start hws-rc-local
	mkdir -p $(RCLOCALDIR)
	cp -f hwsmysqlclear $(PREFIX)
	cp -f hws-rc.local $(RCLOCAL)
	cp -f hwsmysqlcleard $(RCLOCALDIR)

uninstall:
	systemctl disable hws-rc-local
	systemctl stop hws-rc-local
	rm -f $(PREFIX)
	rm -f $(SERVICE)
	rm -f $(RCLOCAL)
	rm -fr $(RCLOCALDIR)


.PHONY: install uninstall
