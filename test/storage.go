/*
Contains various storage benchmarks
*/

package test

var Common = Test{
	Name:   "Common Test Files",
	Prereq: nil,
	Deps: []Dependency{
		Dependency{
			Name:               "build-essential",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "build-essential"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "build-essential"},
		},
		Dependency{
			Name:               "automake",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "automake"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "automake"},
		},
		Dependency{
			Name:               "cmake",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "cmake"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "cmake"},
		},
		Dependency{
			Name:               "wget",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "wget"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "wget"},
		},
		Dependency{
			Name:               "tar",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "tar"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "tar"},
		},
		Dependency{
			Name:               "libtool",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "libtool"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "libtool"},
		},
		Dependency{
			Name:               "bison",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "bison"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "bison"},
		},
		Dependency{
			Name:               "flex",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "flex"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "flex"},
		},
		Dependency{
			Name:               "git",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "git"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "-y", "git"},
		},
	},
}

var Filebench = Test{
	Name:   "Filebench",
	Prereq: &Common,
	Deps: []Dependency{
		Dependency{
			Name:               "build-essential",
			checkCommand:       "dpkg-query",
			checkCommandArgs:   []string{"-W", "-f=${Status}", "build-essential"},
			isInstalledStr:     "install ok installed",
			installCommand:     "apt",
			installCommandArgs: []string{"install", "build-essential"},
		},
	},
}
