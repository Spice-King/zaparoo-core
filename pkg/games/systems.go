package games

type MglParams struct {
	Delay    int
	FileType string // TODO: rename?
	Index    int
}

type FileType struct {
	Label string
	Exts  []string
	Mgl   *MglParams
}

type AltRbfOpts map[string][]string

type System struct {
	Id        string
	Name      string
	Alias     []string
	Folder    string
	Rbf       string
	AltRbf    AltRbfOpts
	FileTypes []FileType
}

const (
	AltRbfLLAPI   = "LLAPI"
	AltRbfYC      = "YC"
	AltRbfDualRAM = "DualRAM"
)

// First in list takes precendence for simple attributes in case there's a
// conflict in the future.
var CoreGroups = map[string][]System{
	"Atari7800": {Systems["Atari7800"], Systems["Atari2600"]},
	"NES":       {Systems["NES"], Systems["FDS"]},
	"Gameboy":   {Systems["Gameboy"], Systems["GameboyColor"]},
	"SMS":       {Systems["MasterSystem"], Systems["GameGear"]},
}

// FIXME: launch game > launch new game same system > not working? should it?
// TODO: setname attribute
// TODO: alternate cores
// TODO: alternate arcade folders
// TODO: custom scan function
// TODO: custom launch function
// TODO: support for multiple folders (think about symlink support here, check for dupes)
// TODO: could cut down on work scanning by folder rather than system
// TODO: add folder name aliases
// TODO: may need to support globbing on extensions

