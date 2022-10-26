package modmanager

type ForgeVersion struct {
	MinecratVersin    string
	ForgeVersion      string
	InstallerLocation string
	Sha256            string
}

type ForgeRepo struct {
	Versions []ForgeVersion
}

type Mod struct {
	Name           string
	ID             string
	MinecratVersin string
	ForgeVersion   string
	Sha256         string
}

type ModRepo struct {
	Mods []Mod
}

