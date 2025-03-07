package system

import (
	"encoding/json"
	"github.com/godbus/dbus/v5"
)

// DBusWrapper – обёртка для системных действий, предназначенная для экспорта через DBus.
type DBusWrapper struct {
	actions *Actions
}

// NewDBusWrapper создаёт новую обёртку для системных действий.
func NewDBusWrapper(a *Actions) *DBusWrapper {
	return &DBusWrapper{actions: a}
}

// Install – обёртка над Actions.Install.
func (w *DBusWrapper) Install(packageName string) (string, *dbus.Error) {
	resp, err := w.actions.Install(packageName)
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// Update – обёртка над Actions.Update.
func (w *DBusWrapper) Update(packageName string) (string, *dbus.Error) {
	resp, err := w.actions.Update(packageName)
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// Info – обёртка над Actions.Info.
func (w *DBusWrapper) Info(packageName string) (string, *dbus.Error) {
	resp, err := w.actions.Info(packageName)
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// Search – обёртка над Actions.Search.
func (w *DBusWrapper) Search(packageName string) (string, *dbus.Error) {
	resp, err := w.actions.Search(packageName)
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// Remove – обёртка над Actions.Remove.
func (w *DBusWrapper) Remove(packageName string) (string, *dbus.Error) {
	resp, err := w.actions.Remove(packageName)
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// ImageGenerate – обёртка над Actions.ImageGenerate.
func (w *DBusWrapper) ImageGenerate(switchFlag bool) (string, *dbus.Error) {
	resp, err := w.actions.ImageGenerate(switchFlag)
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// ImageUpdate – обёртка над Actions.ImageUpdate.
func (w *DBusWrapper) ImageUpdate() (string, *dbus.Error) {
	resp, err := w.actions.ImageUpdate()
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}

// ImageSwitch – обёртка над Actions.ImageSwitch.
func (w *DBusWrapper) ImageSwitch() (string, *dbus.Error) {
	resp, err := w.actions.ImageSwitch()
	if err != nil {
		return "", dbus.MakeFailedError(err)
	}
	data, jerr := json.Marshal(resp)
	if jerr != nil {
		return "", dbus.MakeFailedError(jerr)
	}
	return string(data), nil
}
