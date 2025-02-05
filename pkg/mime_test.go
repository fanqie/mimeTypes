package mime

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMimeType(t *testing.T) {
	t.Log("Starting MIME type recognition tests...")

	tests := []struct {
		name     string
		ext      string
		expected string
	}{
		// Text file tests
		{"HTML file", ".html", TextHTML},
		{"HTM file", ".htm", TextHTML},
		{"SHTML file", ".shtml", TextHTML},
		{"CSS file", ".css", TextCSS},
		{"XML file", ".xml", TextXML},
		{"Plain text", ".txt", TextPlain},
		{"J2ME file", ".jad", TextVndJ2ME},
		{"WML file", ".wml", TextVndWapWML},
		{"Component file", ".htc", TextComponent},

		// Image file tests
		{"GIF image", ".gif", ImageGIF},
		{"JPEG image", ".jpeg", ImageJPEG},
		{"JPG image", ".jpg", ImageJPEG},
		{"PNG image", ".png", ImagePNG},
		{"SVG image", ".svg", ImageSVG},
		{"SVGZ image", ".svgz", ImageSVG},
		{"TIF image", ".tif", ImageTIFF},
		{"TIFF image", ".tiff", ImageTIFF},
		{"WEBP image", ".webp", ImageWEBP},
		{"ICO image", ".ico", ImageICO},
		{"JNG image", ".jng", ImageJNG},
		{"BMP image", ".bmp", ImageBMP},
		{"AVIF image", ".avif", ImageAVIF},
		{"WBMP image", ".wbmp", ImageVndWapWBMP},

		// Font file tests
		{"WOFF font", ".woff", FontWOFF},
		{"WOFF2 font", ".woff2", FontWOFF2},
		{"EOT font", ".eot", FontEOT},

		// Application file tests
		{"JavaScript file", ".js", AppJS},
		{"JSON file", ".json", AppJSON},
		{"PDF file", ".pdf", AppPDF},
		{"ZIP file", ".zip", AppZIP},
		{"WASM file", ".wasm", AppWASM},
		{"Atom XML", ".atom", AppAtomXML},
		{"RSS XML", ".rss", AppRSSXML},
		{"JAR file", ".jar", AppJavaArchive},
		{"WAR file", ".war", AppJavaArchive},
		{"EAR file", ".ear", AppJavaArchive},
		{"HQX file", ".hqx", AppMacBinHex40},
		{"PostScript", ".ps", AppPostScript},
		{"EPS file", ".eps", AppPostScript},
		{"AI file", ".ai", AppPostScript},
		{"RTF file", ".rtf", AppRTF},
		{"M3U8 file", ".m3u8", AppMPEGURL},
		{"KML file", ".kml", AppKML},
		{"KMZ file", ".kmz", AppKMZ},
		{"WMLC file", ".wmlc", AppWMLC},
		{"7Z archive", ".7z", App7Z},
		{"Cocoa file", ".cco", AppCocoa},
		{"JAR diff", ".jardiff", AppJavaDiff},
		{"JNLP file", ".jnlp", AppJNLP},
		{"Run file", ".run", AppMakeSelf},
		{"Perl script", ".pl", AppPerl},
		{"Perl module", ".pm", AppPerl},
		{"PRC file", ".prc", AppPilot},
		{"PDB file", ".pdb", AppPilot},
		{"RAR archive", ".rar", AppRAR},
		{"RPM package", ".rpm", AppRPM},
		{"SEA archive", ".sea", AppSEA},
		{"SWF file", ".swf", AppShockwaveFlash},
		{"StuffIt archive", ".sit", AppStuffIt},
		{"TCL script", ".tcl", AppTCL},
		{"TK script", ".tk", AppTCL},
		{"DER certificate", ".der", AppX509},
		{"PEM certificate", ".pem", AppX509},
		{"CRT certificate", ".crt", AppX509},
		{"XPI install", ".xpi", AppXPInstall},
		{"XHTML file", ".xhtml", AppXHTML},
		{"XSPF file", ".xspf", AppXSPF},

		// Office document tests
		{"Word doc", ".doc", AppMSWord},
		{"Excel xls", ".xls", AppMSExcel},
		{"PowerPoint ppt", ".ppt", AppMSPPT},
		{"ODG file", ".odg", AppODG},
		{"ODP file", ".odp", AppODP},
		{"ODS file", ".ods", AppODS},
		{"ODT file", ".odt", AppODT},
		{"DOCX file", ".docx", AppDocx},
		{"XLSX file", ".xlsx", AppXlsx},
		{"PPTX file", ".pptx", AppPptx},

		// Binary file tests
		{"Binary file", ".bin", AppOctetStream},
		{"Executable", ".exe", AppOctetStream},
		{"DLL file", ".dll", AppOctetStream},
		{"DEB package", ".deb", AppOctetStream},
		{"DMG image", ".dmg", AppOctetStream},
		{"ISO image", ".iso", AppOctetStream},
		{"IMG file", ".img", AppOctetStream},
		{"MSI installer", ".msi", AppOctetStream},
		{"MSP file", ".msp", AppOctetStream},

		// Audio file tests
		{"MIDI file", ".mid", AudioMIDI},
		{"MIDI file", ".midi", AudioMIDI},
		{"KAR file", ".kar", AudioMIDI},
		{"MP3 audio", ".mp3", AudioMP3},
		{"OGG audio", ".ogg", AudioOGG},
		{"M4A audio", ".m4a", AudioM4A},
		{"RA audio", ".ra", AudioRealAudio},

		// Video file tests
		{"3GPP video", ".3gpp", Video3GPP},
		{"3GP video", ".3gp", Video3GPP},
		{"TS video", ".ts", VideoMP2T},
		{"MP4 video", ".mp4", VideoMP4},
		{"MPEG video", ".mpeg", VideoMPEG},
		{"MPG video", ".mpg", VideoMPEG},
		{"QuickTime video", ".mov", VideoQuickTime},
		{"WebM video", ".webm", VideoWEBM},
		{"FLV video", ".flv", VideoFLV},
		{"M4V video", ".m4v", VideoM4V},
		{"MNG video", ".mng", VideoMNG},
		{"ASX video", ".asx", VideoMSASF},
		{"ASF video", ".asf", VideoMSASF},
		{"WMV video", ".wmv", VideoMSWMV},
		{"AVI video", ".avi", VideoMSVideo},

		// Boundary tests
		{"Unknown extension", ".unknown", AppOctetStream},
		{"Empty extension", "", AppOctetStream},
	}

	totalTests := len(tests)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMimeType(tt.ext)
			
			fmt.Printf("Testing: %s, extension '%s' is %s\n", tt.name, tt.ext, got)
			assert.Equal(t, got, tt.expected)

		})
	}

	fmt.Printf("Test completed: %d total test cases\n", totalTests)

}
