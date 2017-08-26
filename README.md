# setup-pi

![](https://raw.github.com/motdotla/ansible-pi/master/ansible-pi.jpg)

Quickly setup your Raspberry Pi.

## How to use

Clone to your host computer

```
go get https://github.com/Code-Hex/setup-pi.git
cd $GOPATH/src/github/Code-Hex/setup-pi
```

And, we need [`go-bindata`](https://github.com/jteeuwen/go-bindata) to compile source code.

```
go get -u github.com/jteeuwen/go-bindata/...
```

If you are completed above flow, you can compile source code after put some files on the `files` directory. Files in the `files` directory are also compiled together, and a single binary file is created.

```
make all
```

## How to write recipe

You must need to write `items` key at first line.  
Next, you should write

    - name: 'task name'
per your task. The whole atmosphere will be as follows.
```yaml
items:
  - name: 'Configure WIFI'
    copy:
      src: ./conf/wpa_supplicant.conf
      dest: /etc/wpa_supplicant/wpa_supplicant.conf
      mode: 0600

  - name: 'Install packages using apt'
    apt:
      packages:
        - git
        - openssl
        - build-essential 
        - guvcview
        - tightvncserver
      before_update: True
      is_upgrade: True

  - name: 'Configure dhcpd'
    copy:
      src: ./conf/dhcpcd.conf
      dest: /etc/dhcpcd.conf

  - name: 'Configure autorun_date'
    copy:
      src: ./init.d/autorun_date
      dest: /etc/init.d/autorun_date
      mode: 0755

  - name: 'Register autorun_date to update-rc.d'
    command: update-rc.d /etc/init.d/autorun_date defaults
```

You can use these keys for the task.
    
    apt, copy, get_url, command

### ・`apt`
`apt` have `packages` and `before_update`.
`packages` have value to install packages.  
- `before_update`: Bool. It update repositories before do relate apt task.
- `is_upgrade`: Bool. If you have the upgrade list, these will upgrade.

```yaml
- name: 'install using apt'
  apt:
    packages:
      - git
      - openssl
      - build-essential 
      - guvcview
      - tightvncserver
    before_update: True # bool
    is_upgrade: True    # bool
```

### ・`copy`
`copy` have `src` and `dest`, `mode`.
Copy the compiled file together to the specified path.
- `src`: You specify the path as the base of `files` directory.  
- `dest`: You specify destenetion path.
- `mode`: permission.
```yaml
- name: 'Copy file'
  copy:
    src: ./init.d/autorun_date
    dest: /etc/init.d/autorun_date
    mode: 0755
```
### ・`get_url`
`get_url` have `url` and `dest`.  
You can download file if you use this key.
- `url`: Specify the URL.
- `dest`: Specify the download destination.
```yaml
- name: 'Download rclone'
   get_url:
     url: "https://downloads.rclone.org/rclone-v1.37-linux-arm.zip"
     dest: "/home/pi/"
```
### ・`command`
You can run command. like this.
```yaml
- name: 'run command'
  command: echo Hello World
```
