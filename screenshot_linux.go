package screenshot

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"sync"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type ShotScreen struct {
}

var (
	xgbC   *xgb.Conn
	lk     *sync.Mutex
	buffer = new(bytes.Buffer)
)

func init() {
	lk = &sync.Mutex{}
	_, err := New()
	if err != nil {
		log.Fatalln(err)
	}
}

func New() (*ShotScreen, error) {
	lk.Lock()
	defer lk.Unlock()

	var err error
	//xgbC, err = xgb.NewConnDisplay(":0")
	xgbC, err = xgb.NewConn()
	if err != nil {
		return nil, err
	}

	shot := &ShotScreen{}
	return shot, nil
}

func (shot *ShotScreen) Close() {
	lk.Lock()
	defer lk.Unlock()
	buffer.Reset()
	if xgbC != nil {
		xgbC.Close()
	}
	//xgbC = &xgb.Conn{}
}

// GetScreen 截屏
// quality 0 - 100
func (shot *ShotScreen) GetScreen(quality int) ([]byte, error) {
	if quality <= 0 {
		quality = 25
	} else if quality > 100 {
		quality = 95
	}
	img, err := shot.CaptureScreen()
	if err != nil {
		return nil, err
	}
	buffer.Reset()
	if err := jpeg.Encode(buffer, img, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (shot *ShotScreen) CaptureScreen() (*image.RGBA, error) {
	r, e := shot.screenRect()
	if e != nil {
		return nil, e
	}
	img, err := shot.captureRect(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (shot *ShotScreen) screenRect() (image.Rectangle, error) {
	screen := xproto.Setup(xgbC).DefaultScreen(xgbC)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}

func (shot *ShotScreen) captureRect(rect image.Rectangle) (*image.RGBA, error) {
	screen := xproto.Setup(xgbC).DefaultScreen(xgbC)
	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(xgbC,
		xproto.ImageFormatZPixmap,
		xproto.Drawable(screen.Root),
		int16(rect.Min.X),
		int16(rect.Min.Y),
		uint16(x), uint16(y),
		0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
}

func (shot *ShotScreen) getScreenXY() (int, int) {
	img, _ := shot.screenRect()
	return img.Dx(), img.Dy()
}
