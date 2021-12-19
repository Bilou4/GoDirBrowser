LINUX=linux
WINDOWS=windows
ARCH64=amd64
ARCH32=386
SOURCE=.
FLAGS=-ldflags "-s -w" -trimpath

all: GoDirBrowser_bin32.out GoDirBrowser_bin64.out GoDirBrowser_exe32.exe GoDirBrowser_exe64.exe

GoDirBrowser_bin64.out: $(SOURCE)
	GOOS=$(LINUX) GOARCH=$(ARCH64) go build -o $@ $(FLAGS) $<

GoDirBrowser_bin32.out: $(SOURCE)
	GOOS=$(LINUX) GOARCH=$(ARCH32) go build -o $@ $(FLAGS) $<

GoDirBrowser_exe64.exe: $(SOURCE)
	GOOS=$(WINDOWS) GOARCH=$(ARCH64) go build -o $@ $(FLAGS) $<

GoDirBrowser_exe32.exe: $(SOURCE)
	GOOS=$(WINDOWS) GOARCH=$(ARCH32) go build -o $@ $(FLAGS) $<

clean:
	rm *.exe *.out


