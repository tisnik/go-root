// Seriál "Programovací jazyk Go"
//    https://www.root.cz/serialy/programovaci-jazyk-go/
//
// Padesátá první část
//    Tvorba grafů v jazyce Go: kreslení ve webovém klientu
//    https://www.root.cz/clanky/tvorba-grafu-v-jazyce-go-kresleni-ve-webovem-klientu/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url string, filename string) (int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	downloaded, err := io.Copy(out, resp.Body)
	return downloaded, err
}

func performDownload(url string, filename string) {
	downloaded, err := downloadFile(url, filename)
	if err != nil {
		fmt.Printf("Download from URL %s failed: %v\n", url, err)
		return
	}
	fmt.Printf("Downloaded %d bytes from URL %s into file %s\n", downloaded, url, filename)
}

func createLinkInDirectory(filename string, directory string) {
	target := directory + filename
	err := os.Link(filename, target)
	if err != nil {
		fmt.Printf("Unable to create hard link from %s to %s: %v\n", filename, target, err)
	}
}

func main() {
	performDownload("https://code.jquery.com/jquery-3.4.1.min.js", "jquery.min.js")
	performDownload("https://cdn.plot.ly/plotly-latest.min.js", "plotly-latest.min.js")
	for example := 1; example <= 10; example++ {
		directory := fmt.Sprintf("plotly%02d/", example)
		createLinkInDirectory("jquery.min.js", directory)
		createLinkInDirectory("plotly-latest.min.js", directory)
	}

}