var Systems = map[string]System{
	// Consoles
	"AdventureVision": {
		Id:     "AdventureVision",
		Name:   "Adventure Vision",
		Folder: "AVision",
		Rbf:    "_Console/AdventureVision",
		FileTypes: []FileType{
			{
				Label: "Game",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Arcadia": {
		Id:     "Arcadia",
		Name:   "Arcadia 2001",
		Folder: "Arcadia",
		Rbf:    "_Console/Arcadia",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Astrocade": {
		Id:     "Astrocade",
		Name:   "Bally Astrocade",
		Folder: "Astrocade",
		Rbf:    "_Console/Astrocade",
		FileTypes: []FileType{
			{
				Exts: []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Atari2600": {
		Id:     "Atari2600",
		Name:   "Atari 2600",
		Folder: "ATARI7800",
		Rbf:    "_Console/Atari7800",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"Atari7800_LLAPI"},
			AltRbfYC:    []string{"Atari7800YC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".a26"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Atari5200": {
		Id:     "Atari5200",
		Name:   "Atari 5200",
		Folder: "ATARI5200",
		Rbf:    "_Console/Atari5200",
		AltRbf: AltRbfOpts{
			AltRbfYC: []string{"Atari5200YC"},
		},
		FileTypes: []FileType{
			{
				Label: "Cart",
				Exts:  []string{".car", ".a52", ".bin", ".rom"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"Atari7800": {
		Id:     "Atari7800",
		Name:   "Atari 7800",
		Folder: "ATARI7800",
		Rbf:    "_Console/Atari7800",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"Atari7800_LLAPI"},
			AltRbfYC:    []string{"Atari7800YC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".a78", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "BIOS",
				Exts:  []string{".rom", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"AtariLynx": {
		Id:     "AtariLynx",
		Name:   "Atari Lynx",
		Folder: "AtariLynx",
		Rbf:    "_Console/AtariLynx",
		FileTypes: []FileType{
			{
				Exts: []string{".lnx"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	// TODO: AY-3-8500
	//       Doesn't appear to have roms even though it has a folder.
	// TODO: C2650
	//       Not in official repos, think it comes with update_all.
	//       https://github.com/Grabulosaure/C2650_MiSTer
	"CasioPV1000": {
		Id:     "CasioPV1000",
		Name:   "Casio PV-1000",
		Folder: "Casio_PV-1000",
		Rbf:    "_Console/Casio_PV-1000",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"ChannelF": {
		Id:     "ChannelF",
		Name:   "Channel F",
		Folder: "ChannelF",
		Rbf:    "_Console/ChannelF",
		FileTypes: []FileType{
			{
				Exts: []string{".rom", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"ColecoVision": {
		// TODO: Remove .sg from here, keep in meta, after multi-folder.
		Id:     "ColecoVision",
		Name:   "ColecoVision",
		Alias:  []string{"Coleco"},
		Folder: "Coleco",
		Rbf:    "_Console/ColecoVision",
		AltRbf: AltRbfOpts{
			AltRbfYC: []string{"ColecoVisionYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".col", ".bin", ".rom"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "SG-1000",
				Exts:  []string{".sg"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"CreatiVision": {
		Id:     "CreatiVision",
		Name:   "VTech CreatiVision",
		Folder: "CreatiVision",
		Rbf:    "_Console/CreatiVision",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".rom", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "Bios",
				Exts:  []string{".rom", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
			{
				Label: "BASIC",
				Exts:  []string{".bas"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    3,
				},
			},
		},
	},
	// TODO: EpochGalaxy2
	//       Has a folder and mount entry but commented as "remove".
	"FDS": {
		Id:     "FDS",
		Name:   "Famicom Disk System",
		Alias:  []string{"FamicomDiskSystem"},
		Folder: "NES",
		Rbf:    "_Console/NES",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"NES_LLAPI"},
			AltRbfYC:    []string{"NESYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".fds"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "FDS BIOS",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"Gamate": {
		Id:     "Gamate",
		Name:   "Gamate",
		Folder: "Gamate",
		Rbf:    "_Console/Gamate",
		FileTypes: []FileType{
			{
				Exts: []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Gameboy": {
		Id:     "Gameboy",
		Name:   "Gameboy",
		Alias:  []string{"GB"},
		Folder: "GAMEBOY",
		Rbf:    "_Console/Gameboy",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"Gameboy_LLAPI"},
			AltRbfYC:    []string{"GameboyYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".gb"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"GameboyColor": {
		Id:     "GameboyColor",
		Name:   "Gameboy Color",
		Alias:  []string{"GBC"},
		Folder: "GAMEBOY",
		Rbf:    "_Console/Gameboy",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"Gameboy_LLAPI"},
			AltRbfYC:    []string{"GameboyYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".gbc"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Gameboy2P": {
		// TODO: Split 2P core into GB and GBC?
		Id:     "Gameboy2P",
		Name:   "Gameboy (2 Player)",
		Folder: "GAMEBOY2P",
		Rbf:    "_Console/Gameboy2P",
		FileTypes: []FileType{
			{
				Exts: []string{".gb", ".gbc"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"GameGear": {
		Id:     "GameGear",
		Name:   "Game Gear",
		Folder: "SMS",
		Rbf:    "_Console/SMS",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"SMS_LLAPI"},
			AltRbfYC:    []string{"SMSYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".gg"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"GameNWatch": {
		Id:     "GameNWatch",
		Name:   "Game & Watch",
		Folder: "GameNWatch",
		Rbf:    "_Console/GnW",
		FileTypes: []FileType{
			{
				Exts: []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"GBA": {
		Id:     "GBA",
		Name:   "Gameboy Advance",
		Alias:  []string{"GameboyAdvance"},
		Folder: "GBA",
		Rbf:    "_Console/GBA",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"GBA_LLAPI"},
			AltRbfYC:    []string{"GBAYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".gba"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"GBA2P": {
		Id:     "GBA2P",
		Name:   "Gameboy Advance (2 Player)",
		Folder: "GBA2P",
		Rbf:    "_Console/GBA2P",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"GBA2P_LLAPI"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".gba"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Genesis": {
		Id:     "Genesis",
		Name:   "Genesis",
		Alias:  []string{"MegaDrive"},
		Folder: "Genesis",
		Rbf:    "_Console/Genesis",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"Genesis_LLAPI"},
			AltRbfYC:    []string{"GenesisYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".bin", ".gen", ".md"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Intellivision": {
		Id:     "Intellivision",
		Name:   "Intellivision",
		Folder: "Intellivision",
		Rbf:    "_Console/Intellivision",
		FileTypes: []FileType{
			{
				Exts: []string{".rom", ".int", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	// TODO: Jaguar
	"MasterSystem": {
		// TODO: Split off SG-1000 (prefer Coleco core).
		Id:     "MasterSystem",
		Name:   "Master System",
		Alias:  []string{"SMS"},
		Folder: "SMS",
		Rbf:    "_Console/SMS",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"SMS_LLAPI"},
			AltRbfYC:    []string{"SMSYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".sms", ".sg"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"MegaCD": {
		Id:     "MegaCD",
		Name:   "Sega CD",
		Alias:  []string{"SegaCD"},
		Folder: "MegaCD",
		Rbf:    "_Console/MegaCD",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"MegaCD_LLAPI"},
			AltRbfYC:    []string{"MegaCDYC"},
		},
		FileTypes: []FileType{
			{
				Label: "Disk",
				Exts:  []string{".cue", ".chd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"NeoGeo": {
		Id:     "NeoGeo",
		Name:   "Neo Geo MVS/AES",
		Folder: "NEOGEO",
		Rbf:    "_Console/NeoGeo",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"NeoGeo_LLAPI"},
			AltRbfYC:    []string{"NeoGeoYC"},
		},
		FileTypes: []FileType{
			{
				// TODO: This also has some special handling re: zip files (darksoft pack).
				// Exts: []strings{".*"}
				Label: "ROM set",
				Exts:  []string{".neo"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "CD Image",
				Exts:  []string{".iso", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	"NES": {
		// TODO: Split off NSF music to separate system.
		Id:     "NES",
		Name:   "NES",
		Folder: "NES",
		Rbf:    "_Console/NES",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"NES_LLAPI"},
			AltRbfYC:    []string{"NESYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".nes", ".nsf"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Odyssey2": {
		Id:     "Odyssey2",
		Name:   "Magnavox Odyssey2",
		Folder: "ODYSSEY2",
		Rbf:    "_Console/Odyssey2",
		FileTypes: []FileType{
			{
				Label: "catridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "XROM",
				Exts:  []string{".rom"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"PokemonMini": {
		Id:     "PokemonMini",
		Name:   "Pokemon Mini",
		Folder: "PokemonMini",
		Rbf:    "_Console/PokemonMini",
		FileTypes: []FileType{
			{
				Label: "ROM",
				Exts:  []string{".min"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"PSX": {
		Id:     "PSX",
		Name:   "Playstation",
		Alias:  []string{"Playstation", "PS1"},
		Folder: "PSX",
		Rbf:    "_Console/PSX",
		AltRbf: AltRbfOpts{
			AltRbfYC:      []string{"PSXYC"},
			AltRbfDualRAM: []string{"PSX_DualSDRAM"},
		},
		FileTypes: []FileType{
			{
				Label: "CD",
				Exts:  []string{".cue", ".chd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "Exe",
				Exts:  []string{".exe"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Sega32X": {
		Id:     "Sega32X",
		Name:   "Genesis 32X",
		Alias:  []string{"S32X", "32X"},
		Folder: "S32X",
		Rbf:    "_Console/S32X",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"S32X_LLAPI"},
			AltRbfYC:    []string{"S32XYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".32x"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	// TODO: SG-1000
	//       Include Coleco and SMS folders.
	"SuperGameboy": {
		Id:     "SuperGameboy",
		Name:   "Super Gameboy",
		Alias:  []string{"SGB"},
		Folder: "SGB",
		Rbf:    "_Console/SGB",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"SGB_LLAPI"},
			AltRbfYC:    []string{"SGBYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".gb", ".gbc"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"SuperVision": {
		Id:     "SuperVision",
		Name:   "SuperVision",
		Folder: "SuperVision",
		Rbf:    "_Console/SuperVision",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin", ".sv"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	// TODO: Saturn
	"SNES": {
		// TODO: Split games and music.
		Id:     "SNES",
		Name:   "SNES",
		Alias:  []string{"SuperNintendo"},
		Folder: "SNES",
		Rbf:    "_Console/SNES",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"SNES_LLAPI"},
			AltRbfYC:    []string{"SNESYC"},
		},
		FileTypes: []FileType{
			{
				Exts: []string{".sfc", ".smc", ".bin", ".bs"},
				Mgl: &MglParams{
					Delay:    2,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Exts: []string{".spc"},
				Mgl: &MglParams{
					Delay:    2,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"TurboGraphx16": {
		Id:     "TurboGraphx16",
		Name:   "TurboGraphx-16",
		Alias:  []string{"TGFX16", "PCEngine"},
		Folder: "TGFX16",
		Rbf:    "_Console/TurboGrafx16",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"TurboGrafx16_LLAPI"},
			AltRbfYC:    []string{"TurboGrafx16YC"},
		},
		FileTypes: []FileType{
			{
				Label: "TurboGrafx",
				Exts:  []string{".bin", ".pce"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "SuperGrafx",
				Exts:  []string{".sgx"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"TurboGraphx16CD": {
		Id:     "TurboGraphx16CD",
		Name:   "TurboGraphx-16 CD",
		Alias:  []string{"TGFX16-CD", "PCEngineCD"},
		Folder: "TGFX16-CD",
		Rbf:    "_Console/TurboGrafx16",
		AltRbf: AltRbfOpts{
			AltRbfLLAPI: []string{"TurboGrafx16_LLAPI"},
			AltRbfYC:    []string{"TurboGrafx16YC"},
		},
		FileTypes: []FileType{
			{
				Label: "CD",
				Exts:  []string{".cue", ".chd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"VC4000": {
		Id:     "VC4000",
		Name:   "VC 4000",
		Folder: "VC4000",
		Rbf:    "_Console/VC4000",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Vectrex": {
		Id:     "Vectrex",
		Name:   "Vectrex",
		Folder: "VECTREX",
		Rbf:    "_Console/Vectrex",
		FileTypes: []FileType{
			{
				Exts: []string{".vec", ".bin", ".rom"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "Overlay",
				Exts:  []string{".ovr"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"WonderSwan": {
		Id:     "WonderSwan",
		Name:   "WonderSwan",
		Folder: "WonderSwan",
		Rbf:    "_Console/WonderSwan",
		FileTypes: []FileType{
			{
				Label: "ROM",
				Exts:  []string{".ws", ".wsc"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	// Computers
	"AcornAtom": {
		Id:     "AcornAtom",
		Name:   "Atom",
		Folder: "AcornAtom",
		Rbf:    "_Computer/AcornAtom",
		FileTypes: []FileType{
			{
				Exts: []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"AcornElectron": {
		Id:     "AcornElectron",
		Name:   "Electron",
		Folder: "AcornElectron",
		Rbf:    "_Computer/AcornElectron",
		FileTypes: []FileType{
			{
				Exts: []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"AliceMC10": {
		Id:     "AliceMC10",
		Name:   "Tandy MC-10",
		Folder: "AliceMC10",
		Rbf:    "_Computer/AliceMC10",
		FileTypes: []FileType{
			{
				Label: "Tape",
				Exts:  []string{".c10"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	// TODO: Altair8800
	//       Has a folder but roms are built in.
	"Amiga": {
		// TODO: New versions of MegaAGS image support launching individual games,
		//       will need support for custom scan and launch functions for a core.
		// TODO: This core has 2 .adf drives and 4 .hdf drives.
		Id:        "Amiga",
		Name:      "Amiga",
		Folder:    "Amiga",
		Rbf:       "_Computer/Minimig",
		FileTypes: nil,
	},
	"Amstrad": {
		Id:     "Amstrad",
		Name:   "Amstrad CPC",
		Folder: "Amstrad",
		Rbf:    "_Computer/Amstrad",
		FileTypes: []FileType{
			{
				Label: "A:",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "B:",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "Expansion",
				Exts:  []string{".e??"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    3,
				},
			},
			{
				Label: "Tape",
				Exts:  []string{".cdt"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    4,
				},
			},
		},
	},
	"AmstradPCW": {
		Id:     "AmstradPCW",
		Name:   "Amstrad PCW",
		Folder: "Amstrad PCW",
		Rbf:    "_Computer/Amstrad-PCW",
		FileTypes: []FileType{
			{
				Label: "A:",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "B:",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	"ao486": {
		Id:     "ao486",
		Name:   "PC (486SX)",
		Folder: "AO486",
		Rbf:    "_Computer/ao486",
		FileTypes: []FileType{
			{
				Label: "Floppy A:",
				Exts:  []string{".img", ".ima", ".vfd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "Floppy B:",
				Exts:  []string{".img", ".ima", ".vfd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "IDE 0-0",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    2,
				},
			},
			{
				Label: "IDE 0-1",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    3,
				},
			},
			{
				Label: "IDE 1-0",
				Exts:  []string{".vhd", ".iso", ".cue", ".chd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    4,
				},
			},
			{
				Label: "IDE 1-1",
				Exts:  []string{".vhd", ".iso", ".cue", ".chd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    5,
				},
			},
		},
	},
	"Apogee": {
		Id:     "Apogee",
		Name:   "Apogee BK-01",
		Folder: "APOGEE",
		Rbf:    "_Computer/Apogee",
		FileTypes: []FileType{
			{
				Exts: []string{".rka", ".rkr", ".gam"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"AppleI": {
		Id:     "AppleI",
		Name:   "Apple I",
		Folder: "Apple-I",
		Rbf:    "_Computer/Apple-I",
		FileTypes: []FileType{
			{
				Label: "ASCII",
				Exts:  []string{".txt"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"AppleII": {
		Id:     "AppleII",
		Name:   "Apple IIe",
		Folder: "Apple-II",
		Rbf:    "_Computer/Apple-II",
		FileTypes: []FileType{
			{
				Exts: []string{".nib", ".dsk", ".do", ".po"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Exts: []string{".hdv"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	"Aquarius": {
		Id:     "Aquarius",
		Name:   "Mattel Aquarius",
		Folder: "AQUARIUS",
		Rbf:    "_Computer/Aquarius",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "Tape",
				Exts:  []string{".caq"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	// TODO: Archie
	//       Can't see anything in CONF_STR.
	"Atari800": {
		Id:     "Atari800",
		Name:   "Atari 800XL",
		Folder: "ATARI800",
		Rbf:    "_Computer/Atari800",
		FileTypes: []FileType{
			{
				Label: "D1",
				Exts:  []string{".atr", ".xex", ".xfd", ".atx"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "D2",
				Exts:  []string{".atr", ".xex", ".xfd", ".atx"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "Cartridge",
				Exts:  []string{".car", ".rom", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    2,
				},
			},
		},
	},
	// TODO: AtariST
	//       CONF_STR does not have any information about the file types.
	"BBCMicro": {
		Id:     "BBCMicro",
		Name:   "BBC Micro/Master",
		Folder: "BBCMicro",
		Rbf:    "_Computer/BBCMicro",
		FileTypes: []FileType{
			{
				Exts: []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Exts: []string{".ssd", ".dsd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Exts: []string{".ssd", ".dsd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    2,
				},
			},
		},
	},
	"BK0011M": {
		Id:     "BK0011M",
		Name:   "BK0011M",
		Folder: "BK0011M",
		Rbf:    "_Computer/BK0011M",
		FileTypes: []FileType{
			{
				Exts: []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "FDD(A)",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "FDD(B)",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    2,
				},
			},
			{
				Label: "HDD",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"C16": {
		Id:     "C16",
		Name:   "Commodore 16",
		Folder: "C16",
		Rbf:    "_Computer/C16",
		FileTypes: []FileType{
			{
				Label: "#8",
				Exts:  []string{".d64", ".g64"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "#9",
				Exts:  []string{".d64", ".g64"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				// TODO: This has a hidden option with only .prg and .tap.
				Exts: []string{".prg", ".tap", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"C64": {
		Id:     "C64",
		Name:   "Commodore 64",
		Folder: "C64",
		Rbf:    "_Computer/C64",
		AltRbf: AltRbfOpts{
			AltRbfYC: []string{"C64YC"},
		},
		FileTypes: []FileType{
			{
				Label: "#8",
				Exts:  []string{".d64", ".g64", ".t64", ".d81"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "#9",
				Exts:  []string{".d64", ".g64", ".t64", ".d81"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Exts: []string{".prg", ".crt", ".reu", ".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"CasioPV2000": {
		Id:     "CasioPV2000",
		Name:   "Casio PV-2000",
		Folder: "Casio_PV-2000",
		Rbf:    "_Computer/Casio_PV-2000",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"CoCo2": {
		Id:     "CoCo2",
		Name:   "TRS-80 CoCo 2",
		Folder: "CoCo2",
		Rbf:    "_Computer/CoCo2",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".rom", ".ccc"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "Disk Drive 0",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "Disk Drive 1",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "Disk Drive 2",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    2,
				},
			},
			{
				Label: "Disk Drive 3",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    3,
				},
			},
			{
				Label: "Cassette",
				Exts:  []string{".cas"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	// TODO: CoCo3
	//       This core has several menu states for different combinations of
	//       files to load. Unsure if MGL is compatible with it.
	// TODO: ColecoAdam
	//       Unsure what folder this uses. Coleco?
	"EDSAC": {
		Id:     "EDSAC",
		Name:   "EDSAC",
		Folder: "EDSAC",
		Rbf:    "_Computer/EDSAC",
		FileTypes: []FileType{
			{
				Label: "Tape",
				Exts:  []string{".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Galaksija": {
		Id:     "Galaksija",
		Name:   "Galaksija",
		Folder: "Galaksija",
		Rbf:    "_Computer/Galaksija",
		FileTypes: []FileType{
			{
				Exts: []string{".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Interact": {
		Id:     "Interact",
		Name:   "Interact",
		Folder: "Interact",
		Rbf:    "_Computer/Interact",
		FileTypes: []FileType{
			{
				Label: "Tape",
				Exts:  []string{".cin", ".k7"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Jupiter": {
		Id:     "Jupiter",
		Name:   "Jupiter Ace",
		Folder: "Jupiter",
		Rbf:    "_Computer/Jupiter",
		FileTypes: []FileType{
			{
				Exts: []string{".ace"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Laser310": {
		Id:     "Laser310",
		Name:   "Laser 350/500/700",
		Folder: "Laser",
		Rbf:    "_Computer/Laser310",
		FileTypes: []FileType{
			{
				Label: "VZ Image",
				Exts:  []string{".vz"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"Lynx48": {
		Id:     "Lynx48",
		Name:   "Lynx 48/96K",
		Folder: "Lynx48",
		Rbf:    "_Computer/Lynx48",
		FileTypes: []FileType{
			{
				Label: "Cassette",
				Exts:  []string{".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"MacPlus": {
		Id:     "MacPlus",
		Name:   "Macintosh Plus",
		Folder: "MACPLUS",
		Rbf:    "_Computer/MacPlus",
		FileTypes: []FileType{
			{
				Label: "Pri Floppy",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "Sec Floppy",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
			{
				Label: "SCSI-6",
				Exts:  []string{".img", ".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "SCSI-5",
				Exts:  []string{".img", ".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	"MSX": {
		Id:     "MSX",
		Name:   "MSX",
		Folder: "MSX",
		Rbf:    "_Computer/MSX",
		FileTypes: []FileType{
			{
				Exts: []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"MultiComp": {
		Id:     "MultiComp",
		Name:   "MultiComp",
		Folder: "MultiComp",
		Rbf:    "_Computer/MultiComp",
		FileTypes: []FileType{
			{
				Exts: []string{".img"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	// TODO: OndraSPO186
	//       Nothing listed in CONF_STR but docs do mention loading files.
	"Orao": {
		Id:     "Orao",
		Name:   "Orao",
		Folder: "ORAO",
		Rbf:    "_Computer/ORAO",
		FileTypes: []FileType{
			{
				Exts: []string{".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Oric": {
		Id:     "Oric",
		Name:   "Oric",
		Folder: "Oric",
		Rbf:    "_Computer/Oric",
		FileTypes: []FileType{
			{
				Label: "Drive A:",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	// TODO: PC88
	//       Nothing listed in CONF_STR.
	"PCXT": {
		Id:     "PCXT",
		Name:   "PC/XT",
		Folder: "PCXT",
		Rbf:    "_Computer/PCXT",
		FileTypes: []FileType{
			{
				Label: "FDD Image",
				Exts:  []string{".img", ".ima"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "HDD Image",
				Exts:  []string{".img"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"PDP1": {
		Id:     "PDP1",
		Name:   "PDP-1",
		Folder: "PDP1",
		Rbf:    "_Computer/PDP1",
		FileTypes: []FileType{
			{
				Exts: []string{".pdp", ".rim", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"PET2001": {
		Id:     "PET2001",
		Name:   "Commodore PET 2001",
		Folder: "PET2001",
		Rbf:    "_Computer/PET2001",
		FileTypes: []FileType{
			{
				Exts: []string{".prg", ".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"PMD85": {
		Id:     "PMD85",
		Name:   "PMD 85-2A",
		Folder: "PMD85",
		Rbf:    "_Computer/PMD85",
		FileTypes: []FileType{
			{
				Label: "ROM Pack",
				Exts:  []string{".rmm"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"QL": {
		Id:     "QL",
		Name:   "Sinclair QL",
		Folder: "QL",
		Rbf:    "_Computer/QL",
		FileTypes: []FileType{
			{
				Label: "HD Image",
				Exts:  []string{".win"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "MDV Image",
				Exts:  []string{".mdv"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"RX78": {
		Id:     "RX78",
		Name:   "RX-78 Gundam",
		Folder: "RX78",
		Rbf:    "_Computer/RX78",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"SAMCoupe": {
		Id:     "SAMCoupe",
		Name:   "SAM Coupe",
		Folder: "SAMCOUPE",
		Rbf:    "_Computer/SAMCoupe",
		FileTypes: []FileType{
			{
				Label: "Drive 1",
				Exts:  []string{".dsk", ".mgt", ".img"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "Drive 2",
				Exts:  []string{".dsk", ".mgt", ".img"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	// TODO: SharpMZ
	//       Nothing listed in CONF_STR.
	"SordM5": {
		Id:     "SordM5",
		Name:   "M5",
		Folder: "Sord M5",
		Rbf:    "_Computer/SordM5",
		FileTypes: []FileType{
			{
				Label: "ROM",
				Exts:  []string{".bin", ".rom"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "Tape",
				Exts:  []string{".cas"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"Specialist": {
		Id:     "Specialist",
		Name:   "Specialist/MX",
		Folder: "SPMX",
		Rbf:    "_Computer/Specialist",
		FileTypes: []FileType{
			{
				Label: "Tape",
				Exts:  []string{".rks"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "Disk",
				Exts:  []string{".odi"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"SVI328": {
		Id:     "SVI328",
		Name:   "SV-328",
		Folder: "SVI328",
		Rbf:    "_Computer/Svi328",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin", ".rom"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "CAS File",
				Exts:  []string{".cas"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
		},
	},
	"TatungEinstein": {
		Id:     "TatungEinstein",
		Name:   "Tatung Einstein",
		Folder: "TatungEinstein",
		Rbf:    "_Computer/TatungEinstein",
		FileTypes: []FileType{
			{
				Label: "Disk 0",
				Exts:  []string{".dsk"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"TI994A": {
		Id:     "TI994A",
		Name:   "TI-99/4A",
		Folder: "TI-99_4A",
		Rbf:    "_Computer/Ti994a",
		FileTypes: []FileType{
			{
				Label: "Full Cart",
				Exts:  []string{".m99", ".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
			{
				Label: "ROM Cart",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
			{
				Label: "GROM Cart",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    3,
				},
			},
			// TODO: Also 3 .dsk entries, inactive on first load.
		},
	},
	"TomyTutor": {
		Id:     "TomyTutor",
		Name:   "Tutor",
		Folder: "TomyTutor",
		Rbf:    "_Computer/TomyTutor",
		FileTypes: []FileType{
			{
				Label: "Cartridge",
				Exts:  []string{".bin"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
			{
				Label: "Tape Image",
				Exts:  []string{".cas"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"TRS80": {
		Id:     "TRS80",
		Name:   "TRS-80",
		Folder: "TRS-80",
		Rbf:    "_Computer/TRS-80",
		FileTypes: []FileType{
			{
				Label: "Disk 0",
				Exts:  []string{".dsk", ".jvi"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "Disk 1",
				Exts:  []string{".dsk", ".jvi"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "Program",
				Exts:  []string{".cmd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
			{
				Label: "Cassette",
				Exts:  []string{".cas"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"TSConf": {
		Id:     "TSConf",
		Name:   "TS-Config",
		Folder: "TSConf",
		Rbf:    "_Computer/TSConf",
		FileTypes: []FileType{
			{
				Label: "Virtual SD",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
		},
	},
	"UK101": {
		Id:     "UK101",
		Name:   "UK101",
		Folder: "UK101",
		Rbf:    "_Computer/UK101",
		FileTypes: []FileType{
			{
				Label: "ASCII",
				Exts:  []string{".txt", ".bas", ".lod"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Vector06C": {
		Id:     "Vector06C",
		Name:   "Vector-06C",
		Folder: "VECTOR06",
		Rbf:    "_Computer/Vector-06C",
		FileTypes: []FileType{
			{
				Exts: []string{".rom", ".com", ".c00", ".edd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
			{
				Label: "Disk A",
				Exts:  []string{".fdd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "Disk B",
				Exts:  []string{".fdd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	"VIC20": {
		Id:     "VIC20",
		Name:   "Commodore VIC-20",
		Folder: "VIC20",
		Rbf:    "_Computer/VIC20",
		FileTypes: []FileType{
			{
				Label: "#8",
				Exts:  []string{".d64", ".g64"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "#9",
				Exts:  []string{".d64", ".g64"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Exts: []string{".prg", ".crt", ".ct?", ".tap"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	"X68000": {
		Id:     "X68000",
		Name:   "X68000",
		Folder: "X68000",
		Rbf:    "_Computer/X68000",
		FileTypes: []FileType{
			{
				Label: "FDD0",
				Exts:  []string{".d88"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "FDD1",
				Exts:  []string{".d88"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "SASI Hard Disk",
				Exts:  []string{".hdf"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    2,
				},
			},
			{
				Label: "RAM",
				Exts:  []string{".ram"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    3,
				},
			},
		},
	},
	// TODO: zx48
	//       https://github.com/Kyp069/zx48-MiSTer
	"ZX81": {
		Id:     "ZX81",
		Name:   "TS-1500",
		Folder: "ZX81",
		Rbf:    "_Computer/ZX81",
		FileTypes: []FileType{
			{
				Label: "Tape",
				Exts:  []string{".0", ".p"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"ZXSpectrum": {
		Id:     "ZXSpectrum",
		Name:   "ZX Spectrum",
		Folder: "Spectrum",
		Rbf:    "_Computer/ZX-Spectrum",
		FileTypes: []FileType{
			{
				Label: "Disk",
				Exts:  []string{".trd", ".img", ".dsk", ".mgt"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "Tape",
				Exts:  []string{".tap", ".csw", ".tzx"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    2,
				},
			},
			{
				Label: "Snapshot",
				Exts:  []string{".z80", ".sna"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    4,
				},
			},
			{
				Label: "DivMMC",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
		},
	},
	"ZXNext": {
		Id:     "ZXNext",
		Name:   "ZX Spectrum Next",
		Folder: "ZXNext",
		Rbf:    "_Computer/ZXNext",
		FileTypes: []FileType{
			{
				Label: "C:",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    0,
				},
			},
			{
				Label: "D:",
				Exts:  []string{".vhd"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "s",
					Index:    1,
				},
			},
			{
				Label: "Tape",
				Exts:  []string{".tzx", ".csw"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    1,
				},
			},
		},
	},
	// Other
	"Arcade": {
		Id:     "Arcade",
		Name:   "Arcade",
		Folder: "_Arcade",
		FileTypes: []FileType{
			{
				Exts: []string{".mra"},
				Mgl:  nil,
			},
		},
	},
	"Arduboy": {
		Id:     "Arduboy",
		Name:   "Arduboy",
		Folder: "Arduboy",
		Rbf:    "_Other/Arduboy",
		FileTypes: []FileType{
			{
				Exts: []string{".bin", ".hex"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	"Chip8": {
		Id:     "Chip8",
		Name:   "CHIP-8",
		Folder: "Chip8",
		Rbf:    "_Other/Chip8",
		FileTypes: []FileType{
			{
				Exts: []string{".ch8"},
				Mgl: &MglParams{
					Delay:    1,
					FileType: "f",
					Index:    0,
				},
			},
		},
	},
	// TODO: Life
	//       Has loadable files, but no folder?
	// TODO: ScummVM
	//       Requires a custom scan and launch function.
	// TODO: SuperJacob
	//       A custom computer?
	//       https://github.com/dave18/MiSTER-SuperJacob
	// TODO: TomyScramble
	//       Has loadable files and a folder but is marked as "remove"?
}
