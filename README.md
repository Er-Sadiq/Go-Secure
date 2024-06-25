# GoSecure 🚀

GoSecure is a command-line tool developed in Golang for cloud-based anti-malware analysis. It allows users to scan their data using open-source cloud-based anti-malicious systems directly from their terminal.

## Features 🌟

- **Cloud-Based Scanning**: Utilizes cloud-based anti-malicious systems for efficient and up-to-date malware analysis. ☁️🛡️
- **Command-Line Interface**: User-friendly CLI powered by Cobra, offering commands for seamless operation. 💻✨
- **Easy Integration**: Built with resty for straightforward integration with RESTful APIs of cloud anti-malware services. 🔗⚙️

## Installation 🛠️

To install GoSecure, follow these steps:

```bash
$ git clone [ https://github.com/Er-Sadiq/Go-Secure.git ]
$ cd GoSecure
$ go build
```

Then move the built executable to your system's binary directory:

```bash
$ sudo mv go-secure /usr/local/bin
$ ./go-secure
Enter your Name: <Your-Name>
Enter Your API Key: <API Key>
Enter Your File Path: <File-Path>
```

## Demo 🎥

![GoSecure](https://github.com/Er-Sadiq/Go-Secure/assets/125464939/63110286-49e8-4339-970e-5fcbac5235bd)

## FAQ ❓

#### Does it require an API key? 🔑

Yes, it requires an API key from VirusTotal for now. More functionality will be added as development progresses.

## Contribution 🤝

Contributions are welcome! Fork the repo and submit a pull request with your enhancements. 🙌

## License 📄

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/). See the LICENSE file for details.

---

If you find GoSecure useful, please leave a like on my repository! 👍 Your support helps me improve and expand the tool. ❤️
