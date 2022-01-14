release:
	mkdir bin; \
	go build -o ./bin/locker ./*.go;

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