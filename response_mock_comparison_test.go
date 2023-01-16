package rem

import "testing"

func TestResponseMockComparison(t *testing.T) {
	t.Run("Status", func(t2 *testing.T) {
		res1 := NewMockResponse()
		res2 := NewMockResponse()


		for i := 100 ; i < 500 ; i += 15 {
			for j := 100 ; j < 500 ; j += 15 {
				res1.Status(i)
				res2.Status(j)

				if i == j {
					if !res1.CompareStatus(res2) {
						t2.Errorf("%d == %d", i, j)
					}
				} else {
					if res1.CompareStatus(res2) {
						t2.Errorf("%d != %d", i, j)
					}
				}
			}
		}
	})

	t.Run("Headers", func(t2 *testing.T) {
		res1 := NewMockResponse()
		res2 := NewMockResponse()

		if !res1.CompareHeaders(res2) {
			t2.Error("Both writers had no headers")
		}

		res1.Header("Content-Type", "application/json")
		if res1.CompareHeaders(res2) {
			t2.Error("First writer had 1 more header")
		}

		res2.Header("Content-Type", "application/json")
		if !res1.CompareHeaders(res2) {
			t2.Error("Both writers had same headers")
		}

		res2.Header("Server", "magic")
		if res1.CompareHeaders(res2) {
			t2.Error("Second writer had 1 more header")
		}

		res1.Header("Server", "two")
		if res1.CompareHeaders(res2) {
			t2.Error("Both writers had same number of headers, but different values")
		}
	})

	t.Run("Body", func(t2 *testing.T) {
		res1 := NewMockResponse()
		res2 := NewMockResponse()

		if !res1.CompareBody(res2) {
			t2.Error("Both bodies were empty")
		}

		res1.Bytes([]byte("a"))
		if res1.CompareBody(res2) {
			t2.Error("First body had 'a' in it")
		}

		res2.Bytes([]byte("b"))
		if res1.CompareBody(res2) {
			t2.Error("First body had 'a' in it and second body had 'b' in it")
		}

		res1.Bytes([]byte("b"))
		if !res1.CompareBody(res2) {
			t2.Error("Bodies were the same")
		}

		res1.Text("b")
		if !res1.CompareBody(res2) {
			t2.Error("Bodies were the same")
		}

		res2.Text("a")
		if res1.CompareBody(res2) {
			t2.Error("First body had 'b' in it and second body had 'a' in it")
		}

		res1.JSON(testStructure{A:1, B:"aa"})
		if res1.CompareBody(res2) {
			t2.Error("First body had an object in it and second body had 'a' in it")
		}

		res2.JSON(testStructure{A:2, B:"aa"})
		if res1.CompareBody(res2) {
			t2.Error("Both bodies had an object in it but the objects were different")
		}

		res2.JSON(testStructure{A:1, B:"aa"})
		if !res1.CompareBody(res2) {
			t2.Error("Both bodies were the same")
		}
	})
}