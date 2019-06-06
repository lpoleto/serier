# Serier
Serier is yet another application to rename TV episode files using data from https://www.thetvdb.com.
The main purpose is to rename from a generic name such as `tv.show.05.04.sbrubles.ddl.mkv` to `TV Show - S05E04 - Episode Name.mkv` where the episode title is obtained from The TVDB.

### Yet another series renamer?
Because all other file renamers I've found required too many steps to install or have configurations that are too complex or both. They mostly work great, yes, but I wanted to make something simpler yet powerful.

## Installing
To install Serier you only need to [download](https://github.com/lpoleto/serier/releases) the executable file specific to your system (Windows, macOS, or Linux) and copy to whatever directory you want.
To make it easier to run, though, you might want to copy it to a directory already in your PATH.

## Basic usage
In the terminal, navigate to the directory where Serier was installed (if it's not in a directory in your PATH) and type:

```bash
serier -s [series name] -p [path to the files to be renamed]
``` 

Follow the instructions that will be presented.

The parameter `-p` can be ommitted; in this case, the current directory where Serier was called from will be used. For exemple, if you navigated to `/home/User/Joe/Files`, Serier will look for the series files in that directory.

### Supported file names
Serier will look for specific patterns in the scanned file names in order to determine the Season and Episode number a file refers to. At this point, it supports the following patters:

* SXXEYY (case insensitive; i.e., S or s and E or e or any combination)
* XX.YY

Where XX is the season number and YY is the episode number.

### File output
Right now the output file name is hardcoded and has this format:

* TV Show Name - SXXEYY - Episode Name.ext

## Questions or issues?
If you have questions about Serier of if you found errors with it, just create an [issue](https://github.com/lpoleto/serier/issues). I'll be glad to help.

## Contributing
If you would like to help, pick an issue (create one with your change if it doesn't exist) and send a PR.
