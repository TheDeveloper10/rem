package rem

import (
	"bytes"
	"testing"
)

func TestResponseMockAssignment(t *testing.T) {
	t.Run("Status", func(t2 *testing.T) {
		res := NewMockResponse()

		for i := 0 ; i < 550 ; i++ {
			res.Status(i)
			if res.StatusCode != i {
				t2.Errorf("Status is not correct (%d)", i)
			}
		}
	})

	t.Run("Headers", func(t2 *testing.T) {
		res := NewMockResponse()

		if len(res.Headers) > 0 {
			t2.Error("Starts with wrong headers!")
		}

		res.Header("Content-Type", "application/json")
		if len(res.Headers) != 1 {
			t2.Error("Failed to assign 'Content-Type' header")
		} else if ct := res.Headers.Get("Content-Type"); ct != "application/json" {
			t2.Error("Failed to assign proper value of 'Content-Type' header")
		}

		res.Header("Server", "Test")
		if len(res.Headers) != 2 {
			t2.Error("Failed to assign 'Server' header")
		} else if s := res.Headers.Get("Server"); s != "Test" {
			t2.Error("Failed to assign proper value of 'Server' header")
		}

		res.Header("tmp", "val")
		if len(res.Headers) != 3 {
			t2.Error("Failed to assign 'tmp' header")
		} else if tmp := res.Headers.Get("tmp"); tmp != "val" {
			t2.Error("Failed to assign proper value of 'tmp' header")
		}
	})

	t.Run("Body", func(t2 *testing.T) {
		res := NewMockResponse()

		if len(res.Body) != 0 {
			t2.Error("Body is not empty!")
		}

		res.Bytes([]byte("aa"))
		if !bytes.Equal(res.Body, []byte("aa")) {
			t2.Error("Body is not properly set!")
		}

		res.Text("aabbcc")
		if !bytes.Equal(res.Body, []byte("aabbcc")) {
			t2.Error("Body is not properly set!")
		}

		res.JSON(testStructure{
			A: 1,
			B: "aa",
		})
		if !bytes.Equal(res.Body, []byte("{\"ab\":1,\"cb\":\"aa\"}")) {
			t2.Error("Body is not properly set!")
		}

		res.JSON(&testStructure{
			A: 54,
			B: "123q",
		})
		if !bytes.Equal(res.Body, []byte("{\"ab\":54,\"cb\":\"123q\"}")) {
			t2.Error("Body is not properly set!")
		}
	})
}

type testStructure struct {
	A int `json:"ab"`
	B string `json:"cb"`
}