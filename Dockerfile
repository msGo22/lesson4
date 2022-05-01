## Kullanacağımız ana image yani golang'in 1.18 alpine işletim sistemindeki halini seçtik
FROM golang:1.18-alpine
## Container içinde çalışma alanımızı ayarladık
WORKDIR /uygulama
## Tüm dosyaları içeri aktardık
COPY . ./
## tüm ihtiyaç duyduğumuz paketleri indirtip sonrasında build ettik
RUN go mod tidy
RUN go build -v -o server ./withStructure/app/main.go

## Çalıştırdık
CMD ["/uygulama/server"]
