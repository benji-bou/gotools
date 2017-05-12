package httputil

import (
	"testing"
)

func BenchmarkGetImageUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetImageUrl("http://s1.lemde.fr/image/2016/10/24/314x157/5019131_7_6e1b_2016-09-09-6c8233f-15016-1gg2s9e_8c5162f6f92866d22109368793afe1de.jpg")
	}
}
