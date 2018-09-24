# DOKUMENTASI LINE BOT GO

## INSTALL
Environment yang perlu di install
- Golang
- Ngrok (https://ngrok.com/)
- Library Line Bot SDK (https://github.com/line/line-bot-sdk-go)

## SETUP
- Buat akun di https://developers.line.me
- Buat Project Messaging API
- Pasang Access Token dan Client Secret pada Dashboard Channel Setting project yang telah anda buat ke `config/config.go`
- Jalankan Aplikasi Golang dengan `go run main.go`
- Jalankan ngrok dengan `./ngrok http {PORT-GOLANG}`
- Tambahkan url ngrok yang https ke Dashboard Channel Setting project yang telah anda buat bagian Webhook URL
- Ye jalan

