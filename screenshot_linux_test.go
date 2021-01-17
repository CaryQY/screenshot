package screenshot

import (
	"image/png"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestShotScreen_CaptureScreen(t *testing.T) {
	shot, err := New()
	if err != nil {
		t.Error(err)
	}
	defer shot.Close()
	img, err := shot.CaptureScreen()
	if err != nil {
		t.Error(err)
	}
	f, err := os.Create("./" + time.Now().Format("200601021504") + ".png")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		t.Error(err)
	}

	t.Log("succeed")
}

func TestShotScreen_GetScreen(t *testing.T) {
	shot,err := New()
	if err != nil {
		t.Error(err)
	}
	defer shot.Close()
	img,err := shot.GetScreen(75)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Create("./" + time.Now().Format("200601021504") + ".jpg")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	_,err = f.Write(img)
	if err != nil {
		t.Error(err)
	}
	t.Log("succeed")
}

func TestLoopCreate1000Images(t *testing.T)  {
	shot,err := New()
	if err != nil {
		t.Error(err)
	}
	defer shot.Close()

	// loop create
	for i:=0;i<1000;i++ {
		img,err := shot.GetScreen(75)
		if err != nil {
			t.Error(err)
		}
		f, err := os.Create("./" + time.Now().Format("200601021504") + "_" + randStr() + ".jpg")
		if err != nil {
			t.Error(err)
		}
		_,err = f.Write(img)
		if err != nil {
			t.Error(err)
		}
		_ = f.Close()
	}

	t.Log("succeed")
}

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randStr() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
