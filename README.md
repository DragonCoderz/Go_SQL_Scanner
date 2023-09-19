# MySQL Scanner

## Introduction

This Go program is a MySQL scanner that detects whether MySQL is running on a specified host and port. It initiates a MySQL handshake with the server and retrieves basic information about the MySQL instance's configuration.

## Usage

### Prerequisites

- Go (Golang) should be installed on your system. If it's not installed, you can download it from [the official website](https://golang.org/dl/).

### Installing MySQL (Depending on Your OS)

#### macOS

1. You can use [Homebrew](https://brew.sh/) to install MySQL on macOS. Open a terminal and run the following commands:

   ```shell
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```

2. Install MySQL with Homebrew:

   ```shell
   brew install mysql
   ```

3. Start the MySQL service:

   ```shell
   brew services start mysql
   ```

#### Windows

1. You can use [MySQL Installer](https://dev.mysql.com/downloads/installer/) to install MySQL on Windows.

2. Download the MySQL Installer for Windows and run it.

3. Follow the installation wizard, which allows you to select various MySQL components and configurations.

4. Once the installation is complete, MySQL should be running as a Windows service.

#### Linux (Ubuntu/Debian)

1. Open a terminal and run the following commands to install MySQL:

   ```shell
   sudo apt update
   sudo apt install mysql-server
   ```

2. During the installation, you will be prompted to set a root password for MySQL.

3. Start the MySQL service:

   ```shell
   sudo systemctl start mysql
   ```

### Running the Scanner

1. Clone or download this repository to your local machine.

2. Open a terminal and navigate to the directory containing the `mysql_scanner.go` file.

3. Modify the `host` and `port` variables in the `mysql_scanner.go` file to specify the target host's IP address and port where MySQL is running (default is `127.0.0.1:3306`).

4. Run the following command to execute the scanner:

   ```shell
   go run mysql_scanner.go
   ```

### Example Output

The scanner will provide output indicating whether MySQL appears to be running on the specified host and port, and it will display information about the MySQL instance's configuration, such as version, capabilities, status, and the server plugin name.

## Testing

You can test the MySQL scanner by following these steps:

1. Ensure that MySQL is running on the target host and port you specified in the `mysql_scanner.go` file.

2. Run the scanner as mentioned in the "Running the Scanner" section above.

3. Verify that the scanner correctly detects MySQL and provides accurate information about the MySQL instance.

## Disclaimer

1. This scanner is a basic tool and does not perform in-depth analysis or security checks. Always ensure you have the necessary permissions and authorization before scanning any host or server.
2. I was only able to test this code on MacOS. There may be version compatibility issues with other Operating Systems. If that's the case, please make a pull request, and I will be sure to make revisements to my code.


---
