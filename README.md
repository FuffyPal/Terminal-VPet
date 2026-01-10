# Terminal V-Pet

**Terminal V-Pet** is a lightweight, CLI-based virtual pet simulation written in Rust. Take care of your digital companion directly from your terminal by managing its hunger, energy, and happiness levels.

## Features

- **Hunger Management**: Feed your pet to keep it well-fed.
- **Energy Management**: Play with your pet to boost its energy.
- **Happiness Management**: Interact with your pet to increase its happiness.

## Installation & Usage

### Download Pre-built Binaries
You can download the latest version of Terminal V-Pet directly from the releases page: [Download Releases](https://gitlab.com/FluffyPal/terminal-vpet/-/releases)

### Running with Docker

If you prefer not to install Rust locally, you can build and run the application using [Docker](https://www.docker.com/):

```shell
# Build the image
docker build -t terminal-vpet .

# Run the application
docker run -it terminal-vpet
```

If you prefer not to install Rust locally, you can build and run the application using [Podman](https://podman.io/docs/installation):

```shell
# Build the image
podman build -t terminal-vpet .

# Run the application
podman run -it terminal-vpet
```

### Building from Source

Ensure you have the [Rust toolchain](https://rustup.rs/) installed.

```shell
git clone https://gitlab.com/FluffyPal/terminal-vpet.git
cd terminal-vpet
cargo build
clear
cargo run
```

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Please refer to our [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

---

**Maintained by [FluffyPal](https://gitlab.com/FluffyPal)**
