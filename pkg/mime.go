// Package mime provides a way to get the mime type of a file.
package mime

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	// text
	TextHTML      = "text/html"
	TextCSS       = "text/css"
	TextXML       = "text/xml"
	TextPlain     = "text/plain"
	TextMathML    = "text/mathml"
	TextComponent = "text/x-component"
	TextVndJ2ME   = "text/vnd.sun.j2me.app-descriptor"
	TextVndWapWML = "text/vnd.wap.wml"

	// image
	ImageGIF        = "image/gif"
	ImageJPEG       = "image/jpeg"
	ImagePNG        = "image/png"
	ImageSVG        = "image/svg+xml"
	ImageTIFF       = "image/tiff"
	ImageWEBP       = "image/webp"
	ImageICO        = "image/x-icon"
	ImageJNG        = "image/x-jng"
	ImageBMP        = "image/x-ms-bmp"
	ImageAVIF       = "image/avif"
	ImageVndWapWBMP = "image/vnd.wap.wbmp"

	// font
	FontWOFF  = "font/woff"
	FontWOFF2 = "font/woff2"
	FontEOT   = "application/vnd.ms-fontobject"

	// application
	AppJS             = "application/javascript"
	AppJSON           = "application/json"
	AppPDF            = "application/pdf"
	AppZIP            = "application/zip"
	AppWASM           = "application/wasm"
	AppAtomXML        = "application/atom+xml"
	AppRSSXML         = "application/rss+xml"
	AppJavaArchive    = "application/java-archive"
	AppMacBinHex40    = "application/mac-binhex40"
	AppPostScript     = "application/postscript"
	AppRTF            = "application/rtf"
	AppMPEGURL        = "application/vnd.apple.mpegurl"
	AppKML            = "application/vnd.google-earth.kml+xml"
	AppKMZ            = "application/vnd.google-earth.kmz"
	AppWMLC           = "application/vnd.wap.wmlc"
	App7Z             = "application/x-7z-compressed"
	AppCocoa          = "application/x-cocoa"
	AppJavaDiff       = "application/x-java-archive-diff"
	AppJNLP           = "application/x-java-jnlp-file"
	AppMakeSelf       = "application/x-makeself"
	AppPerl           = "application/x-perl"
	AppPilot          = "application/x-pilot"
	AppRAR            = "application/x-rar-compressed"
	AppRPM            = "application/x-redhat-package-manager"
	AppSEA            = "application/x-sea"
	AppShockwaveFlash = "application/x-shockwave-flash"
	AppStuffIt        = "application/x-stuffit"
	AppTCL            = "application/x-tcl"
	AppX509           = "application/x-x509-ca-cert"
	AppXPInstall      = "application/x-xpinstall"
	AppXHTML          = "application/xhtml+xml"
	AppXSPF           = "application/xspf+xml"

	// office
	AppMSWord  = "application/msword"
	AppMSExcel = "application/vnd.ms-excel"
	AppMSPPT   = "application/vnd.ms-powerpoint"
	AppODG     = "application/vnd.oasis.opendocument.graphics"
	AppODP     = "application/vnd.oasis.opendocument.presentation"
	AppODS     = "application/vnd.oasis.opendocument.spreadsheet"
	AppODT     = "application/vnd.oasis.opendocument.text"
	AppDocx    = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	AppXlsx    = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	AppPptx    = "application/vnd.openxmlformats-officedocument.presentationml.presentation"

	// binary
	AppOctetStream = "application/octet-stream"

	// audio
	AudioMIDI      = "audio/midi"
	AudioMP3       = "audio/mpeg"
	AudioOGG       = "audio/ogg"
	AudioM4A       = "audio/x-m4a"
	AudioRealAudio = "audio/x-realaudio"

	// video
	Video3GPP      = "video/3gpp"
	VideoMP2T      = "video/mp2t"
	VideoMP4       = "video/mp4"
	VideoMPEG      = "video/mpeg"
	VideoQuickTime = "video/quicktime"
	VideoWEBM      = "video/webm"
	VideoFLV       = "video/x-flv"
	VideoM4V       = "video/x-m4v"
	VideoMNG       = "video/x-mng"
	VideoMSASF     = "video/x-ms-asf"
	VideoMSWMV     = "video/x-ms-wmv"
	VideoMSVideo   = "video/x-msvideo"
)

