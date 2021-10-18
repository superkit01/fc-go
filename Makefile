build-img:
        docker build -t fc-go-runtime  -f build-image/Dockerfile build-image

build: build-img
        docker run --rm -it -v $$(pwd):/tmp fc-go-runtime bash -c "go build -o /tmp/code/bootstrap /tmp/code/main.go"
        chmod +x code/bootstrap  && mv code/bootstrap ./ && zip code.zip bootstrap && rm -rf bootstrap

deploy: build
        fun deploy -y
