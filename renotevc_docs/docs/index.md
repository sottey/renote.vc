# renote.vc
---
renote.vc is a tool for command line fans that allows you to track your tasks. It is a full featured todo manager with tagging, projects, recurring tasks and much more, all stored in a JSON file so it is super portable and tooling new apps for the data format is super easy.

## Getting Started
To learn about available renote.vc commands, see the [usage page](https://renote.vc/usage/)

## Features
* Create nodes (files or folders)
* Store markdown, text, html and edit using your default editor
* Store other file types which will open in their default pp (pdf, docxx, png, etc.)
* Local filesystem AND Filebase storage options
* Quick UI (still on the command line) to search for and select notes
* Auto generate shell completion files and bash, zsh and more
* Edit and rename notes

## Future Plans
* Better listing UI
* Theming
* Markdown rendering

## Downloading
Find the installation for your device in the releases section of the GitHub project [here](https://github.com/sottey/renote.vc/releases)

## Building

renote.vc is available for most configurations. Go to the relase page, download the proper archive for your device. 

Once the file has downloaded, extract the binary and put it somewhere that is accessible from your terminal.

## Contribute

Clone locally and run
```
go build main.go
```

OR (using rake)
```
rake build
```

## Shoutouts

* [notya](https://github.com/insolite-dev/notya) - HUGE thanks to everyone who created [notya](https://github.com/insolite-dev/notya), an amazing note tool!


## License

renote.vc uses the MIT License. Please see the [renote.vc license](https://github.com/sottey/renote.vc/blob/main/LICENSE) for details