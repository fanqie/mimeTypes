package test

import (
	"fmt"
	"github.com/fanqie/mimeTypes/pkg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetMimeType(t *testing.T) {
	t.Log("Starting MIME type recognition tests...")

	tests := []struct {
		name     string
		ext      string
		expected string
	}{
		// Text file tests
		{"HTML file", ".html", mime.TextHTML},
		{"HTM file", ".htm", mime.TextHTML},
		{"SHTML file", ".shtml", mime.TextHTML},
		{"CSS file", ".css", mime.TextCSS},
		{"XML file", ".xml", mime.TextXML},
		{"Plain text", ".txt", mime.TextPlain},
		{"J2ME file", ".jad", mime.TextVndJ2ME},
		{"WML file", ".wml", mime.TextVndWapWML},
		{"Component file", ".htc", mime.TextComponent},

		// Image file tests
		{"GIF image", ".gif", mime.ImageGIF},
		{"JPEG image", ".jpeg", mime.ImageJPEG},
		{"JPG image", ".jpg", mime.ImageJPEG},
		{"PNG image", ".png", mime.ImagePNG},
		{"SVG image", ".svg", mime.ImageSVG},
		{"SVGZ image", ".svgz", mime.ImageSVG},
		{"TIF image", ".tif", mime.ImageTIFF},
		{"TIFF image", ".tiff", mime.ImageTIFF},
		{"WEBP image", ".webp", mime.ImageWEBP},
		{"ICO image", ".ico", mime.ImageICO},
		{"JNG image", ".jng", mime.ImageJNG},
		{"BMP image", ".bmp", mime.ImageBMP},
		{"AVIF image", ".avif", mime.ImageAVIF},
		{"WBMP image", ".wbmp", mime.ImageVndWapWBMP},

		// Font file tests
		{"WOFF font", ".woff", mime.FontWOFF},
		{"WOFF2 font", ".woff2", mime.FontWOFF2},
		{"EOT font", ".eot", mime.FontEOT},

		// Application file tests
		{"JavaScript file", ".js", mime.AppJS},
		{"JSON file", ".json", mime.AppJSON},
		{"PDF file", ".pdf", mime.AppPDF},
		{"ZIP file", ".zip", mime.AppZIP},
		{"WASM file", ".wasm", mime.AppWASM},
		{"Atom XML", ".atom", mime.AppAtomXML},
		{"RSS XML", ".rss", mime.AppRSSXML},
		{"JAR file", ".jar", mime.AppJavaArchive},
		{"WAR file", ".war", mime.AppJavaArchive},
		{"EAR file", ".ear", mime.AppJavaArchive},
		{"HQX file", ".hqx", mime.AppMacBinHex40},
		{"PostScript", ".ps", mime.AppPostScript},
		{"EPS file", ".eps", mime.AppPostScript},
		{"AI file", ".ai", mime.AppPostScript},
		{"RTF file", ".rtf", mime.AppRTF},
		{"M3U8 file", ".m3u8", mime.AppMPEGURL},
		{"KML file", ".kml", mime.AppKML},
		{"KMZ file", ".kmz", mime.AppKMZ},
		{"WMLC file", ".wmlc", mime.AppWMLC},
		{"7Z archive", ".7z", mime.App7Z},
		{"Cocoa file", ".cco", mime.AppCocoa},
		{"JAR diff", ".jardiff", mime.AppJavaDiff},
		{"JNLP file", ".jnlp", mime.AppJNLP},
		{"Run file", ".run", mime.AppMakeSelf},
		{"Perl script", ".pl", mime.AppPerl},
		{"Perl module", ".pm", mime.AppPerl},
		{"PRC file", ".prc", mime.AppPilot},
		{"PDB file", ".pdb", mime.AppPilot},
		{"RAR archive", ".rar", mime.AppRAR},
		{"RPM package", ".rpm", mime.AppRPM},
		{"SEA archive", ".sea", mime.AppSEA},
		{"SWF file", ".swf", mime.AppShockwaveFlash},
		{"StuffIt archive", ".sit", mime.AppStuffIt},
		{"TCL script", ".tcl", mime.AppTCL},
		{"TK script", ".tk", mime.AppTCL},
		{"DER certificate", ".der", mime.AppX509},
		{"PEM certificate", ".pem", mime.AppX509},
		{"CRT certificate", ".crt", mime.AppX509},
		{"XPI install", ".xpi", mime.AppXPInstall},
		{"XHTML file", ".xhtml", mime.AppXHTML},
		{"XSPF file", ".xspf", mime.AppXSPF},

		// Office document tests
		{"Word doc", ".doc", mime.AppMSWord},
		{"Excel xls", ".xls", mime.AppMSExcel},
		{"PowerPoint ppt", ".ppt", mime.AppMSPPT},
		{"ODG file", ".odg", mime.AppODG},
		{"ODP file", ".odp", mime.AppODP},
		{"ODS file", ".ods", mime.AppODS},
		{"ODT file", ".odt", mime.AppODT},
		{"DOCX file", ".docx", mime.AppDocx},
		{"XLSX file", ".xlsx", mime.AppXlsx},
		{"PPTX file", ".pptx", mime.AppPptx},

		// Binary file tests
		{"Binary file", ".bin", mime.AppOctetStream},
		{"Executable", ".exe", mime.AppOctetStream},
		{"DLL file", ".dll", mime.AppOctetStream},
		{"DEB package", ".deb", mime.AppOctetStream},
		{"DMG image", ".dmg", mime.AppOctetStream},
		{"ISO image", ".iso", mime.AppOctetStream},
		{"IMG file", ".img", mime.AppOctetStream},
		{"MSI installer", ".msi", mime.AppOctetStream},
		{"MSP file", ".msp", mime.AppOctetStream},

		// Audio file tests
		{"MIDI file", ".mid", mime.AudioMIDI},
		{"MIDI file", ".midi", mime.AudioMIDI},
		{"KAR file", ".kar", mime.AudioMIDI},
		{"MP3 audio", ".mp3", mime.AudioMP3},
		{"OGG audio", ".ogg", mime.AudioOGG},
		{"M4A audio", ".m4a", mime.AudioM4A},
		{"RA audio", ".ra", mime.AudioRealAudio},

		// Video file tests
		{"3GPP video", ".3gpp", mime.Video3GPP},
		{"3GP video", ".3gp", mime.Video3GPP},
		{"TS video", ".ts", mime.VideoMP2T},
		{"MP4 video", ".mp4", mime.VideoMP4},
		{"MPEG video", ".mpeg", mime.VideoMPEG},
		{"MPG video", ".mpg", mime.VideoMPEG},
		{"QuickTime video", ".mov", mime.VideoQuickTime},
		{"WebM video", ".webm", mime.VideoWEBM},
		{"FLV video", ".flv", mime.VideoFLV},
		{"M4V video", ".m4v", mime.VideoM4V},
		{"MNG video", ".mng", mime.VideoMNG},
		{"ASX video", ".asx", mime.VideoMSASF},
		{"ASF video", ".asf", mime.VideoMSASF},
		{"WMV video", ".wmv", mime.VideoMSWMV},
		{"AVI video", ".avi", mime.VideoMSVideo},

		// Boundary tests
		{"Unknown extension", ".unknown", mime.AppOctetStream},
		{"Empty extension", "", mime.AppOctetStream},
	}

	totalTests := len(tests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mime.GetMimeType(tt.ext)

			assert.Equal(t, got, tt.expected)

		})
	}

	fmt.Printf("Test completed: %d total test cases\n", totalTests)

}

