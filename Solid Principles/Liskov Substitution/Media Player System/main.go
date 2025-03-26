package main

import (
	"fmt"
)

// Media interface defines a method for playing media.
type Media interface {
	Play()
}

// Audio type implements the Media interface.
type Audio struct {
	Title string
}

// Play method for Audio.
func (a Audio) Play() {
	fmt.Printf("Playing audio: %s\n", a.Title)
}

// Video type implements the Media interface.
type Video struct {
	Title string
}

// Play method for Video.
func (v Video) Play() {
	fmt.Printf("Playing video: %s\n", v.Title)
}

// MediaPlayer depends on the Media interface.
type MediaPlayer struct {
	media Media
}

// NewMediaPlayer creates a new MediaPlayer with the given media.
func NewMediaPlayer(m Media) *MediaPlayer {
	return &MediaPlayer{media: m}
}

// PlayMedia calls the Play method on the injected Media type.
func (mp *MediaPlayer) PlayMedia() {
	mp.media.Play()
}

// Main function demonstrates Liskov Substitution Principle in action.
func main() {
	// Create an Audio instance and play it using MediaPlayer.
	audio := Audio{Title: "Imagine - John Lennon"}
	audioPlayer := NewMediaPlayer(audio)
	audioPlayer.PlayMedia()

	// Create a Video instance and play it using MediaPlayer.
	video := Video{Title: "Nature Documentary"}
	videoPlayer := NewMediaPlayer(video)
	videoPlayer.PlayMedia()
}
