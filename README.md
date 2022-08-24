# ARMSim Simulator Starter

This project provides starting to code to write a simulator in Golang using an Electron frontend.  Electron is a 
popular frontend choice that allows you to write HTML, JS and CSS frontends that are cross platform and can communicate
with a variety of backend technologies, like Golang.

This starter code project includes a simple frontend that uses Twitter's Bootstrap CSS framework, JQuery and websockets
to communicate with a Golang backend.  The course will provide a crash course in Golang.  Familiarity with HTML and
JS basics is expected.

## Getting Started

Install Go 1.19 on your platform.

For Windows, [download here](https://go.dev/dl/go1.19.windows-amd64.msi).  For Mac OS X, 
[intel download here](https://go.dev/dl/go1.19.darwin-amd64.pkg) and 
[M1 / M2 download here](https://go.dev/dl/go1.19.darwin-arm64.pkg). For Linux, use your favorite package manager or 
install from source.

After installing Golang, sign up for the [Github Education Pack](https://education.github.com/pack), and once approved,
sign up for the JetBrains product pack.  This will give you access to the Goland IDE, a full featured Golang IDE that
will be used in class.

For more information about the electron portion of the application, refer to the 
[module documentation](https://github.com/asticode/go-astilectron). This project also comes preloaded with a set of 
[electron helpers](https://github.com/asticode/js-toolbox) preinstalled and configured.

> Note: this project uses a bundler to build, you can find [documentation here](https://github.com/asticode/go-astilectron-bundler).

Import this project into Goland, then run `go mod download` in a terminal window of your choice. Add `$GOPATH/bin` to your
`PATH`. Then run `go mod tidy`, followed by `go install github.com/asticode/go-astilectron-bundler/astilectron-bundler` followed by `astilectron-bundler`.

> Note: To run this project in the future, only `astilectron-bundler` is needed.

An executable that you can run will be placed in the `output/` folder.

## Project Overview

Within this starter project, you'll find four files that you should be concerned about.  The first is `go.mod`. This
file contains a list of all the libraries in use by the project.  If you need or want to use any additional libraries
in the project, add them here to make them available running `go mod tidy` and `go mod download` after.  The next is 
`main.go`.  This file contains the bulk of the logic to get the UI running.  You shouldn't have to modify the main
function except to add new command line parameters or to adjust the window size. The `handleMessages` function you will 
update as time goes on.  This function is how messages are received from the UI.  At the current moment, it receives 
only the `load.file` message. Note when defining types in Golang, the first letter needs to be capitalized in order to 
make it public (which is very different from langauges you may be used to).  The next file is `index.html` which is the 
definition of your user interface.  There are several libraries preimported for displaying notifications and loading
animations if you need them.  The final file is `app.js`.  This file contains the brunt of the javascript logic. In
particular pay attention to the click handler which runs when a button is clicked in the UI and the `loadFile` method.
The `loadFile` handler is not traditional javascript but uses `astilectron` to send messages to your Golang backend.
I recommend the body sent to your backend be JSON for ease of debugging.

The application currently has 2 supported command line arguments `-d` and `-h`.  `-d` will enable debug mode in your app
allowing you to select debug from the file dropdown to get a javascript console you can see errors with. Other than
using Goland to attach to a running process and debug, there is no easy way to debug your backend. For instructions, 
on how to do that, refer to the [documentation](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html).
Note that print debugging is always an option, and you can see an example present in `handleMessages`. `-h` is auto 
constructed.

I recommend keeping `main.go` dedicated to handling messages and the UI.  Put all other logic in other files as 
appropriate.  The first two days of class will be a crash course in Golang so in the words of Scar from Lion King, "Be 
Prepared".
