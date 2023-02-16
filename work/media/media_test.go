package media

import (
	"testing"
)

func TestMedia_Get(t *testing.T) {
	//image := "C:/Users/nevermore/Desktop/yqxc/images/89dc1ef7288000ddcb9407dd220752fc.jpeg"
	//file := "C:/Users/nevermore/Desktop/yqxc/images/笔记本电脑发票.pdf"
	video := "C:/Users/nevermore/Desktop/yqxc/images/视屏/v.f30.mp4"

	media := NewMedia()
	//image, e := media.UploadImage(path)
	//image, e := media.UploadFile(file)
	image, e := media.UploadVideo(video)
	if e != nil {
		t.Errorf("失败")
	}
	localpath, e := media.Get(image.Media_id)
	if e != nil {
		t.Errorf("%v", e)
	}

	t.Log(localpath)

}

func TestMedia_UploadImage(t *testing.T) {
	media := NewMedia()
	image, e := media.UploadImage("C:/Users/nevermore/Desktop/yqxc/images/89dc1ef7288000ddcb9407dd220752fc.jpeg")
	if e != nil {
		t.Errorf("失败")
	}
	t.Log(image)
}
