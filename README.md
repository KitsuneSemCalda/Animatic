<p align="center">
  <img src="https://github.com/VitorCarvalho67/faster-news/assets/102667323/4652d5ad-0ce4-4282-8112-9c9c6137b551" />
</p>

# Animatic
[![GitHub license](https://img.shields.io/github/license/KitsuneSemCalda/Animatic)](KitsuneSemCalda/Animatic/blob/master/LICENSE) ![GitHub stars](https://img.shields.io/github/stars/KitsuneSemCalda/Animatic) ![GitHub stars](https://img.shields.io/github/languages/count/KitsuneSemCalda/Animatic) ![GitHub stars](https://img.shields.io/github/languages/top/KitsuneSemCalda/Animatic) ![GitHub stars](https://img.shields.io/github/repo-size/KitsuneSemCalda/Animatic) ![GitHub stars](https://img.shields.io/github/languages/code-size/KitsuneSemCalda/Animatic)

Animatic is a Go-based application designed to search and watch anime episodes from the web, as well as to provide the option to download them. It offers a command-line interface for users to input the name of the anime they wish to access or download. The project is particularly geared towards those who prefer interacting with the application through the terminal, making it a streamlined and efficient way to search for and view anime episodes for enthusiasts of the genre. Its main focus is on anime content with **Portuguese** dubbing and subtitles.

## Features

- Checks for internet connection at startup.

- Searches for the requested anime on AnimeFire.

- Downloads all episodes of the selected anime.

- Provides error messages for various failure cases such as lack of internet connection, failure to locate the anime, and inability to access episode URLs.

## How to use

Create a directory chromeMedia in root:

```bash
sudo mkdir /chromeMedia
```

If chromeMedia is a root owned directory, you need change owner to your user
```bash
sudo chown $USER:$USER /chromeMedia
```

Compile the project in your enviroment:
```bash
go build
```

Setting Path to new local:
```bash
export PATH=$PATH:~/.local/bin
```

Move the project from some directory in your path
```bash
mv animatic ~/.local/bin/animatic
```

## Running:

Run in your shell:

```bash
animatic
```

The code open the prompt-ui label to you write the anime name to be downloaded.

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch
3. Commit your Changes
4. Push to the Branch
5. Open a Pull Request