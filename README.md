# MIME Types Package

A lightweight MIME type recognition package that automatically determines Content-Type based on file extensions.

## Features

- Supports 109 common file type MIME type recognition
- Fast matching based on file extensions
- Fully compatible with **_`nginx/mime.types`_** standard
- Zero dependencies, simple to use
- Comprehensive coverage of unit testing(you try run comand `go test  -v  test\mime_test.go`)

## Supported File Types

- **_The same as nginx/mime.types_**
- Including but not limited to:
  - **_`html,htm,shtml,css,xml,gif,jpeg,jpg,js,atom,rss,mml,txt,jad,wml,htc,avif,png,svg,svgz,tif,tiff,wbmp,webp,ico,jng,bmp,woff,woff2,jar,war,ear,json,hqx,doc,pdf,ps,eps,ai,rtf,m3u8,kml,kmz,xls,eot,ppt,odg,odp,ods,odt,pptx,xlsx,docx,wmlc,wasm,7z,cco,jardiff,jnlp,run,pl,pm,prc,pdb,rar,rpm,sea,swf,sit,tcl,tk,der,pem,crt,xpi,xhtml,xspf,zip,bin,exe,dll,deb,dmg,iso,img,msi,msp,mid,midi,kar,mp3,ogg,m4a,ra,3gpp,3gp,ts,mp4,mpeg,mpg,mov,webm,flv,m4v,mng,asx,asf,wmv,avi`_**

## Usage
```bash
go get -u github.com/fanqie/mimeTypes
```
```go
import (
    mime "github.com/fanqie/mimeTypes/pkg"
)
// Get MIME type for a file
mimeType := mime.GetMimeType(".jpg") // returns "image/jpeg"
// or
mimeType := mime.GetMimeType("jpg") // returns "image/jpeg"
```
## Ignore case
```go
mimeType := mime.GetMimeType(".JPG") // returns "image/jpeg"
mimeType := mime.GetMimeType("JPG") // returns "image/jpeg"
```
## Get MIME type from URL
```go
mimeType := mime.GetMimeTypeFromURL("https://example.com/example.jpg")
// return "image/jpeg"
```
## Get MIME type from file path
```go
mimeType := mime.GetMimeTypeFromPath("example.jpg")
// return "image/jpeg"
```
## Get MIME type from file
```go
func GetMimeTypeFromFileExample() (string, error) {
	
    file, err := os.Open("./types.json")
        if err != nil {
            fmt.Println(err)
            return
        }
    defer func(file *os.File) {
        err := file.Close()
            if err != nil {
            fmt.Printf("error:%s", err)
        }
    }(file)
    return GetMimeTypeFromFile(file)
    // return "application/json"
}
```
## Use on download file 
```go
http.HandleFunc("/download", func(w http.ResponseWriter, r http.Request) {
filename := "example.pdf"
w.Header().Set("Content-Type", mime.GetMimeType(filepath.Ext(filename)))
// ... handle file download
})
```
##
## Get file extension from MIME type
```go
mime.GetExtFromMimeType("application/json")
// return ".json"
```
## Get file extension from MIME type Name
```go
mime.GetExtFromMimeTypeName("image/jpeg")
// return "jpg"
```
