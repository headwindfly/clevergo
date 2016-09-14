// CSRF Middleware

package csrfmiddleware

import (
	"errors"
	"github.com/clevergo/csrf"
	"github.com/clevergo/utils"
	"github.com/headwindfly/clevergo"
	"github.com/valyala/fasthttp"
	"strings"
)

const csrfKey = "_csrf"

var (
	errInvalid        = "Unable to verify your data submission."
	errTokenNotExists = errors.New("The token does not exists.")
	errTokenInvalid   = errors.New("The token is invalid.")
	// Default safe methods.
	safeMethods = map[string]bool{
		"GET":     true,
		"HEAD":    true,
		"OPTIONS": true,
		"TRACE":   true,
	}

	errHandler = func(ctx *clevergo.Context) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(errInvalid)
	}
)

// CSRFMiddleware Cross Site Request Forgery Protection Middleware.
type CSRFMiddleware struct {
	len          int    // length of token.
	maskLen      int    // length mask.
	key          string // be used to save token into Context.UserValue.
	sessionKey   string // be used to save true token into session.
	formKey      string // be used to get token from form's args.
	headerKey    string // be used to get token from header.
	safeMethods  map[string]bool
	errorHandler clevergo.HandlerFunc
}

// NewCSRFMiddleware returns CSRF Middleware instance.
func NewCSRFMiddleware() CSRFMiddleware {
	return CSRFMiddleware{
		len:          32,
		maskLen:      8,
		key:          csrfKey,
		sessionKey:   csrfKey,
		formKey:      csrfKey,
		headerKey:    "X-CSRF-Token",
		safeMethods:  safeMethods,
		errorHandler: errHandler,
	}
}

// Len returns length of the CSRF token.
func (m CSRFMiddleware) Len() int {
	return m.len
}

// SetLen for setting token's length.
func (m *CSRFMiddleware) SetLen(len int) {
	m.len = len
}

// MaskLen returns mask's length.
func (m CSRFMiddleware) MaskLen() int {
	return m.maskLen
}

// SetMaskLen for setting mask's length.
func (m *CSRFMiddleware) SetMaskLen(len int) {
	m.maskLen = len
}

// Key returns key.
func (m CSRFMiddleware) Key() string {
	return m.key
}

// SetKey for setting key.
func (m *CSRFMiddleware) SetKey(key string) {
	m.key = key
}

// SessionKey returns sessionKey.
func (m CSRFMiddleware) SessionKey() string {
	return m.sessionKey
}

// SetSessionKey for setting sessionKey.
func (m *CSRFMiddleware) SetSessionKey(key string) {
	m.sessionKey = key
}

// HeaderKey returns headerKey.
func (m CSRFMiddleware) HeaderKey() string {
	return m.headerKey
}

// SetHeaderKey for setting headerKey.
func (m *CSRFMiddleware) SetHeaderKey(key string) {
	m.headerKey = key
}

// FormKey returns formKey.
func (m CSRFMiddleware) FormKey() string {
	return m.formKey
}

// SetFormKey for setting formKey.
func (m *CSRFMiddleware) SetFormKey(key string) {
	m.formKey = key
}

// SetSafeMethods for set safe methods.
func (m *CSRFMiddleware) SetSafeMethods(methods []string) {
	m.safeMethods = map[string]bool{}
	for i := 0; i < len(methods); i++ {
		m.safeMethods[strings.ToUpper(methods[i])] = true
	}
}

// SetErrorHandler for setting error handler.
func (m *CSRFMiddleware) SetErrorHandler(handler clevergo.HandlerFunc) {
	m.errorHandler = handler
}

// Handle implemented the Middleware Interface.
func (m CSRFMiddleware) Handle(next clevergo.Handler) clevergo.Handler {
	return clevergo.HandlerFunc(func(ctx *clevergo.Context) {
		trueToken, err := m.tureToken(ctx)

		_, safe := m.safeMethods[string(ctx.Method())]

		// Validate CSRF token.
		if !safe {
			if (err != nil) ||
				((csrf.Validate(m.maskLen, ctx.FormValue(m.formKey), trueToken) != nil) &&
					(csrf.Validate(m.maskLen, ctx.Request.Header.Peek(m.headerKey), trueToken) != nil)) {
				m.errorHandler(ctx)
				return
			}
		}

		// Create a new encode token and save it.
		token := csrf.Generate(m.maskLen, trueToken)

		// Save the encoded token into UserValue.
		ctx.SetUserValue(m.key, token)

		// Save the true token
		ctx.Session.Values[m.sessionKey] = trueToken

		next.Handle(ctx)
	})
}

// Get true token from session.
// Returns non-nil error if the token does not exists or invalid.
func (m CSRFMiddleware) tureToken(ctx *clevergo.Context) ([]byte, error) {
	token, ok := ctx.Session.Values[m.sessionKey]
	if !ok || token == nil {
		return utils.RandomBytes(m.len), errTokenNotExists
	}
	if v, ok := token.([]byte); ok {
		return v, nil
	}
	return utils.RandomBytes(m.len), errTokenInvalid
}

// Token returns the encoded token.
func (m CSRFMiddleware) Token(ctx *clevergo.Context) interface{} {
	return ctx.UserValue(m.key)
}
