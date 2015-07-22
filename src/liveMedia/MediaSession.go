package liveMedia

import (
	"fmt"
	. "groupsock"
	"strings"
)

type MediaSession struct {
	controlPath  string
	absStartTime string
	absEndTime   string
}

func NewMediaSession(sdpDesc string) *MediaSession {
	mediaSession := new(MediaSession)
	mediaSession.InitWithSDP(sdpDesc)
	return mediaSession
}

type MediaSubSession struct {
	rtpSocket             *GroupSock
	rtcpSocket            *GroupSock
	Sink                  *MediaSink
	rtpSource             *RTPSource
	readSource            IFramedSource
	rtcpInstance          *RTCPInstance
	rtpTimestampFrequency uint
	rtpPayloadFormat      int
	clientPortNum         uint
	protocolName          string
	controlPath           string
	codecName             string
}

func NewMediaSubSession() *MediaSubSession {
	subsession := new(MediaSubSession)
	return subsession
}

func (this *MediaSession) InitWithSDP(sdpDesc string) {
	//subsession := NewMediaSubSession()
}

func (this *MediaSession) ControlPath() string {
	return this.controlPath
}

func (this *MediaSession) AbsStartTime() string {
	return this.absStartTime
}

func (this *MediaSession) AbsEndTime() string {
	return this.absEndTime
}

// MediaSubSession Implementation
func (this *MediaSubSession) Initiate() bool {
	if len(this.codecName) <= 0 {
		fmt.Println("Codec is unspecified")
		return false
	}

	protocolIsRTP := strings.EqualFold(this.protocolName, "RTP")
	if protocolIsRTP {
		this.clientPortNum = this.clientPortNum &^ 1
	}

	this.rtpSocket = NewGroupSock(this.clientPortNum)
	if this.rtpSocket == nil {
		fmt.Println("Failed to create RTP socket")
		return false
	}

	if protocolIsRTP {
		// Set our RTCP port to be the RTP Port +1
		rtcpPortNum := this.clientPortNum | 1
		this.rtcpSocket = NewGroupSock(rtcpPortNum)
	}

	//this.rtcpInstance = NewRTCPInstance()
	return true
}

func (this *MediaSubSession) deInitiate() {
}

func (this *MediaSubSession) ControlPath() string {
	return this.controlPath
}

func (this *MediaSubSession) RtcpInstance() *RTCPInstance {
	return this.rtcpInstance
}

func (this *MediaSubSession) createSourceObject() {
	if strings.EqualFold(this.protocolName, "RTP") {
		this.readSource = NewBasicUDPSource()
		this.rtpSource = nil

		if strings.EqualFold(this.codecName, "MP2T") {
			// this sets "durationInMicroseconds" correctly, based on the PCR values
			//this.readSource = NewMPEG2TransportStreamFramer(this.readSource)
		}
	} else {
		switch this.codecName {
		case "H264":
			//this.readSource = NewH264VideoRTPSource(this.rtpSocket, this.rtpPayloadFormat, this.rtpTimestampFrequency)
		}
	}
}