package infrastructure

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func NewGormLogger(appEnv string) logger.Interface {
	if appEnv == "development" {
		return logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
}

// catatan pribadiku 
// 1. logger untuk development dan production berbeda
// 2. logger.New ( buat instance logger baru )
// 3. os.Stdout ( tampilkan hasilnya di terminal )
// 4. "\r\n" ( buat baris baru pada output log )
// 5. log.LstdFlags ( tambahkan informasi waktu pada log )
// 6. logger.Config ( konfigurasi logger )
// 7. SlowThreshold ( batas waktu untuk query yang dianggap lambat )
// 8. LogLevel ( tingkat log yang akan ditampilkan )
// 9. logger.Warn menampilkan peringatan
// 10. logger.Info log lebih detail, termasuk query SQL
// 11. IgnoreRecordNotFoundError ( abaikan error jika record tidak ditemukan )  -> ( true )
// 12. ParameterizedQueries menyembunyikan nilai asli parameter query dari log -> ( true )
// 13. Colorful ( menampilkan log dengan warna ) -> ( true )
