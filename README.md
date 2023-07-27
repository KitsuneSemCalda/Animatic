# Animatic - AnimeFire Anime Downloader for Plex

## Overview

Animatic is a Go-based application designed to download anime episodes from AnimeFire and prepare them for indexing in Plex. The purpose of Animatic is to simplify the process of obtaining anime content and making it easily accessible for Plex users to enjoy their favorite anime series.

## Features

- Efficiently downloads anime episodes from AnimeFire.
- Automatically organizes downloaded episodes into a Plex-compatible file structure.
- Provides metadata retrieval and integration with Plex for easy indexing.
- Simple and user-friendly command-line interface.
- Lightweight and fast.

## Requirements

- Go (version X.X.X or higher)
- Internet connection to access AnimeFire and Plex servers.

## How to Use

1. Clone the Animatic repository from GitHub:

```bash
git clone https://github.com/KitsuneSemCalda/Animatic.git
```

2. Navigate to project directory

```bash
cd Animatic
```

3. Build the executable

```bash
go build && sudo mv main /usr/bin/Animatic
```

or 

```bash
sudo bash install.sh
```

4.Run the Animatic

```bash
Animatic
```

## Working

Animatic will download the anime episodes and save them in a folder named after the anime, organized for Plex compatibility.

To add the downloaded anime to your Plex media library, follow Plex's usual procedure for adding new media folders.

Enjoy streaming your anime collection through Plex.


## Configuration
By default, Animatic uses sensible configuration settings, but you can customize them as needed. You can find the configuration file settings.txt in the project's root directory.

## Contribution
We welcome contributions to Animatic! If you find any issues, have ideas for improvement, or want to add new features, please open an issue or submit a pull request on GitHub.

### Disclaimer
Animatic is a project created for educational and personal use only. It is essential to comply with copyright laws and ensure that you have the right to access and use the content you download through AnimeFire and Plex.

### License
Animatic is released under the MIT License. See the LICENSE file for more details.


