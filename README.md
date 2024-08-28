# Age Checker Api

## Overview

This is a web service that processes receipts and calculates points based on given rules.

## Running the Application

### Locally

npm run build
npm start

### Using Docker

docker build -t receipt-processor .
docker run -p 3000:3000 receipt-processor
