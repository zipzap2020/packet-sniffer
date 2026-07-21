# Go Network Packet Sniffer

A lightweight, efficient network packet sniffer written in Go using the `gopacket` library. This tool captures, decodes, and analyzes live network traffic in real time, providing structured insight into multiple network layers.

## Features

* **List Available Network Interfaces:** List all available network interfaces on your device using the `-m` flag.
* **Live Interface Selection:** Choose a specific network interface to capture traffic using the `-i` flag.
* **BPF Filtering:** Apply custom Berkeley Packet Filters on the fly via the `-f` flag (e.g., capturing only specific ports or protocols).
* **Save to .pcap File:** Save output to a .pcap file using the `-t` flag. If no .pcap extension is provided, it will be added automatically.
* **Layer Decoding:** Parsers for IPv4 addresses, TCP/UDP ports, and active DNS query extraction.
* **Smart Application Payload Formatting:** Displays raw application data using a clean Hex dump structure, preventing terminal layout breaking caused by non-printable binary characters.


## Prerequisites

Since this tool hooks into the network stack using `pcap`, you need the C development headers for `libpcap` installed on your system.

For Debian/Ubuntu based systems, run:

```bash
sudo apt update
sudo apt install libpcap-dev
```

Make sure you also have Go installed (version 1.18 or higher recommended).

## Installation

* **Clone the repository:**

```bash
git clone https://github.com/zipzap2020/packet-sniffer.git
cd packet-sniffer
```

* **Download the required dependencies:**

```bash
go mod tidy
```

## Usage

Network packet capture requires root or administrative privileges to put the network interface into promiscuous mode.

```bash
#list network interfaces
sudo ./sniffer -m
# Capture from a specific interface with a filter and save to a .pcap file
sudo ./sniffer -i eth0 -f "tcp port 80" -t output
```

* **Running directly with Go:**

```bash
sudo go run main.go
```
* **Building the executable:**

```bash
go build -o sniffer main.go
sudo ./sniffer
```
