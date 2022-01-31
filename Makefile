release:
	mkdir bin; \
        go build -o ./bin/salt ./salt.go; \
        RAND=$(./bin/salt 32); \
        echo $$RAND; \
        sed -i "s/DEFAULTPASSWORD/$$RAND/" values.go; \
	go build -o ./bin/locker ./locker.go ./main.go ./unlocker.go ./values.go; \
        # sed -i "s/$$RAND/DEFAULTPASSWORD/" values.go; \
        rand=0; 

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