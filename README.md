# MIME Types Package

A lightweight MIME type recognition package that automatically determines Content-Type based on file extensions.

## Features

- Supports 150+ common file type MIME type recognition
- Fast matching based on file extensions
- Fully compatible with nginx/mime.types standard
- Zero dependencies, simple to use

## Supported File Types

Including but not limited to:

- Text files (HTML, CSS, XML, Plain Text, etc.)
- Image files (JPEG, PNG, GIF, SVG, WebP, AVIF, etc.)
- Font files (WOFF, WOFF2, EOT)
- Application files (JS, JSON, PDF, ZIP, WASM, etc.)
- Office documents (Word, Excel, PowerPoint, OpenDocument, etc.)
- Audio files (MP3, OGG, MIDI, etc.)
- Video files (MP4, WebM, MPEG, 3GPP, etc.)

## Usage
```bash
go get -u github.com/fanqie/mimeTypes
```
```go
import (
    "github.com/fanqie/mimeTypes"
)
// Get MIME type for a file
mimeType := mime.GetMimeType(".jpg") // returns "image/jpeg"
```
## Use on download file 
```go
http.HandleFunc("/download", func(w http.ResponseWriter, r http.Request) {
filename := "example.pdf"
w.Header().Set("Content-Type", mime.GetMimeType(filepath.Ext(filename)))
// ... handle file download
})
```
