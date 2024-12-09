package interfaces

// Estruturas que mapeiam o JSON
type SDKParams struct {
	VRTBitrateTimestamp int64  `json:"v_rtbitrate_timestamp"`
	StreamSuffix        string `json:"stream_suffix"`
	VBitrate            int    `json:"vbitrate"`
	VCodec              string `json:"VCodec"`
	Resolution          string `json:"resolution"`
	Gop                 int    `json:"gop"`
	VRTBitrate          int    `json:"v_rtbitrate"`
	VRTPsnr             int    `json:"v_rtpsnr"`
}

type MainData struct {
	Flv      string `json:"flv"`
	Hls      string `json:"hls"`
	Cmaf     string `json:"cmaf"`
	Dash     string `json:"dash"`
	Lls      string `json:"lls"`
	Tsl      string `json:"tsl"`
	Tile     string `json:"tile"`
	SDKParams string `json:"sdk_params"` // String a ser parseada separadamente
}

type Origin struct {
	Main MainData `json:"main"`
}

type Data struct {
	Origin Origin `json:"origin"`
}

type Common struct {
	SessionID       string `json:"session_id"`
	RuleIDs         string `json:"rule_ids"`
	PeerAnchorLevel int    `json:"peer_anchor_level"`
}

type Root struct {
	Common Common `json:"common"`
	Data   Data   `json:"data"`
}