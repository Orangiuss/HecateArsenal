<h1 align="center">
  <img src="img/logo_vf.png" alt="Hecate" width="200px">
  <br>
</h1>

<h1 align="center">
  <br>
    Hecate 🧙🏻‍♀️🔮🪄 - One-liners Framework
  <br>
  <br>
</h1>

<p align="center">
  <a href="https://github.com/Orangiuss/hecate/blob/main/README.md">English</a> 
    -
  <a href="https://github.com/Orangiuss/hecate/blob/main/README_FR.md">Français</a>
    <!-- -
  <a href="https://github.com/Orangiuss/hecate/blob/main/README_CN.md">中文</a>
    -
  <a href="https://github.com/Orangiuss/hecate/blob/main/README_JP.md">日本語</a> -->
</p>

`hecate` is a powerful and versatile framework designed to streamline your offensive security and penetration testing workflows. It offers a rich collection of one-liners, enabling efficient and effective security assessments.

![HecateDemo](img/HecateArsenal.png)

## Description 📝

Hecate provides a comprehensive arsenal of one-liners covering various security tasks. From reconnaissance to post-exploitation, this framework empowers security professionals to perform a wide range of operations quickly and with minimal effort.

## Features ⚙️

- **Extensive One-Liners Collection:** Perform tasks such as:
  - Reconnaissance 🕵️‍♀️
  - Vulnerability Scanning 🔍
  - Exploitation 💥
  - Post-Exploitation 🛠️
  - Privilege Escalation 👑
  - Maintaining Access 🔐
  - Web Application Security Testing 🕸️
  <!-- - Wireless Security Assessments 📡 -->
  - Social Engineering 🎭
  - Password Cracking 🔑
- **Customization:** Easily add and tailor your own one-liners to fit specific testing needs.
- **Efficiency:** Execute commands rapidly, saving time during testing processes.

## Installation 🛠️

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/Orangiuss/hecate.git
    ```

2. **Navigate to the Directory:**

    ```bash
    cd hecate
    ```

3. **Build the Framework:**

    ```bash
    setup.sh
    ```

## Usage 🚀

1. Run the framework with the following command:

    ```bash
    hecate-cli --help
    ```

2. **Example Command:**

    ```bash
    hecate-cli tool list
    hecate-cli ol list
    hecate-cli ol run example-one-liner
    ```

## Customization 🛠️

To add custom one-liners, create new yaml files in the `one-liners` directory. Follow the existing structure and syntax for consistency :

```yaml
id: nmap-http-ports

info:
  name: Nmap HTTP Ports
  default-category: recon
  author: Orangius
  complexity: low
  description: |
    Scan for open HTTP ports using Nmap
  metadata:
    verified: true
  tags:
    - nmap
    - recon
    - port-scanning
    - beginner
    - low
    - simple

variables:
  DOMAIN:
    description: The target domain
    required: true

one-liner:
  cmd: |
    nmap -sV -p 80,443 {DOMAIN}
```

## Roadmap 🗺

- `November 2024` - **Version 0.1** - Alpha Release with Core Features (CLI)
- `December 2024` - **Version 1.0** - Initial Release with Core Features (GUI + CLI)
- `January 2025` - **Version 1.1** - Bug Fixes and Enhancements
- `Summer 2025` - **Version 1.2** - New Features and Improvements
- ? : **Version 2.0** - Advanced Features and Modules + TUI project ?

## Disclaimer ⚠️

Hecate is intended for authorized penetration testing and security assessments. Misuse of this tool is illegal and unethical. The authors are not responsible for any unauthorized activities.

## Contributing 🤝

We welcome contributions! Feel free to submit pull requests for bug fixes, new one-liners, or documentation improvements.

## License 📜

GPL-3.0 license (see [LICENSE](https://github.com/Orangiuss/HecateArsenal/tree/main?tab=GPL-3.0-1-ov-file))

## Contact 📧

For questions or feedback, contact us via [provide contact information, e.g., email address, GitHub repository issue tracker].

## Additional Considerations

- **Security:** Always use this tool responsibly and ethically.
- **Version Control:** Stay updated with the latest changes by following the project's GitHub repository.
