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
	cp -f hwsmysqlcleard $(RCLOCALDIR)
	cp -f hws-rc.local $(RCLOCAL)
	sh /etc/hws-rc.local.d/hwsmysqlcleard


uninstall:
	-killall hwsmysqlclear
	systemctl disable hws-rc-local
	systemctl stop hws-rc-local
	rm -f $(PREFIX)
	rm -fr $(RCLOCALDIR)
	rm -f $(SERVICE)
	rm -f $(RCLOCAL)


.PHONY: install uninstall
