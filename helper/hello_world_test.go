package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRifqi(t *testing.T) {
	result := HelloWorld("Rifqi")
	if result != "Hello Rifqi" {
		t.Fail() // unit test fail and continue next code
	}
	fmt.Println("TestHelloWorldRifqi Done") // executed
}

func TestHelloWorldXenosty(t *testing.T) {
	result := HelloWorld("Xenosty")
	if result != "Hello Xenosty" {
		t.FailNow() // unit test fail and stopped
	}
	fmt.Println("TestHelloWorldXenosty Done") // not executed
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Xenosty")
	if result != "Hello Xenosty" {
		t.Error("Result must be Hello Xenosty") // give error message and running t.Fail() or continue next code
	}
	fmt.Println("TestHelloWorldError Done") // executed
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Xenosty")
	if result != "Hello Xenosty" {
		t.Fatal("Result must be Hello Xenosty") // unit test fail and running t.FailNow() or stopped
	}
	fmt.Println("TestHelloWorldFatal Done") // not executed
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rifqi")
	assert.Equal(t, "Hello Rifqi", result, "Result must be Hello Rifqi") // assert.Equal(t, expextedResult, actualResult, message) then execute t.Fail()
	fmt.Println("TestHelloWorld using Assertion Done")                   // executed
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rifqi")
	require.Equal(t, "Hello Rifqi", result, "Result must be Hello Rifqi") // require.Equal(t, expextedResult, actualResult, message) then execute t.FailNow()
	fmt.Println("TestHelloWorld using Require Done")                      // not executed
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" { // if running on windows, it will be skipped
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Rifqi")
	require.Equal(t, "Hello Rifqi", result, "Result must be Hello Rifqi")
}

func TestMain(m *testing.M) {
	// before (ex. connect database, etc)
	fmt.Println("BEFORE UNIT TESTING")

	// run all testing in this package
	m.Run()

	// after (ex. exit app, etc)
	fmt.Println("AFTER UNIT TESTING")
}

func TestSubTest(t *testing.T) {
	// test case 1
	t.Run("Rifqi", func(t *testing.T) {
		result := HelloWorld("Rifqi")
		require.Equal(t, "Hello Rifqi", result, "Result must be Hello Rifqi")
	})

	// test case 2
	t.Run("Aziz", func(t *testing.T) {
		result := HelloWorld("Aziz")
		require.Equal(t, "Hello Aziz", result, "Result must be Hello Aziz")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rifqi",
			request:  "Rifqi",
			expected: "Hello Rifqi",
		},
		{
			name:     "Bagas",
			request:  "Bagas",
			expected: "Hello Bagas",
		},
		{
			name:     "Kurnia",
			request:  "Kurnia",
			expected: "Hello Kurnia",
		},
		{
			name:     "Bayu",
			request:  "Bayu",
			expected: "Hello Bayu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rifqi")
	}
}

func BenchmarkHelloWorldXenosty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Xenosty")
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Rifqi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rifqi")
		}
	})
	b.Run("Xenosty", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Xenosty")
		}
	})
}

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarksTable := []struct {
		name    string
		request string
	}{
		{
			name:    "Rifqi",
			request: "Rifqi",
		},
		{
			name:    "Xenosty",
			request: "Xenosty",
		},
		{
			name:    "RifqiMuhammadAziz",
			request: "Rifqi Muhammad Aziz",
		},
	}

	for _, benchmark := range benchmarksTable {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
