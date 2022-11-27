package mp3

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Play(ctx context.Context, filepath string) (err error) {
	// 1. 打开mp3文件
	audioFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	// 使用defer防止文件描述服忘记关闭导致资源泄露
	defer audioFile.Close()

	// 对文件进行解码
	audioStreamer, format, err := mp3.Decode(audioFile)
	if err != nil {
		return err
	}

	defer audioStreamer.Close()
	// SampleRate is the number of samples per second. 采样率
	// 通过采样率来更改播放速度, 只能通过整形倍数变换粒度太粗
	sr := format.SampleRate * 1
	_ = speaker.Init(sr, sr.N(time.Second/10))

	// 增加控制节点
	/* 默认是播放 */
	/* count为-1是循环播放 */
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, audioStreamer), Paused: false}

	// 增加音量控制
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0, // 为0是当前系统默认音量
		Silent:   false,
	}

	speed := beep.ResampleRatio(6, 1, volume)

	//Ctrl allows for pausing a Streamer.
	speaker.Play(speed)

	// 增加控制信息
	for {
		// 通过Ctrl控制
		fmt.Println("Press [ENTER] to pause/resume. ")
		_, _ = fmt.Scanln()

		speaker.Lock()
		ctrl.Paused = !ctrl.Paused
		// 增加音量
		volume.Volume += 0.5
		//SetRatio sets the resampling ratio. This does not cause any glitches in the stream.
		// 调整播放速度
		speed.SetRatio(speed.Ratio() + 0.1)
		speaker.Unlock()
	}
}
