package service

import (
	"apm/cmd/common/helper"
	"apm/lib"
	"bufio"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// AltProvider реализует методы для работы с пакетами в ALT linux
type AltProvider struct {
	servicePackage *PackageService
}

// NewAltProvider возвращает новый экземпляр AltProvider.
func NewAltProvider(servicePackage *PackageService) *AltProvider {
	return &AltProvider{
		servicePackage: servicePackage,
	}
}

// GetPackages обновляет базу пакетов, выполняет поиск и отмечает установленные пакеты.
func (p *AltProvider) GetPackages(ctx context.Context, containerInfo ContainerInfo) ([]PackageInfo, error) {
	updateCmd := fmt.Sprintf("%s distrobox enter %s -- sudo apt-get update", lib.Env.CommandPrefix, containerInfo.ContainerName)
	if _, stderr, err := helper.RunCommand(updateCmd); err != nil {
		return nil, fmt.Errorf("не удалось обновить базу пакетов: %v, stderr: %s", err, stderr)
	}

	command := fmt.Sprintf("%s apt-cache dumpavail", lib.Env.CommandPrefix)
	cmd := exec.Command("sh", "-c", command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("ошибка запуска команды: %w", err)
	}

	// Получаем список установленных пакетов.
	installedPackages, err := p.servicePackage.GetAllApplicationsByContainer(ctx, containerInfo)
	if err != nil {
		lib.Log.Error("Ошибка получения установленных пакетов: ", err)
		installedPackages = []string{}
	}

	// Формируем карту для быстрого поиска установленных пакетов.
	installedMap := make(map[string]bool)
	for _, name := range installedPackages {
		installedMap[name] = true
	}

	const maxCapacity = 1024 * 1024 * 350 // 350MB
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(stdout)
	scanner.Buffer(buf, maxCapacity)

	var packages []PackageInfo
	var pkg PackageInfo
	var currentKey string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if pkg.Name != "" {
				packages = append(packages, pkg)
				pkg = PackageInfo{}
				currentKey = ""
			}
			continue
		}

		if parts := strings.SplitN(line, ":", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			currentKey = key

			switch key {
			case "Package":
				pkg.Name = value
			case "Version":
				versionValue, errVersion := helper.GetVersionFromAptCache(value)
				if errVersion != nil {
					pkg.Version = value
				} else {
					pkg.Version = versionValue
				}
			case "Description":
				pkg.Description = value
			default:
			}
		} else {
			if currentKey == "Description" {
				pkg.Description += "\n" + line
			}
		}
	}

	if pkg.Name != "" {
		packages = append(packages, pkg)
	}

	if err = scanner.Err(); err != nil {
		if errors.Is(err, bufio.ErrTooLong) {
			return nil, fmt.Errorf("слишком большая строка: (over %dMB) - ", maxCapacity/(1024*1024))
		}
		return nil, fmt.Errorf("ошибка сканера: %w", err)
	}

	if err = cmd.Wait(); err != nil {
		return nil, fmt.Errorf("ошибка выполнения команды: %w", err)
	}

	for i := range packages {
		if installedMap[packages[i].Name] {
			packages[i].Installed = true
		}
		packages[i].Manager = "apt-get"
	}

	return packages, nil
}

// RemovePackage удаляет указанный пакет с помощью pacman -R.
func (p *AltProvider) RemovePackage(ctx context.Context, containerInfo ContainerInfo, packageName string) error {
	cmdStr := fmt.Sprintf("%s distrobox enter %s -- sudo apt-get remove -y %s", lib.Env.CommandPrefix, containerInfo.ContainerName, packageName)
	_, stderr, err := helper.RunCommand(cmdStr)
	if err != nil {
		return fmt.Errorf("не удалось удалить пакет %s: %v, stderr: %s", packageName, err, stderr)
	}
	return nil
}

// InstallPackage устанавливает указанный пакет с помощью pacman -S.
func (p *AltProvider) InstallPackage(ctx context.Context, containerInfo ContainerInfo, packageName string) error {
	cmdStr := fmt.Sprintf("%s distrobox enter %s -- sudo apt-get install -y %s", lib.Env.CommandPrefix, containerInfo.ContainerName, packageName)
	_, stderr, err := helper.RunCommand(cmdStr)
	if err != nil {
		return fmt.Errorf("не удалось установить пакет %s: %v, stderr: %s", packageName, err, stderr)
	}
	return nil
}

// GetPathByPackageName возвращает список путей для файла пакета, найденных через rpm -ql.
func (p *AltProvider) GetPathByPackageName(ctx context.Context, containerInfo ContainerInfo, packageName, filePath string) ([]string, error) {
	command := fmt.Sprintf("%s distrobox enter %s -- rpm -ql %s | grep '%s'", lib.Env.CommandPrefix, containerInfo.ContainerName, packageName, filePath)
	stdout, stderr, err := helper.RunCommand(command)
	if err != nil {
		lib.Log.Debugf("Ошибка выполнения команды: %s %s", stderr, err.Error())
		return []string{}, err
	}

	lines := strings.Split(stdout, "\n")
	var paths []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && !strings.HasSuffix(trimmed, "/") {
			paths = append(paths, trimmed)
		}
	}
	return paths, nil
}

// GetPackageOwner определяет пакет-владельца файла через rpm -qf.
func (p *AltProvider) GetPackageOwner(ctx context.Context, containerInfo ContainerInfo, filePath string) (string, error) {
	command := fmt.Sprintf("%s distrobox enter %s -- rpm -qf --queryformat '%%{NAME}' %s", lib.Env.CommandPrefix, containerInfo.ContainerName, filePath)
	stdout, stderr, err := helper.RunCommand(command)
	if err != nil {
		lib.Log.Debugf("Ошибка выполнения команды: %s %s", stderr, err.Error())
		return "", err
	}
	return strings.TrimSpace(stdout), nil
}
