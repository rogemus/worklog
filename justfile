setup:
  @echo "Downloading dependencies..."
  go get -v ./...

build: setup
  @echo "Builing app..."
  go build -o tempocli 
  @echo "Build Successfull 🎉!"

install: setup build
  @echo "Installing..."
  go install
  @echo "Install Successfull 🎉!"
  @echo "Run 'tempocli help' to see all availabe options.\n"

uninstall:
  @echo "Uninstalling..."
  rm $HOME/go/bin/tempocli
