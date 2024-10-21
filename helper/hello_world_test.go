package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Rangga", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rangga")
		}
	})
	b.Run("Dwi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dwi")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rangga")
	}
}
func BenchmarkHelloDwi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dwi")
	}
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Rangga",
			request: "Rangga",
		}, {
			name:    "Dwi",
			request: "Dwi",
		}, {
			name:    "Rangga Dwi Mahendra Ganteng",
			request: "Rangga Dwi Mahendra Ganteng",
		}, {
			name:    "Lionel Andreas Messi Ronaldo Critiano",
			request: "Lionel Andreas Messi Ronaldo Critiano",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			HelloWorld(benchmark.request)
		})
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Rangga)",
			request:  "Rangga",
			expected: "Hello Rangga",
		}, {
			name:     "HelloWorld(Dwi)",
			request:  "Dwi",
			expected: "Hello Dwi",
		}, {
			name:     "HelloWorld(Mahendra)",
			request:  "Mahendra",
			expected: "Hello Mahendra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("rangga", func(t *testing.T) {
		result := HelloWorld("rangga")
		require.Equal(t, "Hello rangga", result, "Result must be 'Hello rangga'")
	})
	t.Run("dwi", func(t *testing.T) {
		result := HelloWorld("dwi")
		require.Equal(t, "Hello dwi", result, "Result must be 'Hello dwi'")
	})
	t.Run("mahendra", func(t *testing.T) {
		result := HelloWorld("mahendra")
		require.Equal(t, "Hello mahendra", result, "Result must be 'Hello mahendra'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")

}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		// unit test failed
		// kode program selanjutnya tetap dieksekusi meskipun gagal
		// t.Fail()

		// sama aja cmn bisa log errornya karena apa dan di Fatal juga manggil fungsi Fail
		t.Error("Result must be 'Hello World'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloRangga(t *testing.T) {
	result := HelloWorld("Rangga")
	if result != "Hello Rangga" {
		// unit test failed
		// kode program tidak akan dilanjutkan dan langsung digagalkan
		// t.FailNow()

		// sama aja cmn bisa log errornya karena apa dan di Fatal juga manggil fungsi Failnow
		t.Fatal("Result must be 'Hello Rangga'")
	}

	fmt.Println("TestHelloRangga Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Rangga")
	// kalo pake assert dia ketika gagal akan memanggil Fail
	assert.Equal(t, "Hello Rangga", result, "Result must be 'Hello Rangga'")
	fmt.Println("TestHelloWorldAssertion Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rangga")
	// sedangkan require dia ketika gagal akan memanggil FailNow
	require.Equal(t, "Hello Rangga", result, "Result must be 'Hello Rangga'")
	fmt.Println("TestHelloWorldRequire Done")
}

// skip test, untuk menskip sebuah unit test jika memang kondisi unit test tidak sesuai requirement
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows")
	}

	result := HelloWorld("Rangga")
	// kalo pake assert dia ketika gagal akan memanggil Fail
	assert.Equal(t, "Hello Rangga", result, "Result must be 'Hello Rangga'")
}

func TestHelloDwi(t *testing.T) {
	result := HelloWorld("Dwi")
	if result != "Hello Dwi" {
		t.Fatal("Result must be 'Hello Dwi'")
	}

	fmt.Println("TestHelloDwi Done")
}
