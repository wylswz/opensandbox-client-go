package opensandbox

const (
	defaultUserAgent = "opensandbox-client-go/1.0"
)

const (
	headerContentType      = "Content-Type"
	headerAccept           = "Accept"
	headerUserAgent        = "User-Agent"
	headerExecdAccessToken = "X-EXECD-ACCESS-TOKEN"

	contentTypeJSON = "application/json"
	acceptSSE       = "text/event-stream"
)

const (
	execdPathCode        = "/code"
	execdPathCommand     = "/command"
	execdPathFilesUpload = "/files/upload"
)

const (
	defaultFileOwner = "root"
	defaultFileGroup = "root"
)

const (
	errOnEventCallbackRequired = "onEvent callback is required"
)
