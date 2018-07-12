package record

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

const streamURL = "rtmp://fms-base1.mitene.ad.jp/agqr/aandg22"

// Do function
func Do(title string, minute int) error {
	fmt.Println("Record start:", title)
	err := exec.Command("rtmpdump", "-r", streamURL, "-v", "-q", "-B", fmt.Sprint(minute*60), "-o", title+".flv").Run()
	if err != nil {
		fmt.Println("rtmpdump:", err)
		return err
	}
	defer os.Remove(title + ".flv")

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst).Format("2006-01-02")
	filename := title + now + ".mp3"
	fmt.Println("Encode start:", title)
	err = exec.Command("ffmpeg", "-i", title+".flv", filename).Run()
	if err != nil {
		fmt.Println("ffmpeg:", err)
		return err
	}
	defer os.Remove(filename)

	fmt.Println("Upload start:", title)
	err = upload(filename)
	if err != nil {
		return err
	}
	fmt.Println("Completed:", title)
	return nil
}

func upload(filename string) error {
	config := dropbox.Config{
		Token: os.Getenv("DROPBOX_TOKEN"),
	}
	dbx := files.New(config)
	f, _ := os.Open(filename)
	defer f.Close()
	commitInfo := files.NewCommitInfo("/" + filename)
	_, err := dbx.Upload(commitInfo, f)
	if err != nil {
		fmt.Println("dropbox:", err)
		return err
	}
	return nil
}
