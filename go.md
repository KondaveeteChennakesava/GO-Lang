# Go Programming Learning Track

## 1. Go Basics

### Installation

To get started, install Go by following the official instructions:

- [Go Downloads](https://go.dev/dl/)
- After installing, check your version:
    ```shell
    $ go version
    ```
    Expected output:
    ```
    go version go1.21.0 linux/amd64
    ```

#### Setup Checklist

- Go 1.21+ installed, available with `go version`
- GOPATH and workspace configured ([Go Workspace Guide](https://go.dev/doc/gopath_code))
- VSCode plugin: [Go extension by Google](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- VSCode: enable linting, formatting on save, and install [gopls](https://github.com/golang/tools/blob/master/gopls/README.md)

### Go Style Guide

Familiarize yourself with Go conventions and idiomatic code:

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Go uses [gofmt](https://pkg.go.dev/cmd/gofmt) for formatting—run this regularly on your code.

### Tutorial

- Official [Go Tour](https://tour.golang.org/welcome/1): Do all chapters up to "Methods and Interfaces". Go through others as time permits.
- Beginner video series: [JustForFunc: Programming in Go](https://www.youtube.com/playlist?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ)

*Tip:* Many great Go tutorials exist online. When you're comfortable with the concepts above, mark this section complete.

### Go Advanced Concepts

- Interfaces & Composition: [Go Interfaces Explained](https://gobyexample.com/interfaces)
- Goroutines & Channels: [Concurrency in Go](https://gobyexample.com/goroutines), [Go Channels](https://gobyexample.com/channels)
- Error handling: [Go Error Handling Best Practices](https://blog.golang.org/error-handling-and-go)

## Data Project: IPL Data Analytics [Go Version]

### Guidelines

General guidelines:
* Organize code in functions, e.g., `calculate()`, `plot()`, `main()`.
* Use structs rather than raw slices for CSV records. Prefer well-named fields.
* Minimize memory consumption where possible.
* Descriptive variable and constant names.
* Create `.gitignore` using: [Go .gitignore](https://github.com/github/gitignore/blob/main/Go.gitignore)
* Include a README.md for setup/run instructions.
* Use `go mod` for dependency management.
* Ensure code passes `golint` and `go vet`.

### Sample IPL Project

#### Dataset

Use the [IPL Dataset](https://www.kaggle.com/manasgarg/ipl/version/5).

#### Instructions

1. Download the data; if you can't access, consult a mentor or try alternate sources.
2. Initialize a Go project: `go mod init <project-name>`.
3. Use the [encoding/csv](https://pkg.go.dev/encoding/csv) package to read data.
4. All project files should be in a dedicated repository.
5. README.md and proper code comments required.

#### Tasks (Implement with Go and suitable charting library, e.g., [go-chart](https://github.com/wcharczuk/go-chart), or output data for plotting with Python/Excel):

- Compute total runs per team.
- Compute top batsman for Royal Challengers Bangalore (top 10 by runs).
- Analyze foreign umpires (requires joining umpire data with country source).
- Stacked chart: matches played per team per season.
- Bar chart: matches played per year.
- Stacked chart: matches won per team per year.
- Bar chart: extra runs conceded per team (year 2016).
- Bar chart: top 10 economical bowlers (year 2015).

## Project 2: Company Master Data - Maharashtra

Source: [Company Master Data](https://data.gov.in/catalog/company-master-data)

#### Problems

- Histogram: Authorized Capital (bucket intervals matching Python example).
- Bar Plot: Company registration count by year.
- Bar Plot: Company registration in 2015 by district (map zip code to district, e.g., [Mumbai Pin Codes](https://www.goldenchennai.com/pin-code/maharashtra-postal-code/)).
- Grouped Bar Plot: Registration counts by principal business activity and year (top 5 activities, last 10 years).

*Implement with Go CSV, and output for plotting, or use [go-echarts](https://github.com/go-echarts/go-echarts) for advanced charting.*

## Project 3: UN Population Data Analytics

Source: [UN Population Data CSV](https://datahub.io/core/population-growth-estimates-and-projections/r/population-estimates.csv)

#### Tasks

- Bar Chart: India population over years.
- Bar Chart: ASEAN countries' population for 2014 ([ASEAN Countries](https://en.wikipedia.org/wiki/ASEAN))
- Bar Chart: Total SAARC population over years ([SAARC Countries](https://en.wikipedia.org/wiki/South_Asian_Association_for_Regional_Cooperation))
- Grouped Bar Chart: ASEAN population by years (2004–2014).

## Game Project

Familiarize with Go game development libraries (simpler than Pygame):
- [Ebiten](https://ebiten.org/documents/overview.html) (2D games)
- [Raylib-go](https://github.com/gen2brain/raylib-go)
- Try building a simple game: Pong, Snake, or a small shooter.

General mechanics:
- Game loop
- Sprite/object coordinates
- Rendering
- Score display
- End condition

## Testing in Go

- Use Go's [testing package](https://golang.org/pkg/testing/) for unit tests.
- For code coverage: [`go test -cover`](https://golang.org/pkg/testing/#hdr-Testing_flags).
- Create a small test fixture dataset (CSV) and write tests for core transformation logic.

## Go Self Review

Write your own technical notes or short docs on:
- Array and slice functions (e.g., `append`, slicing).
- String functions (`strings` package: `Split`, `Replace`, etc.).
- Structs, interfaces, and Go's approach to OOP.

***

## References and Resources

- [Go Official Website](https://go.dev/)
- [Go Documentation](https://pkg.go.dev/)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Wiki](https://github.com/golang/go/wiki)
- [Gophercises](https://gophercises.com/) (Project-based learning)

### For Data and Charting

- [go-chart](https://github.com/wcharczuk/go-chart)
- [go-echarts](https://github.com/go-echarts/go-echarts)
- [CSV handling](https://pkg.go.dev/encoding/csv)

*Feel free to expand or adjust these projects and resources based on your interests or specific data sets.*
