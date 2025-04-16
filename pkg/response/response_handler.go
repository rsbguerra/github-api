package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
)

// HandleGithubErrors processes an error and sends an appropriate HTTP response
// based on the error type and status code.
//
// Supported Errors:
// - 400 Bad Request: if the request is invalid.
// - 401 Unauthorized: invalid token or authentication failed.
// - 403 Forbidden: user does not have permission to access the repository.
// - 404 Not Found: repository does not exist for a given owner and repo name.
// - 422 Unprocessable Entity: if the request is invalid.
// - 500 Internal Server Error: if an error occurs while retrieving the pull requests.
//
// Parameters:
// - c: *gin.Context - the Gin context used to send the HTTP response.
// - err: error - the error to be handled.
func HandleGithubErrors(c *gin.Context, err error) {
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			switch ghErr.Response.StatusCode {
			// 400 Bad Request: if the request is invalid
			case 400:
				StatusBadRequest(c)
			// 401 Unauthorized: invalid token or authentication failed
			case 401:
				StatusUnauthorized(c)
			// 403 Forbidden: user does not have permission to access the repository.
			case 403:
				StatusForbidden(c)
			// 404 Not Found: repository does not exist for a given owner and repo name
			case 404:
				StatusNotFound(c)
			// 422 Unprocessable Entity: if the request is invalid
			case 422:
				StatusUnprocessableEntity(c, err)
			// 500 Internal Server Error: if an error occurs while retrieving the pull requests
			case 500:
				StatusInternalServerError(c, err)
			default:
				StatusInternalServerError(c, err)
			}
		}
	}
}
