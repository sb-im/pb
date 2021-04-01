package signal

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestJSONUnmarshalSessionDescription(t *testing.T) {
	rawSdp := `{"id":"","track_source":0,"sdp":{"type":"offer","sdp":"v=0\r\no=- 3130295548462269111 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\na=group:BUNDLE 0\r\na=extmap-allow-mixed\r\na=msid-semantic: WMS\r\nm=video 9 UDP/TLS/RTP/SAVPF 96 97 98 99 100 101 102 121 127 120 125 107 108 109 124 119 123\r\nc=IN IP4 0.0.0.0\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=ice-ufrag:qgyo\r\na=ice-pwd:DRNdp7V/NeGghkwnjgtegiXP\r\na=ice-options:trickle\r\na=fingerprint:sha-256 D7:C1:6E:12:7A:BF:11:EB:75:8F:03:C3:09:47:88:6A:6A:CF:BE:AB:37:C5:CE:02:59:FB:9A:F2:60:F7:A4:99\r\na=setup:actpass\r\na=mid:0\r\na=extmap:1 urn:ietf:params:rtp-hdrext:toffset\r\na=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\r\na=extmap:3 urn:3gpp:video-orientation\r\na=extmap:4 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=extmap:5 http://www.webrtc.org/experiments/rtp-hdrext/playout-delay\r\na=extmap:6 http://www.webrtc.org/experiments/rtp-hdrext/video-content-type\r\na=extmap:7 http://www.webrtc.org/experiments/rtp-hdrext/video-timing\r\na=extmap:8 http://www.webrtc.org/experiments/rtp-hdrext/color-space\r\na=extmap:9 urn:ietf:params:rtp-hdrext:sdes:mid\r\na=extmap:10 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id\r\na=extmap:11 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id\r\na=sendrecv\r\na=msid:- 2a0548b1-7ec9-4bea-af2d-7a0bf6f1c9ae\r\na=rtcp-mux\r\na=rtcp-rsize\r\na=rtpmap:96 VP8/90000\r\na=rtcp-fb:96 goog-remb\r\na=rtcp-fb:96 transport-cc\r\na=rtcp-fb:96 ccm fir\r\na=rtcp-fb:96 nack\r\na=rtcp-fb:96 nack pli\r\na=rtpmap:97 rtx/90000\r\na=fmtp:97 apt=96\r\na=rtpmap:98 VP9/90000\r\na=rtcp-fb:98 goog-remb\r\na=rtcp-fb:98 transport-cc\r\na=rtcp-fb:98 ccm fir\r\na=rtcp-fb:98 nack\r\na=rtcp-fb:98 nack pli\r\na=fmtp:98 profile-id=0\r\na=rtpmap:99 rtx/90000\r\na=fmtp:99 apt=98\r\na=rtpmap:100 VP9/90000\r\na=rtcp-fb:100 goog-remb\r\na=rtcp-fb:100 transport-cc\r\na=rtcp-fb:100 ccm fir\r\na=rtcp-fb:100 nack\r\na=rtcp-fb:100 nack pli\r\na=fmtp:100 profile-id=2\r\na=rtpmap:101 rtx/90000\r\na=fmtp:101 apt=100\r\na=rtpmap:102 H264/90000\r\na=rtcp-fb:102 goog-remb\r\na=rtcp-fb:102 transport-cc\r\na=rtcp-fb:102 ccm fir\r\na=rtcp-fb:102 nack\r\na=rtcp-fb:102 nack pli\r\na=fmtp:102 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f\r\na=rtpmap:121 rtx/90000\r\na=fmtp:121 apt=102\r\na=rtpmap:127 H264/90000\r\na=rtcp-fb:127 goog-remb\r\na=rtcp-fb:127 transport-cc\r\na=rtcp-fb:127 ccm fir\r\na=rtcp-fb:127 nack\r\na=rtcp-fb:127 nack pli\r\na=fmtp:127 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f\r\na=rtpmap:120 rtx/90000\r\na=fmtp:120 apt=127\r\na=rtpmap:125 H264/90000\r\na=rtcp-fb:125 goog-remb\r\na=rtcp-fb:125 transport-cc\r\na=rtcp-fb:125 ccm fir\r\na=rtcp-fb:125 nack\r\na=rtcp-fb:125 nack pli\r\na=fmtp:125 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f\r\na=rtpmap:107 rtx/90000\r\na=fmtp:107 apt=125\r\na=rtpmap:108 H264/90000\r\na=rtcp-fb:108 goog-remb\r\na=rtcp-fb:108 transport-cc\r\na=rtcp-fb:108 ccm fir\r\na=rtcp-fb:108 nack\r\na=rtcp-fb:108 nack pli\r\na=fmtp:108 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f\r\na=rtpmap:109 rtx/90000\r\na=fmtp:109 apt=108\r\na=rtpmap:124 red/90000\r\na=rtpmap:119 rtx/90000\r\na=fmtp:119 apt=124\r\na=rtpmap:123 ulpfec/90000\r\na=ssrc-group:FID 523105335 1512610506\r\na=ssrc:523105335 cname:iv0H4wwy4aajyVGb\r\na=ssrc:523105335 msid:- 2a0548b1-7ec9-4bea-af2d-7a0bf6f1c9ae\r\na=ssrc:523105335 mslabel:-\r\na=ssrc:523105335 label:2a0548b1-7ec9-4bea-af2d-7a0bf6f1c9ae\r\na=ssrc:1512610506 cname:iv0H4wwy4aajyVGb\r\na=ssrc:1512610506 msid:- 2a0548b1-7ec9-4bea-af2d-7a0bf6f1c9ae\r\na=ssrc:1512610506 mslabel:-\r\na=ssrc:1512610506 label:2a0548b1-7ec9-4bea-af2d-7a0bf6f1c9ae\r\n"}}`
	var sdp SessionDescription
	if err := json.Unmarshal([]byte(rawSdp), &sdp); err != nil {
		t.Fatal(err)
	}

	p, err := proto.Marshal(&sdp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("size fo protobuf encoding: %d\n", binary.Size(p))

	b, err := json.Marshal(&sdp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("size fo json encoding: %d\n", binary.Size(b))
}
