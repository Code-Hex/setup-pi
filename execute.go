package setup

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

var client = http.DefaultClient

func init() {
	client.Timeout = time.Duration(900) * time.Second
}

func (s *Setup) execute() error {
	for idx, item := range s.Recipe.Items {
		err := s.invoke(item)
		if err != nil {
			return errors.Wrapf(err, "Failed to invoke item[%d]: %s", idx, item.Name)
		}
	}
	return nil
}

func (s *Setup) invoke(item *Item) error {
	if item.Apt != nil {
		if err := s.invokeApt(item); err != nil {
			return errors.Wrap(err, "Failed to invokeApt")
		}
	}

	if item.Copy != nil {
		if err := s.invokeCopy(item); err != nil {
			return errors.Wrap(err, "Failed to invokeCopy")
		}
	}

	if item.GetURL != nil {
		if err := s.invokeGetURL(item); err != nil {
			return errors.Wrap(err, "Failed to invokeGetURL")
		}
	}

	if err := s.invokeCommand(item); err != nil {
		return errors.Wrap(err, "Failed to invokeCommand")
	}

	return nil
}

func (s *Setup) invokeApt(item *Item) error {
	apt := item.Apt
	if apt.BeforeUpdate {
		cmd := exec.Command("apt", "update")
		cmd.Env = s.Env
		if err := cmd.Run(); err != nil {
			return exit.MakeOSErr(err)
		}
	}

	args := append([]string{"install", "-y"}, item.Packages...)
	cmd := exec.Command("apt", args...)
	cmd.Env = s.Env
	if err := cmd.Run(); err != nil {
		return exit.MakeOSErr(err)
	}
	return nil
}

func (s *Setup) invokeCopy(item *Item) error {
	copy := item.Copy
	info, err := AssetInfo(copy.Src)
	if err != nil {
		return exit.MakeOSFile(err)
	}
	b, err := Asset(copy.Src)
	if err != nil {
		return exit.MakeOSFile(err)
	}

	fi, err := os.Stat(copy.Dest)
	mode := fi.Mode()
	if !mode.IsDir() && err != nil {
		return exit.MakeOSFile(err)
	}

	if mode.IsDir() {
		dest, err := filepath.Abs(fi.Name())
		if err != nil {
			return exit.MakeOSFile(err)
		}
		f, err := os.Create(filepath.Join(dest, info.Name()))
		if err != nil {
			return exit.MakeOSFile(err)
		}
		defer f.Close()
		f.Write(b)
		if copy.Mode > 0 {
			if err := f.Chmod(copy.Mode); err != nil {
				return exit.MakeOSFile(err)
			}
		}
		return nil
	}

	dest, err := filepath.Abs(fi.Name())
	if err != nil {
		return exit.MakeOSFile(err)
	}
	f, err := os.Create(dest)
	if err != nil {
		return exit.MakeOSFile(err)
	}
	defer f.Close()
	if copy.Mode > 0 {
		if err := f.Chmod(copy.Mode); err != nil {
			return exit.MakeOSFile(err)
		}
	}
	f.Write(b)
	return nil
}

func (s *Setup) invokeGetURL(item *Item) error {
	get := item.GetURL
	resp, err := client.Get(get.URL)
	if err != nil {
		return exit.MakeSoftWare(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(get.Dest)
	if err != nil {
		return exit.MakeOSFile(err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil && err != io.EOF {
		return exit.MakeOSFile(err)
	}
	return nil
}

func (s *Setup) invokeCommand(item *Item) error {
	split := strings.Split(item.Command, " ")
	cmd := exec.Command(split[0], split[1:]...)
	cmd.Env = s.Env
	if err := cmd.Run(); err != nil {
		return exit.MakeOSErr(err)
	}
	return nil
}
