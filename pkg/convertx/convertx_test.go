package convertx

import (
	"testing"
)

func TestChinese2Spell(t *testing.T) {
	type args struct {
		words string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{words: "sdadadasd好人啊赛☺sdadas"}, "sdadadasdhaorenasaisdadas"},
		{"test2", args{words: "sdadadasd"}, "sdadadasd"},
		{"test3", args{words: "好人啊赛"}, "haorenasai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chinese2Spell(tt.args.words); got != tt.want {
				t.Errorf("Chinese2Spell() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkChinese2Spell(b *testing.B) {
	type args struct {
		words string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{words: "sdadadasd好人啊赛☺sdadas"}, "sdadadasdhaorenasaisdadas"},
		{"test2", args{words: "sdadadasd"}, "sdadadasd"},
		{"test3", args{words: "好人啊赛"}, "haorenasai"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chinese2Spell(tests[i%3].args.words)
	}
}
