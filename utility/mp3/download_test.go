package mp3

import (
	"path/filepath"
	"strings"
	"testing"
)

const dirPath = "/Users/huoyinghui/Music/网易云音乐/"

func TestStr(t *testing.T) {
	src := "https://music.163.com/#/song?id=1901371647&userid=429237423"
	src = strings.Replace(src, "#/", "", 1)
	t.Log(src)
}
func TestParseSrcLink(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantId  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				src: "https://music.163.com/#/song?id=1901371647&userid=429237423",
			},
			wantId:  "1901371647",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := ParseSrcLink(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSrcLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("ParseSrcLink() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestDownById(t *testing.T) {
	type args struct {
		id   string
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				id:   "1901371647",
				path: filepath.Join(dirPath, "tt.mp3"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownById(tt.args.id, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("DownById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMusicGet(t *testing.T) {
	type args struct {
		src      string
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				src:      "https://music.163.com/#/song?id=1331905859&userid=429237423",
				filename: "Sinus 400Hz &gt; 10dB",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MusicGet(tt.args.src, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("MusicGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
