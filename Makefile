release:
	mkdir bin; \
        go run salt.go l 64 values.go; \
	go build -o ./bin/locker ./locker.go ./main.go ./unlocker.go ./values.go; \
        go run salt.go u 64 values.go

windows:
	mkdir bin; \
        go run salt.go l 64 values.go; \
	go build -o .\bin\locker.exe .\locker.go .\main.go .\unlocker.go .\values.go; \
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