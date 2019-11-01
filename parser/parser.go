package parser

import (
	"regexp"
	"strconv"
	"strings"
)

// VideoType is the video codec we need in a given playlist
type VideoType string

// AudioType is the audio codec we need in a given playlist
type AudioType string

// AudioLanguage is the audio language we need in a given playlist
type AudioLanguage string

// CaptionLanguage is the audio language we need in a given playlist
type CaptionLanguage string

const (
	videoHDR10       VideoType = "hdr10"
	videoDolbyVision VideoType = "dovi"
	videoHEVC        VideoType = "hevc"
	videoH264        VideoType = "avc"

	audioAAC                AudioType = "aac"
	audioNoAudioDescription AudioType = "noAd"

	audioLangPTBR AudioLanguage = "pt-br"
	audioLangES   AudioLanguage = "es"
	audioLangEN   AudioLanguage = "en"

	captionPTBR CaptionLanguage = "pt-br"
	captionES   CaptionLanguage = "es"
	captionEN   CaptionLanguage = "en"
)

// MediaFilters is a struct that carry all the information passed via url
type MediaFilters struct {
	Videos           []VideoType       `json:"Videos,omitempty"`
	Audios           []AudioType       `json:"Audios,omitempty"`
	AudioLanguages   []AudioLanguage   `json:"AudioLanguages,omitempty"`
	CaptionLanguages []CaptionLanguage `json:"CaptionLanguages,omitempty"`
	MaxBitrate       int               `json:"MinBitrate,omitempty"`
	MinBitrate       int               `json:"MaxBitrate,omitempty"`
}

// Parse will generate a MediaFilters struct with
// all the filters that needs to be applied to the
// master manifest.
func Parse(filters string) (*MediaFilters, error) {
	mf := new(MediaFilters)
	parts := strings.Split(filters, "/")
	re := regexp.MustCompile(`(.*)\((.*)\)`)

	for _, part := range parts {
		subparts := re.FindStringSubmatch(part)
		// FindStringSubmatch should return a slice with
		// the full string, the key and filters (3 elements)
		if len(subparts) != 3 {
			continue
		}

		filters := strings.Split(subparts[2], ",")

		switch key := subparts[1]; key {
		case "v":
			for _, videoType := range filters {
				mf.Videos = append(mf.Videos, VideoType(videoType))
			}
		case "a":
			for _, audioType := range filters {
				mf.Audios = append(mf.Audios, AudioType(audioType))
			}
		case "al":
			for _, audioLanguage := range filters {
				mf.AudioLanguages = append(mf.AudioLanguages, AudioLanguage(audioLanguage))
			}
		case "c":
			for _, captionLanguage := range filters {
				mf.CaptionLanguages = append(mf.CaptionLanguages, CaptionLanguage(captionLanguage))
			}
		case "b":
			mf.MinBitrate, _ = strconv.Atoi(filters[0])
			mf.MaxBitrate, _ = strconv.Atoi(filters[1])
		}
	}

	return mf, nil
}