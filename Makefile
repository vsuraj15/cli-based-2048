GONEW = go
GOCMD=$(GONEW) build

COMMON_ENV=CGO_ENABLED=0
Twenty48_MAIN=cmd/twenty48/main.go

# Set OS
LINUX_ENV=$(COMMON_ENV) GOOS=linux
WINDOWS_ENV=$(COMMON_ENV) GOOS=windows
MAC_ENV_AMD64=$(COMMON_ENV) GOOS=darwin GOARCH=amd64
MAC_ENV_ARM64=$(COMMON_ENV) GOOS=darwin GOARCH=arm64

# Create binaries
Twenty48_BIN_LINUX=target/twenty48-linux
Twenty48_BIN_WINDOWS=target/twenty48-windows.exe
Twenty48_BIN_MACOS_AMD64=target/amd64/twenty48-macos
Twenty48_BIN_MACOS_ARM64=target/arm64/twenty48-macos

# Create target
CREATE_TARGET=mkdir -p target
CREATE_AMD64_TARGET=mkdir -p target/amd64
CREATE_ARM64_TARGET=mkdir -p target/arm64

# Get GIT COMMIT and VERSION
GIT_COMMIT=$(shell git rev-list -1 HEAD)
Twenty48_VERSION=$(shell cat twenty48-version.txt)

# Build command
BUILD_CMD=$(GOCMD) -ldflags="-X 'loconav.com/projects/version.GitCommit=$(GIT_COMMIT)' -X 'loconav.com/projects/version.Version=$(VERSION)'" -o

# Setup
SETUP_LINUX=$(CREATE_TARGET) && $(LINUX_ENV) $(BUILD_CMD)
SETUP_WINDOWS=$(CREATE_TARGET) && $(WINDOWS_ENV) $(BUILD_CMD)
SETUP_MAC_AMD64=$(CREATE_TARGET) && $(CREATE_AMD64_TARGET) && $(MAC_ENV_AMD64) $(BUILD_CMD)
SETUP_MAC_ARM64=$(CREATE_TARGET) && $(CREATE_ARM64_TARGET) && $(MAC_ENV_ARM64) $(BUILD_CMD)

# All
all: twenty48

twenty48:
	$(eval VERSION=$(Twenty48_VERSION))
	$(SETUP_LINUX) $(Twenty48_BIN_LINUX) $(Twenty48_MAIN)
	$(SETUP_WINDOWS) $(Twenty48_BIN_WINDOWS) $(Twenty48_MAIN)
	$(SETUP_MAC_AMD64) $(Twenty48_BIN_MACOS_AMD64) $(Twenty48_MAIN)
	$(SETUP_MAC_ARM64) $(Twenty48_BIN_MACOS_ARM64) $(Twenty48_MAIN)

clean:
	rm -rf target
