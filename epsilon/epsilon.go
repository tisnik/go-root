// Backend for the Epsilon project

/*
Apache NON-AI License, Version 2.0

TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

1. Definitions.

“License” shall mean the terms and conditions for use, reproduction, and
distribution as defined by Sections 1 through 9 of this document.

“Licensor” shall mean the copyright owner or entity authorized by the copyright
owner that is granting the License.

“Legal Entity” shall mean the union of the acting entity and all other entities
that control, are controlled by, or are under common control with that entity.
For the purposes of this definition, “control” means (i) the power, direct or
indirect, to cause the direction or management of such entity, whether by
contract or otherwise, or (ii) ownership of fifty percent (50%) or more of the
outstanding shares, or (iii) beneficial ownership of such entity.

“You” (or “Your”) shall mean an individual or Legal Entity exercising
permissions granted by this License.

“Source” form shall mean the preferred form for making modifications, including
but not limited to software source code, documentation source, and
configuration files.

“Object” form shall mean any form resulting from mechanical transformation or
translation of a Source form, including but not limited to compiled object
code, generated documentation, and conversions to other media types.

“Work” shall mean the work of authorship, whether in Source or Object form,
made available under the License, as indicated by a copyright notice that is
included in or attached to the work (an example is provided in the Appendix
below).

“Derivative Works” shall mean any work, whether in Source or Object form, that
is based on (or derived from) the Work and for which the editorial revisions,
annotations, elaborations, or other modifications represent, as a whole, an
original work of authorship.  For the purposes of this License, Derivative
Works shall not include works that remain separable from, or merely link (or
bind by name) to the interfaces of, the Work and Derivative Works thereof.

“Contribution” shall mean any work of authorship, including the original
version of the Work and any modifications or additions to that Work or
Derivative Works thereof, that is intentionally submitted to Licensor for
inclusion in the Work by the copyright owner or by an individual or Legal
Entity authorized to submit on behalf of the copyright owner. For the purposes
of this definition, “submitted” means any form of electronic, verbal, or
written communication sent to the Licensor or its representatives, including
but not limited to communication on electronic mailing lists, source code
control systems, and issue tracking systems that are managed by, or on behalf
of, the Licensor for the purpose of discussing and improving the Work, but
excluding communication that is conspicuously marked or otherwise designated in
writing by the copyright owner as “Not a Contribution.”

“Contributor” shall mean Licensor and any individual or Legal Entity on behalf
of whom a Contribution has been received by Licensor and subsequently
incorporated within the Work.

2. Grant of Copyright License.

Subject to the terms and conditions of this License, each Contributor hereby
grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free,
irrevocable copyright license to reproduce, prepare Derivative Works of,
publicly display, publicly perform, sublicense, and distribute the Work and
such Derivative Works in Source or Object form, under the following conditions:

  2.1. You shall not use the Covered Software in the creation of an Artificial
  Intelligence training dataset, including but not limited to any use that
  contributes to the training or development of an AI model or algorithm,
  unless You obtain explicit written permission from the Contributor to do so.

  2.2. You acknowledge that the Covered Software is not intended for use in the
  creation of an Artificial Intelligence training dataset, and that the
  Contributor has no obligation to provide support or assistance for any use
  that violates this license.

3. Grant of Patent License.

Subject to the terms and conditions of this License, each Contributor hereby
grants to You a perpetual, worldwide, non-exclusive, no-charge, royalty-free,
irrevocable (except as stated in this section) patent license to make, have
made, use, offer to sell, sell, import, and otherwise transfer the Work, where
such license applies only to those patent claims licensable by such Contributor
that are necessarily infringed by their Contribution(s) alone or by combination
of their Contribution(s) with the Work to which such Contribution(s) was
submitted. If You institute patent litigation against any entity (including a
cross-claim or counterclaim in a lawsuit) alleging that the Work or a
Contribution incorporated within the Work constitutes direct or contributory
patent infringement, then any patent licenses granted to You under this License
for that Work shall terminate as of the date such litigation is filed.

4. Redistribution.

You may reproduce and distribute copies of the Work or Derivative Works thereof
in any medium, with or without modifications, and in Source or Object form,
provided that You meet the following conditions:

  1. You must give any other recipients of the Work or Derivative Works a copy
  of this License; and

  2. You must cause any modified files to carry prominent notices stating that
  You changed the files; and

  3. You must retain, in the Source form of any Derivative Works that You
  distribute, all copyright, patent, trademark, and attribution notices from
  the Source form of the Work, excluding those notices that do not pertain to
  any part of the Derivative Works; and

  4. If the Work includes a “NOTICE” text file as part of its distribution,
  then any Derivative Works that You distribute must include a readable copy of
  the attribution notices contained within such NOTICE file, excluding those
  notices that do not pertain to any part of the Derivative Works, in at least
  one of the following places: within a NOTICE text file distributed as part of
  the Derivative Works; within the Source form or documentation, if provided
  along with the Derivative Works; or, within a display generated by the
  Derivative Works, if and wherever such third-party notices normally appear.
  The contents of the NOTICE file are for informational purposes only and do
  not modify the License.

You may add Your own attribution notices within Derivative Works that You
distribute, alongside or as an addendum to the NOTICE text from the Work,
provided that such additional attribution notices cannot be construed as
modifying the License.  You may add Your own copyright statement to Your
modifications and may provide additional or different license terms and
conditions for use, reproduction, or distribution of Your modifications, or for
any such Derivative Works as a whole, provided Your use, reproduction, and
distribution of the Work otherwise complies with the conditions stated in this
License.

5. Submission of Contributions.

Unless You explicitly state otherwise, any Contribution intentionally submitted
for inclusion in the Work by You to the Licensor shall be under the terms and
conditions of this License, without any additional terms or conditions.
Notwithstanding the above, nothing herein shall supersede or modify the terms
of any separate license agreement you may have executed with Licensor regarding
such Contributions.

6. Trademarks.

This License does not grant permission to use the trade names, trademarks,
service marks, or product names of the Licensor, except as required for
reasonable and customary use in describing the origin of the Work and
reproducing the content of the NOTICE file.

7. Disclaimer of Warranty.

Unless required by applicable law or agreed to in writing, Licensor provides
the Work (and each Contributor provides its Contributions) on an “AS IS” BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied,
including, without limitation, any warranties or conditions of TITLE,
NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.  You
are solely responsible for determining the appropriateness of using or
redistributing the Work and assume any risks associated with Your exercise of
permissions under this License.

8. Limitation of Liability.

In no event and under no legal theory, whether in tort (including negligence),
contract, or otherwise, unless required by applicable law (such as deliberate
and grossly negligent acts) or agreed to in writing, shall any Contributor be
liable to You for damages, including any direct, indirect, special, incidental,
or consequential damages of any character arising as a result of this License
or out of the use or inability to use the Work (including but not limited to
damages for loss of goodwill, work stoppage, computer failure or malfunction,
or any and all other commercial damages or losses), even if such Contributor
has been advised of the possibility of such damages.

9. Accepting Warranty or Additional Liability.

While redistributing the Work or Derivative Works thereof, You may choose to
offer, and charge a fee for, acceptance of support, warranty, indemnity, or
other liability obligations and/or rights consistent with this License.
However, in accepting such obligations, You may act only on Your own behalf and
on Your sole responsibility, not on behalf of any other Contributor, and only
if You agree to indemnify, defend, and hold each Contributor harmless for any
liability incurred by, or claims asserted against, such Contributor by reason
of your accepting any such warranty or additional liability.

END OF TERMS AND CONDITIONS
*/

