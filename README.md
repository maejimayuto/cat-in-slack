# 猫の宅急便

## どんなことを？
slack に画像を送るだけのサービスです。

## Architechture
- heroku で github connect させている
- main ブランチにマージされると deploy される
- Heroku scheduler で 1日に1回画像を送信している
- [The Cat API \- Cats as a Service\.](https://thecatapi.com/) から画像を取得している

## Load Map
slack app として、公開

### Idea
1. このアプリを導入しているチャンネルでは、猫のスタンプが押されると、そのスレッドに猫の画像が送信される
1. 配信頻度を slack から設定できるようにする

などなど、あまり良いアイデアはない

# License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.
