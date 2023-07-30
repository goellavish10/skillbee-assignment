# Static Site Generator

The static site generator application developed as an assignment is a tool that automates the process of generating static web pages from dynamic data fetched via an API. During the build time, the application makes use of the API to retrieve data based on the number of pages specified in an environment variable.

## Tech Stack

**Client:** HTML, Plain CSS

**Server:** Golang (Go Fiber)

## Run Locally

Clone the project

```bash
  git clone https://github.com/goellavish10/skillbee-assignment
```

Install Go (v1.20.5 or above)

For Mac OS

```bash
  brew install go
```

Others: https://go.dev/dl/

Install Air

```bash
  # binary will be $(go env GOPATH)/bin/air
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

  # or install it into ./bin/
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

  air -v
```

Start the development server

```bash
  make server
```

Build the project

```bash
  make static
```

Run the Production Build

```bash
  make run
```

Generate Desired Static Pages

```bash
  # where number after PAGES= will define number of static pages to generate
  PAGES=30 make run

  # or while running dev server
  PAGES=30 make server
```

## Documentation

### Configuration

Make a .env file in Root Directory setting **PORT** env variable

To generate desired number of static pages while running dev server or production build, set the **PAGES** variable to the number of static pages required.

If **PAGES** env variable is not set the generator will default to 10 pages.

### Additional Functionality

**The Home Route**

It contains a simple form asking for number of pages to generate. If the user wishes to not generate staic pages via defining on terminal and need an interface then enter the number of static pages to generate in the input and press generate button.

![Reference Image](https://i.ibb.co/9NCyqH1/Screenshot-2023-07-30-at-4-03-10-PM.png)

**API**

Similarly you can make a POST Request on home route with **x-www-form-urlencoded** format for generating static pages by defining number of pages in the key **numberOfPages**

## API Reference

#### Get static pages

```http
  GET /staic/:pageId
```

| Parameter | Type  | Description                            |
| :-------- | :---- | :------------------------------------- |
| `pageId`  | `int` | **Required**. Static Page ID (integer) |

#### Generate Static Pages

```http
  POST /
```

| Parameter       | Type  | Description                                      |
| :-------------- | :---- | :----------------------------------------------- |
| `numberOfPages` | `int` | **Required**. Number of static pages to generate |
