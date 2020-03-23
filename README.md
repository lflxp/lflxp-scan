IP scan and port scan

scan是扫描模块，提供：

* ip扫描
    * 192.168.0.1
* ip段扫描
    * 192.168.1-10.55
    * 192.168.1.1-255
* 端口扫描
    * 1-65535
* 扫描结果展示

![](https://github.com/lflxp/lflxp-scan/blob/master/asset/scan.png)
![](https://github.com/lflxp/lflxp-scan/blob/master/asset/scan1.png)

## Install

```
git clone https://github.com/lflxp/lflxp-scan
cd lflxp-scsan
make install
lflxp-scan -h
```

## Usage

Scan is the GUI display of showme terminal, providing dynamic parameter prompt function

> showme

```bash
=> host@127.0.0.1 # showme
>>> scan 192.168.50.1-255
            192.168.50.1-255  192网段  
            192.168.40.1-255  192网段  
            192.168.1.1-255   192网段  
```

## Args

`IP segment parameter description`

* Single IP
* Supports IP of [-] specified range, including a, B, C and D segments

`GUI operating instructions`

- Tab: Next View
- Enter: Select IP/Commit Input
- F5: Input New Scan IP or Port range
- ↑ ↓: Move View
- ^c: Exit
- F1: Help
- Space: search result with ip view and port view