# Goday
Best day ever

![](https://www.websequencediagrams.com/cgi-bin/cdraw?lz=ZG93bmxvYWRlciAtPisgZnRwOiByZXF1ZXN0IGZpbGUKZnRwIC0tPi0gACAKOiByZWNlaXZlABwGCgA1D2FyY2hpdmU6IHN0ZWFtAEAGAA0HIC0-ABQLYXZlcwBgBSB0byBkaXNrABwLKyBoYXNoZXI6IGdlbmVyYXRlAAwFCgAQBgCBDAYAYAlyZXR1cm4AGwYANw1qYW1lcwBkBgBLBSBhbmQgbG9jCgAUBQCBbQVkYgCBBAYKZGIAgWIGAC0HZG9uZQoAIAgAYQ0AFQUAgUkIAIIKEgA0BQ&s=magazine)

## Diagram code
downloader ->+ ftp: request file  
ftp -->- downloader: receive file  
downloader ->+ archive: steam file  
archive -> archive: saves file to disk  
archive ->+ hasher: generate hash  
hasher -->- archive: return hash  
archive ->+ james: save hash and loc  
james ->+ db: save  
db -->- james: done  
james -->- archive: done  
archive  -->- downloader: done  
