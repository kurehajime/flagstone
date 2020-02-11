package flagstone

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/icza/gowut/gwu"
)

var port int = 8081
var subURL string = ""
var nonFlagArgs []string = []string{}
var sortArray []string = []string{}
var silent bool = false
var useNonFlagArgs bool = false
var css = `
body {
	background: linear-gradient(-45deg, #fbf3dc, #d8f6fa, #d3dbf6, #d3f6ee);
	background-size: 400% 400%;
	animation: gradient 12s ease infinite;
	font-size: 18px; font-style: normal; 
}
.name,.title,.desc,.description { 
	font-family: "Arial Black", "Arial Bold", Gadget, sans-serif; 
	font-variant: normal; 
	font-weight: 400; 
	line-height: 30px; 
	display:block;
	floet:left;
} 
.name{
	font-size: 22px; font-style: normal; 
	width:150px;
	text-align:right;
	margin-right: 5px;
	text-decoration: underline; 
}
.title{
	font-size: 64px; font-style: bold; 
	line-height: 80px; 
}
.input{
	font-size: 150%;
	min-width:200px;
}
.desc{
	width:350px;
	word-break: break-all;
	margin-left: 5px;
}
.description{
	width: 70%;
}
button {
	display: inline-block;
	width: 70%;
	text-align: center;
	background-color: #5373dc;
	font-family: "Arial Black", "Arial Bold", Gadget, sans-serif; 
	font-size: 24px;
	color: #FFF;
	text-decoration: none;
	font-weight: bold;
	padding: 10px 24px;
	margin-top: 30px;
	border-radius: 4px;
	border-bottom: 2px solid #2649bc;
	margin-bottom: 10px;
}
button:hover {
	transform: translateY(2px);
}

@keyframes gradient {
	0% {
		background-position: 0% 50%;
	}
	60% {
		background-position: 100% 50%;
	}
	100% {
		background-position: 0% 50%;
	}
}
`

// SetPort : Set port number
func SetPort(n int) {
	port = n
}

// SetSubURL : Set Window Name
func SetSubURL(s string) {
	subURL = s
}

// SetSort : Sort By Index
func SetSort(sa []string) {
	sortArray = sa
}

// SetSilent : mute log
func SetSilent(b bool) {
	silent = b
}

// SetUseNonFlagArgs : make non-flag arguments
func SetUseNonFlagArgs(b bool) {
	useNonFlagArgs = b
}

// SetCSS : Specify CSS
func SetCSS(s string) {
	css = s
}

// Launch : Capture the flag
func Launch(title string, message string) ([]string, bool) {
	done := make(chan bool)
	go lanchServer(done, title, message)
	return nonFlagArgs, <-done
}

func indexOf(arr []string, str string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}
	return -1
}

// lanchServer : lanch server
func lanchServer(done chan bool, exeName string, message string) {
	var flags []*flag.Flag
	tbArgs := gwu.NewTextBox("")
	if len(flag.Args()) > 0 {
		tbArgs.SetText(strings.Join(flag.Args(), "\n"))
	}

	// randomize window name
	if subURL == "" {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
		b := []byte(n.String())
		sha512 := sha512.Sum512(b)
		subURL = hex.EncodeToString(sha512[:])
	}

	// get flags
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})

	// sort
	if len(sortArray) > 0 {
		sort.Slice(flags, func(i, j int) bool {
			a := indexOf(sortArray, flags[i].Name)
			b := indexOf(sortArray, flags[j].Name)
			return a < b
		})
	}

	// run server
	server := gwu.NewServer("x", "localhost:"+strconv.Itoa(port))
	server.AddWin(createPage(done, flags, exeName, message))
	if silent {
		w := log.Writer()
		log.SetOutput(ioutil.Discard)
		server.Start(subURL)
		log.SetOutput(w)
	} else {
		server.Start(subURL)
	}
}

// create page
func createPage(done chan bool, flags []*flag.Flag, exeName string, message string) gwu.Window {
	var submit bool = false
	tbArgs := gwu.NewTextBox("")

	// create  window
	if exeName == "" {
		if exe, err := os.Executable(); err == nil {
			exeName = filepath.Base(exe)
		}
	}

	win := gwu.NewWindow(subURL, exeName)
	win.Style().SetFullWidth()
	win.SetHAlign(gwu.HACenter)
	win.SetCellPadding(2)

	title := gwu.NewLabel(exeName)
	title.Style().AddClass("title")
	win.Add(title)

	description := gwu.NewLabel(message)
	description.Style().AddClass("description")
	win.Add(description)

	ctrls := []gwu.HasText{}

	// if button clicked then button disable
	btn := gwu.NewButton("Run")
	btn.AddEHandlerFunc(func(e gwu.Event) {
		var hasErr bool = false
		for i := 0; i < len(ctrls); i++ {
			c := ctrls[i].(gwu.Comp)
			if err := flags[i].Value.Set(ctrls[i].Text()); err != nil {
				hasErr = true
				if c != nil {
					c.Style().SetColor(gwu.ClrRed)
					e.MarkDirty(c)
				}
			} else {
				if c != nil {
					c.Style().SetColor(gwu.ClrBlack)
					e.MarkDirty(c)
				}
			}
		}

		nonFlagArgs = strings.Split(tbArgs.Text(), "\n")

		if hasErr == false {
			submit = true
			btn.SetEnabled(false)
			e.MarkDirty(btn)
		}
	}, gwu.ETypeClick)
	win.Add(btn)

	// args textbox
	if useNonFlagArgs {
		p := gwu.NewHorizontalPanel()
		lbl := gwu.NewLabel("args")
		lbl.Style().AddClass("name")
		p.Add(lbl)
		tbArgs.SetRows(2)
		tbArgs.Style().AddClass("input")
		p.Add(tbArgs)
		desc := gwu.NewLabel(" ")
		desc.Style().AddClass("desc")
		p.Add(desc)
		win.Add(p)
	}

	// add flags textbox
	for i := 0; i < len(flags); i++ {
		p := gwu.NewHorizontalPanel()
		lbl := gwu.NewLabel(flags[i].Name)
		lbl.Style().AddClass("name")
		p.Add(lbl)
		tb := gwu.NewTextBox(flags[i].DefValue)
		tb.Style().AddClass("input")
		ctrls = append(ctrls, tb)
		p.Add(tb)
		desc := gwu.NewLabel(" " + flags[i].Usage)
		desc.Style().AddClass("desc")
		p.Add(desc)
		win.Add(p)
	}

	// if button is disabled then close
	html := `
	<script defer>
		var btn = document.getElementsByTagName("button");
		window.setInterval(function(){
			if(btn[0].disabled == true){
				window.close();
			}
		},100)
	</script>`
	h := gwu.NewHTML(html)
	win.Add(h)

	// if window is closed then action
	loaded := false
	lastTime := time.Now()

	win.AddEHandlerFunc(func(e gwu.Event) {
		lastTime = time.Now()
		loaded = true
	}, gwu.ETypeWinLoad)

	t := gwu.NewTimer(250 * time.Millisecond)
	t.SetRepeat(true)
	t.AddEHandlerFunc(func(e gwu.Event) {
		lastTime = time.Now()
	}, gwu.ETypeStateChange)
	win.Add(t)

	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				if loaded && time.Now().Sub(lastTime).Seconds() >= 1 {
					done <- submit
					ticker.Stop()
					break
				}
			}
		}
	}()

	// css
	cssHTML := "<style>" + css + "</style>"
	win.Add(gwu.NewHTML(cssHTML))
	return win
}
