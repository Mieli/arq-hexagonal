package weapon_usecase

type WeaponPresenter struct {
	ID          int64               `json:"oid"`
	Description string              `json:"description"`
	Type        TypeWeaponPresenter `json:"type"`
}

type TypeWeaponPresenter struct {
	Category    string `json:"category"`
	Description string `json:"description"`
}
