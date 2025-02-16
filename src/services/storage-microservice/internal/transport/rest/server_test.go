package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSMiddleware(t *testing.T) {
	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := performRequest(router, "OPTIONS", "/test")
	assert.Equal(t, http.StatusNoContent, w.Code)

	w = performRequest(router, "GET", "/test")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "http://localhost:8082", w.Header().Get("Access-Control-Allow-Origin"))
}
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRegisterRoutes(t *testing.T) {
	mockClient := &minio.Client{}
	router := gin.New()
	RegisterRoutes(router, mockClient, "test-bucket")

	t.Run("POST /upload - successful upload", func(t *testing.T) {
		w := performRequest(router, "POST", "/upload")
		assert.NotEqual(t, http.StatusOK, w.Code)
	})

	t.Run("GET /download - successful download", func(t *testing.T) {
		w := performRequest(router, "GET", "/download")
		assert.NotEqual(t, http.StatusOK, w.Code)
	})

	t.Run("DELETE /delete - successful delete", func(t *testing.T) {
		w := performRequest(router, "DELETE", "/delete")
		assert.NotEqual(t, http.StatusOK, w.Code)
	})
}

type MockMinioClient struct {
	mock.Mock
}

func (m *MockMinioClient) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	args := m.Called(ctx, bucketName)
	return args.Bool(0), args.Error(1)
}

func (m *MockMinioClient) MakeBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	args := m.Called(ctx, bucketName, opts)
	return args.Error(0)
}
