package main

import "github.com/dentych/goday/apprication/downloader"

func main() {
	downloader.DownloadFile("test.rebex.net:21", "demo", "password", "readme.txt")
}
