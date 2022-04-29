# go-copy-file-progress
Show progress of downloading file

How to use:
=
Run:
```
go get github.com/shapito27/go-copy-file-progress
```

Code example:
```
package main

import downloader "github.com/shapito27/go-copy-file-progress"

func main() {
	//first file
	downloader.DownloadFile("https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb", "./")
	//second file
	downloader.DownloadFile("https://wordpress.org/latest.tar.gz", "./")
	//third file
	downloader.DownloadFile("https://thumbs.dreamstime.com/z/die-fantacy-immage-woman-s-hand-draws-31994627.jpg", "./")
}
```

