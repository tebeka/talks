package nrsc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const (
	port = 9888
)

var root string

func testDir() string {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}
	return fmt.Sprintf("%s/nrsc-test-%s-%s", os.TempDir(), os.Getenv("USER"), host)
}

func TestMask(t *testing.T) {
	resp := getResp(t, "/static/i.gif")
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("got masked resource - %d\n", resp.StatusCode)
	}
}

func TestText(t *testing.T) {
	expected := map[string]string{
		"Content-Size": "12",
		"Content-Type": "text/plain",
	}
	checkPath(t, "ht.txt", expected)
}

func TestSub(t *testing.T) {
	expected := map[string]string{
		"Content-Size": "1150",
		"Content-Type": "image/",
	}
	checkPath(t, "sub/favicon.ico", expected)
}

// / serves a template
func TestTempalte(t *testing.T) {
	server := startServer(t)
	if server == nil {
		t.Fatalf("can't start server")
	}
	defer server.Process.Kill()

	url := fmt.Sprintf("http://localhost:%d", port)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("can't GET / - %s", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("can't read body - %s", err)
	}

	if string(data) != "The number is 7\n" {
		t.Fatalf("bad template reply - %s", string(data))
	}
}

func createMain() error {
	filename := fmt.Sprintf("%s/main.go", root)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, code, port)
	return nil
}

func initDir() error {
	// Ignore error value, since it might not be there
	os.RemoveAll(root)

	err := os.Mkdir(root, 0777)
	if err != nil {
		return err
	}

	return createMain()
}

func get(path string) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:%d/static/%s", port, path)
	return http.Get(url)
}

func startServer(t *testing.T) *exec.Cmd {
	cmd := exec.Command(fmt.Sprintf("%s/nrsc-test", root))
	// Ignore errors, test will fail anyway if server not running
	cmd.Start()

	// Wait for server
	url := fmt.Sprintf("http://localhost:%d", port)
	start := time.Now()
	for time.Since(start) < time.Duration(2*time.Second) {
		_, err := http.Get(url)
		if err == nil {
			return cmd
		}
		time.Sleep(time.Second / 10)
	}

	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	t.Fatalf("can't connect to server")
	return nil
}

func init() {
	root = testDir()

	if err := initDir(); err != nil {
		panic(err)
	}

	cwd, _ := os.Getwd()
	path := func(name string) string {
		return fmt.Sprintf("%s/%s", cwd, name)
	}

	os.Chdir("nrsc")
	cmd := exec.Command("go", "build")
	if err := cmd.Run(); err != nil {
		fmt.Printf("error building nrsc: %s\n", err)
		panic(err)
	}

	os.Chdir(root)
	defer os.Chdir(cwd)

	testExe := "nrsc-test"

	cmd = exec.Command("go", "build", "-o", testExe)
	if err := cmd.Run(); err != nil {
		fmt.Printf("error building: %s\n", err)
		panic(err)
	}

	cmd = exec.Command(path("nrsc/nrsc"), testExe, path("test-resources"))
	if err := cmd.Run(); err != nil {
		fmt.Printf("error packing: %s\n", err)
		panic(err)
	}
}

func checkHeaders(t *testing.T, expected map[string]string, headers http.Header) {
	for key := range expected {
		v1 := expected[key]
		v2 := headers.Get(key)
		if !strings.HasPrefix(v2, v1) {
			t.Fatalf("bad header %s: %s <-> %s", key, v1, v2)
		}
	}

	key := "Last-Modified"
	value := headers.Get(key)
	if value == "" {
		t.Fatalf("no %s header", key)
	}
}

func getResp(t *testing.T, path string) *http.Response {
	server := startServer(t)
	if server == nil {
		return nil
	}
	defer server.Process.Kill()

	resp, err := get(path)
	if err != nil {
		t.Fatalf("%s\n", err)
		return nil
	}

	return resp
}

func checkPath(t *testing.T, path string, expected map[string]string) {
	resp := getResp(t, path)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("bad reply - %s", resp.Status)
	}

	checkHeaders(t, expected, resp.Header)
}

const code = `
package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"bitbucket.org/tebeka/nrsc"
)

type params struct {
	Number  int
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	t, err := nrsc.LoadTemplates(nil, "t.html")
	if err != nil {
		http.NotFound(w, req)
	}
	if err = t.Execute(w, params{7}); err != nil {
		http.NotFound(w, req)
	}
}

func main() {
	nrsc.Handle("/static/")
	nrsc.Mask(regexp.MustCompile(".gif$"))
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":%d", nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %%s\n", err)
		os.Exit(1)
	}
}
`
