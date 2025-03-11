# ssha - Simple SSH Host Manager

`ssha` is a command-line tool written in Go that provides a simple and efficient way to manage and connect to your SSH hosts. It uses Bubble Tea to provide an interactive terminal user interface (TUI).

## Features

* **Host Management:**
    * Stores host configurations in a `config.json` file located in `~/.ssha/`.
    * Supports both password-based and private key-based authentication.
    * Displays a table of configured hosts with alias, hostname, user, and port.
    * Planned features:
        * Add, update, delete, and search hosts.
* **SSH Connections:**
    * Connects to selected hosts via SSH directly within the terminal.
    * Handles both password and private key authentication.
    * Displays connection errors within the TUI.
* **Interactive TUI:**
    * Uses Bubble Tea for a clean and interactive terminal user interface.
    * Table-based host display with selection and navigation.
* **Future Features:**
    * File transfer with split window terminal.
    * Improved TUI.
    * Better Error handling.
    * Autocomplete and cursor control inside the ssh session.

## Getting Started

### Prerequisites

* Go 1.21 or later.
* `sshpass` if using password authentication.

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/saini128/ssha
    cd ssha
    ```

2.  **Build the application:**

    ```bash
    go build
    ```

3.  **Run the application:**

    ```bash
    ./ssha
    ```

### Configuration

* Host configurations are stored in `~/.ssha/config.json`.
* You can manually edit this file to add, modify, or remove hosts.
* The application also provides features to add, update, and delete hosts through the UI.

### Usage

1.  **Run the application:** `./ssha`
2.  **Navigate:** Use the arrow keys to navigate the host table.
3.  **Connect:** Press `Enter` to connect to the selected host.
4.  **Quit:** Press `Esc` or `q` to quit the application.

## Planned Enhancements

* **Improved TUI:** Enhance the user interface with better styling and navigation.
* **Host Management Features:** Implement features to add, update, delete, and search hosts within the TUI.
* **Enhanced Error Handling:** Improve error handling for connection failures and display more informative error messages.
* **File Transfer:** Add a split-window terminal feature for file and folder transfer between the local and remote systems.
* **Autocomplete and Cursor Control:** Add autocomplete and cursor control features to the ssh session within the go application.
* **Host Key Verification:** Implement proper host key verification for secure connections.

## Security Considerations

* **Private Key Authentication:** It is highly recommended to use private key authentication instead of password authentication for better security.
* **Host Key Verification:** In future releases, proper host key verification will be implemented.
* **`sshpass`:** If using password authentication, `sshpass` is used, which is generally considered insecure. Use with caution.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to report bugs or suggest new features.

## License

This project is licensed under the MIT License.