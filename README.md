# Web-scraper-go

## Environment Setup / Pre-requisite

> Please install the following for running OS environment before continue:

1. [Go](https://go.dev/doc/install)

---

## Quick Start

1. Git clone current repository to your local environment and open with preferred IDE [eg: VS code].
2. Open terminal in `project root path` and run following command to start

```bash
# Start web scrapper function
go run main.go
```

3. Result will be stored as JSON data in `generated` folder.
   - All quotes extracted from [example link here](https://quotes.toscrape.com) as examples
   - User can change page to be extracted from [this line of code](./main.go#L63). Max valid page is `10`

<!-- Instruction to guide user change page URL for bookstore -->

---

## Sample Extracted Quotes

```json
[
  {
    "quoteText": "The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking.",
    "author": "Albert Einstein",
    "authorUrl": "https://quotes.toscrape.com/author/Albert-Einstein",
    "tags": ["change", "deep-thoughts", "thinking", "world"]
  },
  {
    "quoteText": "It is our choices, Harry, that show what we truly are, far more than our abilities.",
    "author": "J.K. Rowling",
    "authorUrl": "https://quotes.toscrape.com/author/J-K-Rowling",
    "tags": ["abilities", "choices"]
  },
  {...},
  {...},
  {...},
]
```

---

## Third-party Golang package/modules installed

- Colly
