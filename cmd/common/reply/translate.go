package reply

import (
	"apm/lib"
)

// TranslateKey принимает ключ и возвращает английский текст.
func TranslateKey(key string) string {
	switch key {
	case "package":
		return lib.T_("Package")
	case "count":
		return lib.T_("Count")
	case "isConsole":
		return lib.T_("Console Application")
	case "packageInfo":
		return lib.T_("Package Information")
	case "install":
		return lib.T_("Install")
	case "store":
		return lib.T_("Storage Type")
	case "timestamp":
		return lib.T_("Date")
	case "imageDigest":
		return lib.T_("Image Digest")
	case "os":
		return lib.T_("Distribution")
	case "container":
		return lib.T_("Container")
	case "name":
		return lib.T_("Name")
	case "extraInstalled":
		return lib.T_("Extra Installed")
	case "upgradedCount":
		return lib.T_("Upgraded Count")
	case "bootedImage":
		return lib.T_("Booted Image")
	case "removedPackages":
		return lib.T_("Removed Packages")
	case "providers":
		return lib.T_("Providers")
	case "version":
		return lib.T_("Version")
	case "history":
		return lib.T_("History")
	case "depends":
		return lib.T_("Dependencies")
	case "installedSize":
		return lib.T_("Installed Size")
	case "removedCount":
		return lib.T_("Removed Count")
	case "upgradedPackages":
		return lib.T_("Upgraded Packages")
	case "packageName":
		return lib.T_("Package Name")
	case "image":
		return lib.T_("Image")
	case "commands":
		return lib.T_("Commands")
	case "maintainer":
		return lib.T_("Maintainer")
	case "versionInstalled":
		return lib.T_("Installed Version")
	case "remove":
		return lib.T_("Remove")
	case "containers":
		return lib.T_("Containers")
	case "paths":
		return lib.T_("Paths")
	case "description":
		return lib.T_("Description")
	case "date":
		return lib.T_("Date")
	case "newInstalledCount":
		return lib.T_("Newly Installed Count")
	case "active":
		return lib.T_("Active")
	case "info":
		return lib.T_("Information")
	case "totalCount":
		return lib.T_("Total Count")
	case "installed":
		return lib.T_("Installed")
	case "manager":
		return lib.T_("Package Manager")
	case "lastChangelog":
		return lib.T_("Last Changelog")
	case "section":
		return lib.T_("Section")
	case "spec":
		return lib.T_("Specification")
	case "booted":
		return lib.T_("Booted")
	case "staged":
		return lib.T_("Staged")
	case "size":
		return lib.T_("Size")
	case "newInstalledPackages":
		return lib.T_("Newly Installed Packages")
	case "notUpgradedCount":
		return lib.T_("Not Upgraded Count")
	case "containerName":
		return lib.T_("Container Name")
	case "config":
		return lib.T_("Configuration")
	case "exporting":
		return lib.T_("Exporting")
	case "status":
		return lib.T_("Status")
	case "imageDate":
		return lib.T_("Image Date")
	case "packages":
		return lib.T_("Packages")
	case "filename":
		return lib.T_("Filename")
	case "containerInfo":
		return lib.T_("Container Information")
	case "imageName":
		return lib.T_("Image Name")
	case "transport":
		return lib.T_("Transport")
	case "pinned":
		return lib.T_("Pinned")
	case "list":
		return lib.T_("List")
	default:
		return lib.T_(key)
	}
}
