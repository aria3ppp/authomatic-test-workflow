//go:generate go-enum --file=$GOFILE --marshal

package response

/*
ENUM(
OK
NotFound
InvalidURLParameter
InvalidRequest
EmailAlreadyUsed
EmailNotFound
IncorrectPassword
SameNewPassword
TokenInvalid
TokenMissingOrMalformed
InternalServerError
)
*/
type Status int
