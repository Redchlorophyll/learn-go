package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("before unit test")
	m.Run()
	fmt.Println("after unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "hello Eko" {
		t.Fail() // will not stop test function if end up here
		// t.Error("error reason here") <- this will show what error is and run t.Fail() afterward
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		HelloWorld("Dhonni")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("able to return correct string", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			HelloWorld("Dhonni")
		}
	})

	b.Run("able to return correct string", func(b *testing.B) {
		for i := 0; i < b.N; i += 1 {
			HelloWorld("Eko")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	tests := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(eko)",
			request: "djo",
		},
		{
			name:    "HelloWorld(eko)",
			request: "djo",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i += 1 {
				HelloWorld(test.request)
			}
		})
	}
}

func TestHelloWorldDhonni(t *testing.T) {
	result := HelloWorld("Dhonni")
	if result != "hello Dhonni" {
		t.FailNow() // will stop test function if found this
		// t.Fatal("error reason here") <- this will show what error is and run t.FailNow() afterward
	}
}

func TestHelloWorldUsingAssert(t *testing.T) {
	result := HelloWorld("Dhonni")
	assert.Equal(t, "hello Dhonni", result, "Result must be 'hello Dhonni'")
}

func TestHelloWroldUsingRequire(t *testing.T) {
	result := HelloWorld("Dhonni")
	require.Equal(t, "hello Dhonni", result, "Result must be 'hello Dhonni'")
}

func TestHelloWroldUsingGOOS(t *testing.T) {
	result := HelloWorld("Dhonni")

	if runtime.GOOS == "darwin" {
		t.Skip("can not run on MacOS")
	}

	require.Equal(t, "hello Dhonni", result, "Result must be 'hello Dhonni'")
}

func TestHelloWorldWithSubTest(t *testing.T) {
	t.Run("able to return correct string", func(t *testing.T) {
		result := HelloWorld("Dhonni")

		require.Equal(t, "hello Dhonni", result, "Result must be 'hello Dhonni'")
	})
}

func TestHelloWorldWithTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(eko)",
			request:  "djo",
			expected: "hello djo",
		},
		{
			name:     "HelloWorld(eko)",
			request:  "djo",
			expected: "hello djo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
