package repositorymemory

import (
	"context"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/config"
	"github.com/lekan-pvp/go-musthave-shortener-tpl.git/internal/models"
	"testing"
)

var mr = MemoryRepository{}

func BenchmarkMemoryRepository_InsertUserRepo(b *testing.B) {
	b.Run("Insert User Repo", func(b *testing.B) {
		cfg := config.Config{
			FileStoragePath: "test.json",
		}
		mr.New(&cfg)
		urls := models.URLs{
			UUID:          "324234234",
			ShortURL:      "http://localhost:8080/RXggLwCj",
			OriginalURL:   "http://google.com",
			CorrelationID: "123",
			DeleteFlag:    false,
		}
		ctx := context.Background()

		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := mr.InsertUserRepo(ctx, urls.UUID, urls.ShortURL, urls.OriginalURL)
			if err != nil {
				panic(err)
			}
		}
	})
}
