package weapon_usecase

type WeaponAssembler struct {
	ID          int64               `json:"id"`
	Description string              `json:"description"`
	Type        TypeWeaponAssembler `json:"type"`
}

type TypeWeaponAssembler struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}
