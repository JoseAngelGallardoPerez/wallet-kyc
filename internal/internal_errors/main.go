package internal_errors

const (
	InternalServerError = "INTERNAL_SERVER_ERROR"

	RequestIsApproved     = "REQUEST_IS_APPROVED"
	RequestIsPending      = "REQUEST_IS_PENDING"
	RequestIsNotAvailable = "REQUEST_IS_NOT_AVAILABLE"
	RequestAlreadyExists  = "REQUEST_ALREADY_EXISTS"
	RequestNotFound       = "REQUEST_NOT_FOUND"
	RequestNotPending     = "REQUEST_NOT_PENDING"
	RequestBeenApproved   = "REQUEST_BEEN_APPROVED"

	TierNotFound       = "TIER_NOT_FOUND"
	TierIsNotAvailable = "TIER_IS_NOT_AVAILABLE"

	RequirementNotFound       = "REQUIREMENT_NOT_FOUND"
	RequirementIsApproved     = "REQUIREMENT_IS_APPROVED"
	RequirementIsPending      = "REQUIREMENT_IS_PENDING"
	RequirementIsNotAvailable = "REQUIREMENT_IS_NOT_AVAILABLE"
	RequirementsFilled        = "REQUIREMENTS_FILLED"
	RequirementsNotApproved   = "REQUIREMENTS_NOT_APPROVED"
	RequirementsIsPending     = "REQUIREMENTS_IS_PENDING"

	UserNotFound = "USER_NOT_FOUND"
	FileNotFound = "FILE_NOT_FOUND"

	LimitingExceededMaxBalance        = "LIMITING_EXCEEDED_MAX_BALANCE"
	LimitingExceededMaxPerDay         = "LIMITING_EXCEEDED_MAX_PER_DAY"
	LimitingExceededMaxPerMonth       = "LIMITING_EXCEEDED_MAX_PER_MONTH"
	LimitingExceededMaxPerTransaction = "LIMITING_EXCEEDED_MAX_PER_TRANSACTION"
	LimitingExceededMaxSingleTransfer = "LIMITING_EXCEEDED_MAX_SINGLE_TRANSFER"
)

type Error struct {
	Parent         error
	Code           string
	PublicMessages string
}

func (e *Error) Error() string {
	return e.PublicMessages
}

func CreateError(err error, code string, mess string) *Error {
	return &Error{
		err,
		code,
		mess,
	}
}
