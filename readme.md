# GoCR Worker for Debian

This is a guide to stand up a GoCR worker for Debian.

## Primary steps
1. `sudo apt update` To get fresh package list
1. Install GoLang
2. Install Tesseract-OCR



### Install Go
1. `sudo rm -rf /usr/local/go` Remove existing installs
3.  `cd /usr/local` 
4. Download 1.22.2 `sudo wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz`
5. Unpack the tar `sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz`
6. Add go to PATH `export PATH=$PATH:/usr/local/go/bin`
7. Confirm install `go version` should output version info

### Install Tesseract-OCR
1. `sudo apt install tesseract-ocr libtesseract-dev` Install tesseract and header files/libraries
2. `sudo apt install tesseract-ocr-eng` was probably already installed
3. `tesseract --version` Confirm Tesseract version