# agrec

超A&Gでの放送を録音してmp3に変換し、Dropboxの任意のディレクトリにアップロードするアプリケーションです。

## Docker

Raspberry Pi Zero W向けのイメージを用意しています。

```sh
docker run -e "DROPBOX_TOKEN=yourtoken" -v /path/to/your/programs.toml:/root/programs.toml -d eiryyy/agrec
```

## Programs

`programs.toml` に予約したい番組を設定します。

```toml
[[programs]]
title = "favorite_program"
cron = "0 30 21 * * Mon" # CRON式
min = 30 # 番組の時間
```