package main

import (
	"embed"
	_ "embed"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// default server port
const Port = 8080

const HttpRequestTimeout = 5 * time.Second
const MaxRows = 99
const MaxColumns = 'Z'

// ---------------------------------------------------------------------------
// New data types
// ---------------------------------------------------------------------------

type Storage interface {
}

type MemoryStorage struct {
}

type Configuration struct {
	port             uint `json:"port"`
	showHelp         bool
	showVersion      bool
	generateExamples bool
}

// application holds the dependencies for our web application.
type Application struct {
	config Configuration
	logger *log.Logger
}

// Server interface
type Server interface {
	Serve()
}

// ServerImpl is a simple HTTP server implementation
type ServerImpl struct {
	app *Application
}

type Coordinate struct {
	Column string
	Row    int
}

type Cell string

type WorkSheet struct {
	Name  string
	Cells map[Coordinate]Cell
}

// ---------------------------------------------------------------------------
// Resources embedded into the final binary file
// ---------------------------------------------------------------------------

// HTML pages, templates, and part of pages
//
//go:embed html/page_header.htm
var PageHeader string

//
//go:embed html/page_footer.htm
var PageFooter string

// Static images in PNG format
//
//go:embed images/*.png
var StaticImages embed.FS

// Static icons in SVG format
//
//go:embed icons/*.svg
var StaticIcons embed.FS

//go:emded icons/favicon.ico
var Favicon []byte

// Stylesheets
//
//go:embed css/*.css
var StaticStylesheets embed.FS

// Scripts
//
//go:embed scripts/*.js
var Scripts embed.FS

// ---------------------------------------------------------------------------
// Interface implementations
// ---------------------------------------------------------------------------

// NewStorage creates and returns an in-memory storage implementation.
func NewStorage() Storage {
	return MemoryStorage{}
}

// ---------------------------------------------------------------------------
// HTTP server implementations
// ---------------------------------------------------------------------------

// NewServer creates a server for the provided application.
func NewServer(app *Application) Server {
	return ServerImpl{
		app: app,
	}
}

// serveStaticImage serves a PNG image from the embedded static image bundle.
//
// It maps requests under /image/ to files in the img/ directory, reads the
// requested file from StaticImages, and writes it to the response with a
// Content-Type of image/png.
//
// If the image cannot be found or read, it logs the error and returns a 404
// Not Found.
func (s ServerImpl) serveStaticImage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()

	// construct proper image name from provided path
	imageName := "img/" + strings.TrimPrefix(path, "/image/")

	// read binary data bundled together with the application
	binaryData, err := StaticImages.ReadFile(imageName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(binaryData)
}

// serveStaticIcon serves a SVG icon from the embedded static icon bundle.
//
// It maps requests under /icons/ to files in the icons/ directory, reads the
// requested file from StaticIcons, and writes it to the response with a
// Content-Type of image/svg+xml.
//
// If the icon cannot be found or read, it logs the error and returns a 404
// Not Found.
func (s ServerImpl) serveStaticIcon(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()

	// construct proper image name from provided path
	imageName := "icons/" + strings.TrimPrefix(path, "/icons/")

	// read text data bundled together with the application
	textData, err := StaticIcons.ReadFile(imageName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(textData)
}

// serveStaticStylesheet serves a CSS stylesheet from the embedded static
// assets.
//
// It maps the request URL path to a file under the embedded css/
// directory, reads the file contents, and returns them with a text/css
// content type.
//
// If the stylesheet cannot be found, it logs the error and
// responds with 404 Not Found.
func (s ServerImpl) serveStaticStylesheet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	styleSheetName := "css/" + strings.TrimPrefix(path, "/css/")
	data, err := StaticStylesheets.ReadFile(styleSheetName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/css")
	w.Write(data)
}

// serveFavicon serves the application's favicon as an ICO image.
//
// It sets the response content type to image/x-icon and writes the embedded
// icon bytes.
func (s ServerImpl) serveFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.Write(Favicon)
}

// serveScripts serves a JavaScript file from the embedded scripts directory.
//
// It maps the request path to an embedded file, returns the file contents,
// and responds with 404 if the script does not exist.
func (s ServerImpl) serveScripts(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	scriptFileName := "scripts/" + strings.TrimPrefix(path, "/scripts/")
	data, err := Scripts.ReadFile(scriptFileName)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/javascript")
	w.Write(data)
}

func (s ServerImpl) renderTable(writer http.ResponseWriter) {
	io.WriteString(writer, "            <tr><th>&nbsp;</th>")
	for column := 'A'; column <= MaxColumns; column++ {
		io.WriteString(writer, fmt.Sprintf("<th>%c</th>", column))

	}
	io.WriteString(writer, "</tr>\n")
	for row := 1; row <= MaxRows; row++ {
		io.WriteString(writer, fmt.Sprintf("            <tr><th>%d</th>", row))
		for column := 'A'; column <= MaxColumns; column++ {
			cellID := fmt.Sprintf("%c%d", column, row)
			io.WriteString(writer, fmt.Sprintf("<td id='%s'></td>", cellID))
		}
		io.WriteString(writer, "</tr>\n")
	}
}

// mainEndpoint serves the main HTML page at the root path.
//
// It returns 404 for any non-root request, writes the page header and footer,
// and renders the main table content in between.
func (s ServerImpl) mainEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(writer, PageHeader)
	s.renderTable(writer)
	io.WriteString(writer, PageFooter)
}

