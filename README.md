# Go CLI Web Scrapper

This is a CLI application built using Go Colly for web scraping and Cobra for command-line interface.

## Overview

This CLI application allows users to scrape data from web pages using Go Colly. It provides a simple and efficient way to extract information from HTML pages using CSS selectors.

## Features

- **Web Scraping**: Scrapes data from web pages using Go Colly.
- **Flexible**: Supports dynamic CSS selectors for scraping different types of data.
- **Easy-to-use CLI**: Built with Cobra, offering a user-friendly command-line interface.

## Prerequisites

Before running the application, ensure you have the following installed:

- Go (version 1.22.1)
- Go Colly library

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/puffyguy/go-web-scrapper.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-web-scrapper
    ```

3. Build the application:

    ```bash
    go build
    ```

## Example
```
./go-web-scrapper --url="https://example.com" selector --class=span.text,small.author
```