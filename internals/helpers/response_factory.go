package helpers

/**
 * ResponseFactory provides helper functions to create standardized HTTP responses.
 */
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * JSONResponse writes a JSON response to the gin context with the given status code and data.
 * @param c The gin context.
 * @param statusCode The HTTP status code.
 * @param data The data to be written as JSON.
 */
func JSONResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

/**
 * ErrorResponse writes an error response to the gin context with the given status code and error.
 * The response will be a JSON object with the key "error" and the value of the error's message.
 * @param c The gin context.
 * @param statusCode The HTTP status code.
 * @param err The error to be written as JSON.
 */
func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err})
}

/**
 * SuccessResponse writes a successful response to the gin context with the given status code and data.
 * The response will be a JSON object with the key "data" and the value of the data.
 * @param c The gin context.
 * @param statusCode The HTTP status code.
 * @param data The data to be written as JSON.
 */
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

/**
 * BadRequestResponse writes a bad request error response to the gin context with the given error.
 * The response will be a JSON object with the key "error" and the value of the error's message.
 * @param c The gin context.
 * @param err The error to be written as JSON.
 */
func BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

/**
 * ValidationErrorResponse writes a validation error response to the gin context with the given error messages.
 * The response will be a JSON object with the key "errors" and the value of the error messages.
 * @param c The gin context.
 * @param err The error messages to be written as JSON.
 */
func ValidationErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"errors": ValidateInputJSON(c, err)})
}