func (s ServerImpl) returnCell(writer http.ResponseWriter, r *http.Request) {
	ID := r.PathValue("id")
	log.Printf("Cell ID provided: %s", ID)

	/*
		ID, err := strconv.Atoi(IDs)
		if err != nil {
			writer.Header().Set("Content-Type", "text/plain")
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

			user, found := s.storage.ReadUser(ID)
			if !found {
				writer.Header().Set("Content-Type", "text/plain")
				writer.WriteHeader(http.StatusNotFound)
				return
			}*/
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(ID)
}

func (s ServerImpl) Serve() {
	port := s.app.config.port
	log.Printf("Starting server on port %d", port)

	http.HandleFunc("/", s.mainEndpoint)

	// static content
	http.HandleFunc("/favicon.ico", s.serveFavicon)

	// scripts
	http.HandleFunc("/scripts/{path}", s.serveScripts)

	// images
	http.HandleFunc("/image/{path}", s.serveStaticImage)

	// icons
	http.HandleFunc("/icons/{path}", s.serveStaticIcon)

	// stylesheets
	http.HandleFunc("/css/{path}", s.serveStaticStylesheet)

	// REST API endpoints
	http.HandleFunc("GET /cell/{id}", s.returnCell)

	// start the server
	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           http.DefaultServeMux,
		ReadHeaderTimeout: HttpRequestTimeout,
		ReadTimeout:       HttpRequestTimeout,
		WriteTimeout:      HttpRequestTimeout,
		IdleTimeout:       HttpRequestTimeout,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------------------------------------------------
// Import/export
// ---------------------------------------------------------------------------

func readCsv(filename string) ([][]string, error) {
	var cells [][]string
	fin, err := os.Open(filename)
	if err != nil {
		return cells, err
	}
	defer fin.Close()

	reader := csv.NewReader(fin)
	records, err := reader.ReadAll()
	if err != nil {
		return cells, err
	}

	var row []string
	for i, r := range records {
		if i == 0 {
			// skip header
			continue
		}
		for _, c := range r {
			row = append(row, c)
		}

		cells = append(cells, row)
	}

	return cells, nil
}

func readWorkSheet(filename string) (WorkSheet, error) {
	return WorkSheet{}, nil
}

func writeWorksheet(worksheet WorkSheet, filename string) error {
	fout, err := os.Create("test.epsilon")
	if err != nil {
		return err
	}
	defer fout.Close()

	enc := gob.NewEncoder(fout)
	err = enc.Encode(worksheet)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return nil
}

func generateExamples() {
	cells := make(map[Coordinate]Cell)
	cells[Coordinate{"A", 1}] = "foo"
	cells[Coordinate{"Z", 99}] = "bar"

	worksheet := WorkSheet{
		Name:  "test1",
		Cells: cells,
	}
	writeWorksheet(worksheet, "ws.epsilon")
}

// main parses command-line options, configures the application, and starts the HTTP server.
// It prints help or version information when requested and reports invalid port values.
func main() {
	var cfg Configuration
	var test bool = false

	if test {
		cells, err := readCsv("sheets/test.csv")
		if err != nil {
			panic(err)
		}
		fmt.Println(len(cells))
		fmt.Println(cells)
		return
	}

	// Read env vars (if any)
	defaultPort := os.Getenv("PORT")
	if defaultPort == "" {
		defaultPort = strconv.Itoa(Port)
	}
	defaultPortNumber, err := strconv.Atoi(defaultPort)
	if err != nil {
		fmt.Printf("Improper port value: %s\n", defaultPort)
		return
	}

	flag.BoolVar(&cfg.showHelp, "h", false, "display help")
	flag.BoolVar(&cfg.showVersion, "v", false, "display version")
	flag.UintVar(&cfg.port, "p", uint(defaultPortNumber), "port for the server (shorthand)")
	flag.UintVar(&cfg.port, "port", uint(defaultPortNumber), "port for the server")
	flag.BoolVar(&cfg.generateExamples, "g", false, "generate example sheets")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		config: cfg,
		logger: logger,
	}

	switch {
	case cfg.showHelp:
		flag.PrintDefaults()
	case cfg.showVersion:
		fmt.Println("version")
	case cfg.generateExamples:
		generateExamples()
	default:
		// default operation: start the HTTP server
		//storage := NewStorage()
		server := NewServer(app)
		server.Serve()
	}
}