var (
	mimeTypes = map[string]string{
		".html":  TextHTML,
		".htm":   TextHTML,
		".shtml": TextHTML,
		".css":   TextCSS,
		".xml":   TextXML,
		".txt":   TextPlain,
		".jad":   TextVndJ2ME,
		".wml":   TextVndWapWML,
		".htc":   TextComponent,

		".gif":  ImageGIF,
		".jpeg": ImageJPEG,
		".jpg":  ImageJPEG,
		".png":  ImagePNG,
		".svg":  ImageSVG,
		".svgz": ImageSVG,
		".tif":  ImageTIFF,
		".tiff": ImageTIFF,
		".webp": ImageWEBP,
		".ico":  ImageICO,
		".jng":  ImageJNG,
		".bmp":  ImageBMP,
		".avif": ImageAVIF,
		".wbmp": ImageVndWapWBMP,

		".woff":  FontWOFF,
		".woff2": FontWOFF2,
		".eot":   FontEOT,

		".js":      AppJS,
		".json":    AppJSON,
		".pdf":     AppPDF,
		".zip":     AppZIP,
		".wasm":    AppWASM,
		".atom":    AppAtomXML,
		".rss":     AppRSSXML,
		".jar":     AppJavaArchive,
		".war":     AppJavaArchive,
		".ear":     AppJavaArchive,
		".hqx":     AppMacBinHex40,
		".ps":      AppPostScript,
		".eps":     AppPostScript,
		".ai":      AppPostScript,
		".rtf":     AppRTF,
		".m3u8":    AppMPEGURL,
		".kml":     AppKML,
		".kmz":     AppKMZ,
		".wmlc":    AppWMLC,
		".7z":      App7Z,
		".cco":     AppCocoa,
		".jardiff": AppJavaDiff,
		".jnlp":    AppJNLP,
		".run":     AppMakeSelf,
		".pl":      AppPerl,
		".pm":      AppPerl,
		".prc":     AppPilot,
		".pdb":     AppPilot,
		".rar":     AppRAR,
		".rpm":     AppRPM,
		".sea":     AppSEA,
		".swf":     AppShockwaveFlash,
		".sit":     AppStuffIt,
		".tcl":     AppTCL,
		".tk":      AppTCL,
		".der":     AppX509,
		".pem":     AppX509,
		".crt":     AppX509,
		".xpi":     AppXPInstall,
		".xhtml":   AppXHTML,
		".xspf":    AppXSPF,

		".doc":  AppMSWord,
		".xls":  AppMSExcel,
		".ppt":  AppMSPPT,
		".odg":  AppODG,
		".odp":  AppODP,
		".ods":  AppODS,
		".odt":  AppODT,
		".docx": AppDocx,
		".xlsx": AppXlsx,
		".pptx": AppPptx,

		".bin": AppOctetStream,
		".exe": AppOctetStream,
		".dll": AppOctetStream,
		".deb": AppOctetStream,
		".dmg": AppOctetStream,
		".iso": AppOctetStream,
		".img": AppOctetStream,
		".msi": AppOctetStream,
		".msp": AppOctetStream,

		".mid":  AudioMIDI,
		".midi": AudioMIDI,
		".kar":  AudioMIDI,
		".mp3":  AudioMP3,
		".ogg":  AudioOGG,
		".m4a":  AudioM4A,
		".ra":   AudioRealAudio,

		".3gpp": Video3GPP,
		".3gp":  Video3GPP,
		".ts":   VideoMP2T,
		".mp4":  VideoMP4,
		".mpeg": VideoMPEG,
		".mpg":  VideoMPEG,
		".mov":  VideoQuickTime,
		".webm": VideoWEBM,
		".flv":  VideoFLV,
		".m4v":  VideoM4V,
		".mng":  VideoMNG,
		".asx":  VideoMSASF,
		".asf":  VideoMSASF,
		".wmv":  VideoMSWMV,
		".avi":  VideoMSVideo,
	}
)

// GetMimeType gets the MIME type for a given file extension
func GetMimeType(ext string) string {
	if ext == "" {
		return AppOctetStream
	}
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	ext = strings.ToLower(ext)

	if mime, ok := mimeTypes[ext]; ok {
		return mime
	}
	return AppOctetStream
}
func GetMimeTypeFromURL(url string) string {
	ext := filepath.Ext(url)
	return GetMimeType(ext)
}
func GetMimeTypeFromPath(path string) string {
	ext := filepath.Ext(path)
	return GetMimeType(ext)
}
func GetMimeTypeFromFile(file *os.File) string {
	return GetMimeType(filepath.Ext(file.Name()))
}
func GetExtFromMimeType(mimeType string) string {
	for ext, mime := range mimeTypes {

		if mimeType == mime {
			return ext
		}
	}
	return ""
}

func GetExtFromMimeTypeName(mimeType string) string {
	for ext, mime := range mimeTypes {

		if mimeType == mime {
			return strings.TrimLeft(ext, ".")
		}
	}
	return ""
}
