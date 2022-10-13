# minitino

> A tool generate minimalist tech blog. See demo https://qunv.github.io

## Why it's awesome

It would be great if we could share something and have a site of our own.

## You should have skills

+ Markdown - to write your content

## Prerequisites

[Golang >=1.19](https://go.dev/)

## Usages

1. Clone this repo:

```shell
git clone https://github.com/qunv/minitino.git
```

2. Build

Move to `cd minitino`

```shell
go build -o <YOUR FOLDER>
```

3. Run

First create `config.yaml`:

```yaml
app:
  rootName: <YOUR NAME>
  intro: <YOUR INTRO>
  findOn:
    github: <YOUR GITHUB URL>
    twitter: <YOUR TWITTER URL>
```
Then run `./minitino`

The first run will create necessary folders that we have 2 important folders:

+ `_posts`: This folder contains all your Markdown file post
- `_about`: Write something about you. It's also Markdown file

Create posts by create markdown file inside `_posts` folder. And we have several rules:

- File name with format: `<YEAR>-<MONTH>-<DAY>-<YOUR_FILE_NAME>.md`

File contents start with 2 line:

```
[comment]: <> (Rambling a little) // your post title

[comment]: <> (self, self2) // your post tags
```

Finally, run `./minitino` one again.


## Contributing

Happy to your contribution

## License

Licensed under the [GNU License](https://github.com/qunv/minitino/blob/main/LICENSE).