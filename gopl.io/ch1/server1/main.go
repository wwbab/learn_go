// server1 是一个迷你回声器
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)
var mu sync.Mutex
var palette = []color.Color{color.RGBA{0x11,0x22,0xBB,0xff}, color.Black}
const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的第二种颜色
	res = 0.001 // 角度分辨率
	size = 100 // 图像中的画布包含[-size..+size]
	nframes = 64 // 动画中的帧数
	delay = 8 // 以10ms为单位的帧间延迟
)

func main() {
	// 回声请求调用处理程序
	// http.HandleFunc 将特定的 URL 路径与对应的处理函数绑定。
	handler := func (w http.ResponseWriter, r *http.Request)  {
		lissajous(w)
	}
	handler1 := func (w http.ResponseWriter, r *http.Request)  {
		lissajous2(w)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/?cycles=", handler1)
	// log.Fatal 用于输出日志信息并终止程序的执行。
	// http.ListenAndServe 的作用是启动一个 HTTP 服务器，
	// 监听指定的地址和端口，
	// 并处理客户端的 HTTP 请求。
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 处理程序回显请求 URL r 的路径部分
// http.ResponseWriter 用于向客户端发送响应
// *http.Request 则包含了客户端的请求信息。
/* func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
*/

func lissajous(out io.Writer) {
	mu.Lock()
	cycles := 5 // 完整的 x 振荡器变化的个数
	freq := rand.Float64() * 3.0 // y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意忽略编码错误。
	mu.Unlock()
}

func lissajous2(out io.Writer) {
	mu.Lock()
	var r *http.Request
	cycles, _ := strconv.Atoi(r.FormValue("cycles")) // 完整的 x 振荡器变化的个数
	freq := rand.Float64() * 3.0 // y 振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意忽略编码错误
	mu.Unlock()
}

    

    
