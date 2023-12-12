package patterns

import (
	"fmt"
)

// Реализовать паттерн «строитель».
// Объяснить применимость паттерна, его плюсы и минусы, а также реальные
// примеры использования данного примера на практике.

// Паттерн «строитель» является порождающим паттерном проектирования,
// предназначен для создания сложных объектов, многосоставных объектов.

// Применимость: необходимость создания сложных объектов, для чего требуется
// много шагов и деталей. необходимость создания различных версий одного
// объекта. необходимость сокрытия деталей реализации объекта и создаваемых
// временных сущностей, возникающих при создании объекта. Детерминация состояния
// объекта до двух возможных вариантов - объект создан, объекта не существует.

// Плюсы: упрощение создания сложных объектов, возможность создания различных
// версий одного объекта меньшим количеством вызовов.

// Минусы: увеличение количества классов и объектов, увеличение сложности
// программы.

// Пример: создание объекта конфигурации для приложения, зависящего от
// операционной системы, на которой оно запущено.

// OSConfig конфигурация приложения, зависящая от операционной системы
type OSConfig struct {
	OSName    string
	OSVersion string
	OSArch    string
}

// OSConfigBuilder интерфейс для создания конфигурации приложения
type OSConfigBuilder interface {
	SetOSName() OSConfigBuilder
	SetOSVersion(OSVersion string) OSConfigBuilder
	SetOSArch(OSArch string) OSConfigBuilder
	GetOSConfig() OSConfig
}

// WindowsConfigBuilder реализация интерфейса OSConfigBuilder для Windows
type WindowsConfigBuilder struct {
	config OSConfig
}

// NewWindowsConfigBuilder создаёт новый экземпляр билдера для Windows
func NewWindowsConfigBuilder() *WindowsConfigBuilder {
	return &WindowsConfigBuilder{}
}

// SetOSName устанавливает имя операционной системы
func (w *WindowsConfigBuilder) SetOSName() OSConfigBuilder {
	w.config.OSName = "Windows"
	return w
}

// SetOSVersion устанавливает версию операционной системы
func (w *WindowsConfigBuilder) SetOSVersion(OSVersion string) OSConfigBuilder {
	w.config.OSVersion = OSVersion
	return w
}

// SetOSArch устанавливает архитектуру операционной системы
func (w *WindowsConfigBuilder) SetOSArch(OSArch string) OSConfigBuilder {
	w.config.OSArch = OSArch
	return w
}

// GetOSConfig возвращает конфигурацию приложения
func (w *WindowsConfigBuilder) GetOSConfig() OSConfig {
	return w.config
}

// LinuxConfigBuilder реализация интерфейса OSConfigBuilder для Linux
type LinuxConfigBuilder struct {
	config OSConfig
}

// NewLinuxConfigBuilder создаёт новый экземпляр билдера для Linux
func NewLinuxConfigBuilder() *LinuxConfigBuilder {
	return &LinuxConfigBuilder{}
}

// SetOSName устанавливает имя операционной системы
func (l *LinuxConfigBuilder) SetOSName() OSConfigBuilder {
	l.config.OSName = "Linux"
	return l
}

// SetOSVersion устанавливает версию операционной системы
func (l *LinuxConfigBuilder) SetOSVersion(OSVersion string) OSConfigBuilder {
	l.config.OSVersion = OSVersion
	return l
}

// SetOSArch устанавливает архитектуру операционной системы
func (l *LinuxConfigBuilder) SetOSArch(OSArch string) OSConfigBuilder {
	l.config.OSArch = OSArch
	return l
}

// GetOSConfig возвращает конфигурацию приложения
func (l *LinuxConfigBuilder) GetOSConfig() OSConfig {
	return l.config
}

// OSConfigDirector директор, который управляет процессом создания конфигурации
// приложения
type OSConfigDirector struct {
	builder OSConfigBuilder
}

// SetBuilder устанавливает билдер для директора
func (d *OSConfigDirector) SetBuilder(builder OSConfigBuilder) {
	d.builder = builder
}

// BuildOSConfig создаёт конфигурацию приложения
func (d *OSConfigDirector) BuildOSConfig(OSVersion, OSArch string) OSConfig {
	d.builder.SetOSName().SetOSVersion(OSVersion).SetOSArch(OSArch)
	return d.builder.GetOSConfig()
}

// NewOSConfigDirector создаёт новый экземпляр директора
func NewOSConfigDirector() *OSConfigDirector {
	return &OSConfigDirector{}
}

// NewOSConfig создаёт новый экземпляр конфигурации приложения
func NewOSConfig(OSVersion, OSArch string) OSConfig {
	director := NewOSConfigDirector()
	var builder OSConfigBuilder
	switch OSName := getOSName(); OSName {
	case "Windows":
		builder = NewWindowsConfigBuilder()
	case "Linux":
		builder = NewLinuxConfigBuilder()
	default:
		panic("Unknown OS")
	}
	director.SetBuilder(builder)
	return director.BuildOSConfig(OSVersion, OSArch)
}

// getOSName возвращает имя операционной системы
func getOSName() string {
	var OSName string
	fmt.Print("Enter OS name: ")
	fmt.Scan(&OSName)
	return OSName
}

// пример использования
// func main() {
// 	OSVersion := "10.0.18363.0"
// 	OSArch := "amd64"
// 	config := NewOSConfig(OSVersion, OSArch)
// 	fmt.Printf("OS name: %s\nOS version: %s\nOS arch: %s\n", config.OSName, config.OSVersion, config.OSArch)
// }
