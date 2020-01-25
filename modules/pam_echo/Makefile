UNAME := $(shell uname -s)

# PAM_POLICY_PATH is where the PAM policy for PAM-aware applications is
# defined.
PAM_POLICY_PATH = /etc/pam.d/

# PAM_MODULE_PATH is determined by the OS.
PAM_MODULE_PATH = /lib/x86_64-linux-gnu/security/
ifeq ($(UNAME),Darwin)
    PAM_MODULE_PATH = /usr/local/lib/pam/
endif

# LD and LD_FLAGS controls the linker and linker flags to use and are
# determined by the OS.
LD = ld
LD_FLAGS = -lpam --shared -x
ifeq ($(UNAME),Darwin)
    LD = clang
    LD_FLAGS = -lpam -shared
endif

all: pam_echo.so

install:
	mkdir -p $(PAM_MODULE_PATH)
	cp pam_echo.so $(PAM_MODULE_PATH)
	sudo cp policy/teleport-acct-failure $(PAM_POLICY_PATH)
	sudo cp policy/teleport-session-failure $(PAM_POLICY_PATH)
	sudo cp policy/teleport-success $(PAM_POLICY_PATH)

pam_echo.so: pam_echo.o
	$(LD) $(LD_FLAGS) -o pam_echo.so pam_echo.o
	chmod 644 pam_echo.so

pam_echo.o: clean pam_echo.c
	gcc -fPIC -c pam_echo.c

clean:
	rm -f pam_echo.o pam_echo.so
