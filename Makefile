#locker (https://github.com/X3NOOO/locker)
#Copyright (C) 2022 X3NO <X3NO@disroot.org> [https://X3NO.ct8.pl] [https://github.com/X3NOOO]
#
#This program is free software: you can redistribute it and/or modify
#it under the terms of the GNU General Public License as published by
#the Free Software Foundation, either version 3 of the License, or
#(at your option) any later version.
#
#This program is distributed in the hope that it will be useful,
#but WITHOUT ANY WARRANTY; without even the implied warranty of
#MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#GNU General Public License for more details.
#
#You should have received a copy of the GNU General Public License
#along with this program.  If not, see <http://www.gnu.org/licenses/>.

release:
	mkdir bin; \
        go run salt.go l 64 values.go; \
	go build -o ./bin/locker ./locker.go ./main.go ./unlocker.go ./values.go; \
        go run salt.go u 64 values.go

windows:
	mkdir bin; \
        go run salt.go l 64 values.go; \
	GOOS=windows go build -o .\bin\locker.exe .\locker.go .\main.go .\unlocker.go .\values.go; \
        go run salt.go u 64 values.go

install:
	if test `whoami` != "root" ; \
        then \
                echo "You are not root. Installing to user path ($$HOME/.local/bin/)"; \
				cp ./bin/locker $$HOME/.local/bin/; \
        else \
                echo "Installing to /usr/local/bin/" ; \
                cp ./bin/locker /usr/local/bin/; \
        fi

uninstall:
		if test `whoami` != "root" ; \
        then \
                echo "You are not root. Removing from user path ($$HOME/.local/bin/)"; \
				rm $$HOME/.local/bin/locker \
        else \
                echo "Removing from /usr/local/bin/" ; \
            	rm /usr/local/bin/locker; \
        fi