func TestGetMimeTypeFromURL(t *testing.T) {
	t.Run("TestGetMimeTypeFromURL", func(t *testing.T) {
		assert.Equal(t, mime.GetMimeTypeFromURL("https://github.com/fanqie/mimeTypes/demo.gif"), mime.ImageGIF)

	})
}

func TestGetMimeTypeFromPath(t *testing.T) {
	t.Run("TestGetMimeTypeFromPath", func(t *testing.T) {
		assert.Equal(t, mime.GetMimeTypeFromPath("path/demo.gif"), mime.ImageGIF)

	})
}

func TestGetMimeTypeFromFile(t *testing.T) {
	t.Run("TestGetMimeTypeFromFile", func(t *testing.T) {
		file, err := os.Open("./types.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				t.Errorf("error:%s", err)
			}
		}(file)

		assert.Equal(t, mime.GetMimeTypeFromFile(file), mime.AppJSON)

	})
}

func TestGetExtFromMimeType(t *testing.T) {
	t.Run("TestGetExtFromMimeType", func(t *testing.T) {
		assert.Equal(t, mime.GetExtFromMimeType(mime.AppDocx), ".docx")
	})
	t.Run("TestGetExtFromMimeTypeName", func(t *testing.T) {
		assert.Equal(t, mime.GetExtFromMimeTypeName(mime.AppJSON), "json")
	})
}
