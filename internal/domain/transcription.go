package domain

type Transcription struct {
	Text string
}

type transcriptionV2 struct {
	operatorText string
	callerText   string
}

type Response struct {
	Status  int
	Message Transcription
}
