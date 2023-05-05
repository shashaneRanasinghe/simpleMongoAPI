package consts

//Handler Errors
const (
	JSONMarshalError      = "Json Marshalling Error "
	GetStudentsError      = "Error Getting Students "
	ResponseWriteError    = "Error Writing Response "
	RequestBodyReadError  = "Error Reading The Request Body"
	RequestBodyCloseError = "Error Closing The Request Body"
	IDError               = "Error Getting The ID"
	ObjectIdConvertError  = "Error converting to ObjectID"
)

//DB ERRORS
const (
	QueryPrepareError  = "Error Preparing Query "
	DBResultsError     = "Error Getting Results From The DB "
	DBRowCloseError    = "Error Closing DB Rows "
	DBScanRowError     = "Error Scanning DB Rows"
	DBRowsError        = "Error In DB Rows"
	DBResultIDError    = "Error Getting Insert ID "
	DBCursorCloseError = "Error Closing Cursor "
)

const (
	StudentNotFound    = "student Not Found"
	StudentDeleteError = "Error Deleting Student"
)
