package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Umar")
	}
}

func BenchmarkHelloBawazir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bawazir")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Umar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Umar")
		}
	})

	b.Run("Bawazir", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bawazir")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	tests := []struct {
		name    string
		request string
	}{
		{
			name:    "Umar",
			request: "Umar",
		},
		{
			name:    "Bawazir",
			request: "Bawazir",
		},
	}

	for _, benchmark := range tests {
		for i := 0; i < b.N; i++ {
			HelloWorld(benchmark.request)
		}
	}
}

// membuat test main agar bisa menjalankan before dan after
func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit test")

	m.Run()

	fmt.Println("Sesudah unit test")
}

// fungsi skip() akan membatalkan eksekusi unit test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run in windows")
	}

	result := HelloWorld("Umar")
	require.Equal(t, "Hello Umar", result, "Result must be 'Hello Umar'")
}

// fungsi error akan tetap melanjutkan testing
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Umar")
	if result != "Hello Umar" {
		// error
		t.Error("Result must be 'Hello Umar'")
	}

	fmt.Println("TestHelloWorld Umar")
}

// fungsi fatal akan menghentikan testing
func TestHelloWorldBawazir(t *testing.T) {
	result := HelloWorld("Bawazir")
	if result != "Hello Bawazir" {
		// error
		t.Fatal("Result must be 'Hello Bawazir'")
	}

	fmt.Println("TestHelloWorld Bawazir")
}

// assertion dan require menggunakan testify
// jika gagal langsung berhenti dan tetap menjalankan code selanjutnya
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Umar")
	assert.Equal(t, "Hello Umar", result, "Result must be 'Hello Umar'")
}

// jika gagal langsung berhenti dan tidak menjalankan code selanjutnya
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Umar")
	require.Equal(t, "Hello Umar", result, "Result must be 'Hello Umar'")
}

// sub test = membuat function test di dalam function unit test
// jika ingin menjalankan unit sub test mengunakan perintah go test -run TestSubTest/Bawazir
func TestSubTest(t *testing.T) {
	t.Run("Umar", func(t *testing.T) {
		result := HelloWorld("Umar")
		require.Equal(t, "Hello Umar", result, "Result must be 'Hello Umar'")
	})
	t.Run("Bawazir", func(t *testing.T) {
		result := HelloWorld("Bawazir")
		require.Equal(t, "Hello Bawazir", result, "Result must be 'Hello Bawazirs'")
	})
}

// table test melakukan testing dengan iterasi
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Umar)",
			request:  "Umar",
			expected: "Hello Umar",
		},
		{
			name:     "HelloWorld(Bawazir)",
			request:  "Bawazir",
			expected: "Hello Bawazir",
		},
		{
			name:     "HelloWorld(Faiz)",
			request:  "Faiz",
			expected: "Hello Faiz",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}
