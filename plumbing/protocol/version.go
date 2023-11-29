package protocol

type Version string

const (
	PROTOCOL_V0 Version = "protocol_version_0"
	PROTOCOL_V1         = "protocol_version_1"
	PROTOCOL_V2         = "protocol_version_2"
)

func DetermineProtocolVersionClient() Version {
	version := PROTOCOL_V0

}
